// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package api

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"strings"

	"github.com/defenseunicorns/pkg/exec"
	"github.com/defenseunicorns/uds-runtime/pkg/api/auth"
	_ "github.com/defenseunicorns/uds-runtime/pkg/api/docs" //nolint:staticcheck
	udsMiddleware "github.com/defenseunicorns/uds-runtime/pkg/api/middleware"
	"github.com/defenseunicorns/uds-runtime/pkg/api/monitor"
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"k8s.io/client-go/rest"
)

type K8sResources struct {
	client         *k8s.Clients
	cache          *resources.Cache
	currentCtx     string
	currentCluster string
	cancel         context.CancelFunc
}

// @title UDS Runtime API
// @version 0.0.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @schemes http https
func Setup(assets *embed.FS) (*chi.Mux, error) {
	apiAuth, token, err := checkForLocalAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to set auth: %w", err)
	}

	authSVC := checkForClusterAuth()

	r := chi.NewRouter()

	r.Use(udsMiddleware.ConditionalCompress)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if authSVC {
		r.Use(auth.RequireJWT)
	}

	// Middleware chain for api token authentication
	apiAuthMiddleware := func(next http.Handler) http.Handler {
		if apiAuth {
			return udsMiddleware.ValidateSession(next)
		}
		return next
	}

	// Setup k8s resources
	k8sResources, err := setupK8sResources()
	if err != nil {
		return nil, fmt.Errorf("failed to setup k8s resources: %w", err)
	}

	// Create the disconnected channel
	disconnected := make(chan error)

	inCluster, err := isRunningInCluster()
	if err != nil {
		k8sResources.cancel()
		return nil, fmt.Errorf("failed to check if running in cluster: %w", err)
	}

	// Get current k8s context and start the reconnection goroutine if NOT in cluster
	if !inCluster {
		currentCtx, currentCluster, err := k8s.GetCurrentContext()
		if err != nil {
			k8sResources.cancel()
			return nil, fmt.Errorf("failed to get current context: %w", err)
		}

		k8sResources.currentCtx = currentCtx
		k8sResources.currentCluster = currentCluster

		go handleReconnection(disconnected, k8sResources, k8s.NewClient, resources.NewCache)
	}

	// Add Swagger UI routes
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/health", checkHealth(k8sResources, disconnected))
	r.Route("/api/v1", func(r chi.Router) {
		// Require a valid token for API calls
		if apiAuth {
			// If api auth is enabled, require a valid token for all routes under /api/v1
			// authenticate token
			r.With(auth.TokenAuthenticator(token)).Head("/api-auth", func(_ http.ResponseWriter, _ *http.Request) {})
		} else {
			r.Head("/api-auth", func(_ http.ResponseWriter, _ *http.Request) {})
		}

		r.With(apiAuthMiddleware).Route("/monitor", func(r chi.Router) {
			r.Get("/pepr/", monitor.Pepr)
			r.Get("/pepr/{stream}", monitor.Pepr)
			r.Get("/cluster-overview", monitor.BindClusterOverviewHandler(k8sResources.cache))
		})

		r.With(apiAuthMiddleware).Route("/resources", func(r chi.Router) {
			r.Get("/nodes", withLatestCache(k8sResources, getNodes))
			r.Get("/nodes/{uid}", withLatestCache(k8sResources, getNode))

			r.Get("/events", withLatestCache(k8sResources, getEvents))
			r.Get("/events/{uid}", withLatestCache(k8sResources, getEvent))

			r.Get("/namespaces", withLatestCache(k8sResources, getNamespaces))
			r.Get("/namespaces/{uid}", withLatestCache(k8sResources, getNamespace))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", withLatestCache(k8sResources, getPods))
				r.Get("/pods/{uid}", withLatestCache(k8sResources, getPod))

				r.Get("/deployments", withLatestCache(k8sResources, getDeployments))
				r.Get("/deployments/{uid}", withLatestCache(k8sResources, getDeployment))

				r.Get("/daemonsets", withLatestCache(k8sResources, getDaemonsets))
				r.Get("/daemonsets/{uid}", withLatestCache(k8sResources, getDaemonset))

				r.Get("/statefulsets", withLatestCache(k8sResources, getStatefulsets))
				r.Get("/statefulsets/{uid}", withLatestCache(k8sResources, getStatefulset))

				r.Get("/jobs", withLatestCache(k8sResources, getJobs))
				r.Get("/jobs/{uid}", withLatestCache(k8sResources, getJob))

				r.Get("/cronjobs", withLatestCache(k8sResources, getCronJobs))
				r.Get("/cronjobs/{uid}", withLatestCache(k8sResources, getCronJob))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					getPodMetrics(w, r, k8sResources.cache)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", withLatestCache(k8sResources, getUDSPackages))
				r.Get("/uds-packages/{uid}", withLatestCache(k8sResources, getUDSPackage))

				r.Get("/uds-exemptions", withLatestCache(k8sResources, getUDSExemptions))
				r.Get("/uds-exemptions/{uid}", withLatestCache(k8sResources, getUDSExemption))

				r.Get("/configmaps", withLatestCache(k8sResources, getConfigMaps))
				r.Get("/configmaps/{uid}", withLatestCache(k8sResources, getConfigMap))

				r.Get("/secrets", withLatestCache(k8sResources, getSecrets))
				r.Get("/secrets/{uid}", withLatestCache(k8sResources, getSecret))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", withLatestCache(k8sResources, getMutatingWebhooks))
				r.Get("/mutatingwebhooks/{uid}", withLatestCache(k8sResources, getMutatingWebhook))

				r.Get("/validatingwebhooks", withLatestCache(k8sResources, getValidatingWebhooks))
				r.Get("/validatingwebhooks/{uid}", withLatestCache(k8sResources, getValidatingWebhook))

				r.Get("/hpas", withLatestCache(k8sResources, getHPAs))
				r.Get("/hpas/{uid}", withLatestCache(k8sResources, getHPA))

				r.Get("/priority-classes", withLatestCache(k8sResources, getPriorityClasses))
				r.Get("/priority-classes/{uid}", withLatestCache(k8sResources, getPriorityClass))

				r.Get("/runtime-classes", withLatestCache(k8sResources, getRuntimeClasses))
				r.Get("/runtime-classes/{uid}", withLatestCache(k8sResources, getRuntimeClass))

				r.Get("/poddisruptionbudgets", withLatestCache(k8sResources, getPodDisruptionBudgets))
				r.Get("/poddisruptionbudgets/{uid}", withLatestCache(k8sResources, getPodDisruptionBudget))

				r.Get("/limit-ranges", withLatestCache(k8sResources, getLimitRanges))
				r.Get("/limit-ranges/{uid}", withLatestCache(k8sResources, getLimitRange))

				r.Get("/resource-quotas", withLatestCache(k8sResources, getResourceQuotas))
				r.Get("/resource-quotas/{uid}", withLatestCache(k8sResources, getResourceQuota))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", withLatestCache(k8sResources, getServices))
				r.Get("/services/{uid}", withLatestCache(k8sResources, getService))

				r.Get("/networkpolicies", withLatestCache(k8sResources, getNetworkPolicies))
				r.Get("/networkpolicies/{uid}", withLatestCache(k8sResources, getNetworkPolicy))

				r.Get("/endpoints", withLatestCache(k8sResources, getEndpoints))
				r.Get("/endpoints/{uid}", withLatestCache(k8sResources, getEndpoint))

				r.Get("/virtualservices", withLatestCache(k8sResources, getVirtualServices))
				r.Get("/virtualservices/{uid}", withLatestCache(k8sResources, getVirtualService))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", withLatestCache(k8sResources, getPersistentVolumes))
				r.Get("/persistentvolumes/{uid}", withLatestCache(k8sResources, getPersistentVolume))

				r.Get("/persistentvolumeclaims", withLatestCache(k8sResources, getPersistentVolumeClaims))
				r.Get("/persistentvolumeclaims/{uid}", withLatestCache(k8sResources, getPersistentVolumeClaim))

				r.Get("/storageclasses", withLatestCache(k8sResources, getStorageClasses))
				r.Get("/storageclasses/{uid}", withLatestCache(k8sResources, getStorageClass))
			})
		})

		r.With(apiAuthMiddleware).Route("/security", func(r chi.Router) {
			r.Get("/reports", getSecurityReports)
		})
	})

	if apiAuth {
		port := "8080"
		ip := "127.0.0.1"
		colorYellow := "\033[33m"
		colorReset := "\033[0m"
		url := fmt.Sprintf("http://%s:%s?token=%s", ip, port, token)
		log.Printf("%sRuntime API connection: %s%s", colorYellow, url, colorReset)
		err := exec.LaunchURL(url)
		if err != nil {
			return nil, fmt.Errorf("failed to launch URL: %w", err)
		}
	}

	// Serve static files from embed.FS
	if assets != nil {
		staticFS, err := fs.Sub(assets, "ui/build")
		if err != nil {
			return nil, fmt.Errorf("failed to create static file system: %w", err)
		}

		if err := fileServer(r, http.FS(staticFS)); err != nil {
			return nil, fmt.Errorf("failed to serve static files: %w", err)
		}
	}
	return r, nil
}

func setupK8sResources() (*K8sResources, error) {
	k8sClient, err := k8s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cache, err := resources.NewCache(ctx, k8sClient)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	// K8sResources struct to hold references
	k8sResources := &K8sResources{
		client: k8sClient,
		cache:  cache,
		cancel: cancel,
	}

	return k8sResources, nil
}

// fileServer is a custom file server handler for embedded files
func fileServer(r chi.Router, root http.FileSystem) error {
	// Load index.html content and modification time at startup
	f, err := root.Open("index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}
	indexModTime := stat.ModTime()

	indexHTML, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// Create a new file server handler
	fs := http.FileServer(root)

	// Serve the index.html file if the requested file doesn't exist
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Try to open the file from the embedded filesystem
		file, err := root.Open(r.URL.Path)
		if err != nil {
			// If the file doesn't exist, serve the pre-loaded index.html
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			// Serve the index.html file with the pre-loaded content
			http.ServeContent(w, r, "index.html", indexModTime, strings.NewReader(string(indexHTML)))
			return
		}
		file.Close()

		// If the file exists, serve it using the http.FileServer
		fs.ServeHTTP(w, r)
	})

	return nil
}

func checkForLocalAuth() (bool, string, error) {
	apiAuth := true
	if strings.ToLower(os.Getenv("API_AUTH_DISABLED")) == "true" {
		apiAuth = false
	}

	// If the env variable API_TOKEN is set, use that for the API secret
	token := os.Getenv("API_TOKEN")
	var err error
	// Otherwise, generate a random secret
	if token == "" {
		token, err = auth.RandomString(96)
		if err != nil {
			return true, "", fmt.Errorf("failed to generate random string: %w", err)
		}
	}

	return apiAuth, token, nil
}

func checkForClusterAuth() bool {
	authSVC := false
	if strings.ToLower(os.Getenv("AUTH_SVC_ENABLED")) == "true" {
		authSVC = true
	}

	return authSVC
}

// withLatestCache returns a wrapper lambda function, creating a closure that can dynamically access the latest cache
func withLatestCache(k8sResources *K8sResources, handler func(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(k8sResources.cache)(w, r)
	}
}

type createClient func() (*k8s.Clients, error)
type createCache func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error)

func getRetryInterval() time.Duration {
	if interval, exists := os.LookupEnv("RETRY_INTERVAL_MS"); exists {
		parsed, err := strconv.Atoi(interval)
		if err == nil {
			return time.Duration(parsed) * time.Millisecond
		}
	}
	return 5 * time.Second // Default to 5 seconds if not set
}

// isRunningInCluster checks if the application is running in cluster
func isRunningInCluster() (bool, error) {
	_, err := rest.InClusterConfig()

	if err == rest.ErrNotInCluster {
		return false, nil
	} else if err != nil {
		return true, err
	}

	return true, nil
}

// handleReconnection is a goroutine that handles reconnection to the k8s API
// passing createClient and createCache instead of calling k8s.NewClient and resources.NewCache for testing purposes
func handleReconnection(disconnected chan error, k8sResources *K8sResources, createClient createClient,
	createCache createCache) {
	for err := range disconnected {
		log.Printf("Disconnected error received: %v\n", err)
		for {
			// Cancel the previous context
			k8sResources.cancel()
			time.Sleep(getRetryInterval())

			currentCtx, currentCluster, err := k8s.GetCurrentContext()
			if err != nil {
				log.Printf("Error fetching current context: %v\n", err)
				continue
			}

			// If the current context or cluster is different from the original, skip reconnection
			if currentCtx != k8sResources.currentCtx || currentCluster != k8sResources.currentCluster {
				log.Println("Current context has changed. Skipping reconnection.")
				continue
			}

			k8sClient, err := createClient()
			if err != nil {
				log.Printf("Retrying to create k8s client: %v\n", err)
				continue
			}

			// Create a new context and cache
			ctx, cancel := context.WithCancel(context.Background())
			cache, err := createCache(ctx, k8sClient)
			if err != nil {
				log.Printf("Retrying to create cache: %v\n", err)
				continue
			}

			k8sResources.client = k8sClient
			k8sResources.cache = cache
			k8sResources.cancel = cancel
			log.Println("Successfully reconnected to k8s and recreated cache")
			break
		}
	}
}
