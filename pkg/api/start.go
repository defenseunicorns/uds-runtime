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
	"time"

	"strings"

	"encoding/json"

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
)

type K8sResources struct {
	Client *k8s.Clients
	Cache  *resources.Cache
	Cancel context.CancelFunc
}

// @title UDS Runtime API
// @version 0.0.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @schemes http https
func Setup(assets *embed.FS) (*chi.Mux, error) {
	apiAuth := true
	if strings.ToLower(os.Getenv("API_AUTH_DISABLED")) == "true" {
		apiAuth = false
	}
	port := "8080"

	ip := "127.0.0.1"

	// If the env variable API_TOKEN is set, use that for the API secret
	token := os.Getenv("API_TOKEN")
	var err error
	// Otherwise, generate a random secret
	if token == "" {
		token, err = auth.RandomString(96)
		if err != nil {
			return nil, fmt.Errorf("failed to generate random string: %w", err)
		}
	}

	r := chi.NewRouter()

	r.Use(udsMiddleware.ConditionalCompress)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ctx, cancel := context.WithCancel(context.Background())
	k8sClient, err := k8s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %w", err)
	}

	cache, err := resources.NewCache(ctx, k8sClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	// Create the disconnected channel
	disconnected := make(chan error)

	// Create a K8sResources struct to hold the references
	k8sResources := &K8sResources{
		Client: k8sClient,
		Cache:  cache,
		Cancel: cancel,
	}

	// Goroutine to handle retries on disconnection
	go func() {
		for {
			select {
			case err := <-disconnected:
				fmt.Printf("Disconnected error received: %v\n", err)
				for {
					time.Sleep(5 * time.Second) // Retry interval
					k8sClient, err := k8s.NewClient()
					if err != nil {
						fmt.Printf("Retrying to create k8s client: %v\n", err)
						continue
					}
					// Cancel the previous context and create a new one
					k8sResources.Cancel()
					ctx, cancel := context.WithCancel(context.Background())
					cache, err = resources.NewCache(ctx, k8sClient)
					if err != nil {
						fmt.Printf("Retrying to create cache: %v\n", err)
						continue
					}
					// Update the references in the K8sResources struct
					k8sResources.Client = k8sClient
					k8sResources.Cache = cache
					k8sResources.Cancel = cancel
					fmt.Println("Successfully reconnected to k8s and recreated cache")
					break
				}
			}
		}
	}()

	// Add Swagger UI routes
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// expose API_AUTH_DISABLED env var to frontend via endpoint
	r.Get("/auth-status", serveAuthStatus)
	r.Get("/health", checkHealth(k8sResources, disconnected))
	r.Route("/api/v1", func(r chi.Router) {
		// Require a valid token for API calls
		if apiAuth {
			// If api auth is enabled, require a valid token for all routes under /api/v1
			r.Use(auth.RequireSecret(token))
			// Endpoint to test if connected with auth
			r.Head("/", auth.Connect)
		}
		r.Route("/monitor", func(r chi.Router) {
			r.Get("/pepr/", monitor.Pepr)
			r.Get("/pepr/{stream}", monitor.Pepr)
			r.Get("/cluster-overview", monitor.BindClusterOverviewHandler(k8sResources.Cache))
		})

		r.Route("/resources", func(r chi.Router) {
			r.Get("/nodes", getNodes(k8sResources.Cache))
			r.Get("/nodes/{uid}", getNode(k8sResources.Cache))

			r.Get("/events", getEvents(k8sResources.Cache))
			r.Get("/events/{uid}", getEvent(k8sResources.Cache))

			r.Get("/namespaces", getNamespaces(k8sResources.Cache))
			r.Get("/namespaces/{uid}", getNamespace(k8sResources.Cache))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", func(w http.ResponseWriter, r *http.Request) { getPods(k8sResources.Cache)(w, r) })
				r.Get("/pods/{uid}", getPod(k8sResources.Cache))

				r.Get("/deployments", getDeployments(k8sResources.Cache))
				r.Get("/deployments/{uid}", getDeployment(k8sResources.Cache))

				r.Get("/daemonsets", getDaemonsets(k8sResources.Cache))
				r.Get("/daemonsets/{uid}", getDaemonset(k8sResources.Cache))

				r.Get("/statefulsets", getStatefulsets(k8sResources.Cache))
				r.Get("/statefulsets/{uid}", getStatefulset(k8sResources.Cache))

				r.Get("/jobs", getJobs(k8sResources.Cache))
				r.Get("/jobs/{uid}", getJob(k8sResources.Cache))

				r.Get("/cronjobs", getCronJobs(k8sResources.Cache))
				r.Get("/cronjobs/{uid}", getCronJob(k8sResources.Cache))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					getPodMetrics(w, r, k8sResources.Cache)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", getUDSPackages(k8sResources.Cache))
				r.Get("/uds-packages/{uid}", getUDSPackage(k8sResources.Cache))

				r.Get("/uds-exemptions", getUDSExemptions(k8sResources.Cache))
				r.Get("/uds-exemptions/{uid}", getUDSExemption(k8sResources.Cache))

				r.Get("/configmaps", getConfigMaps(k8sResources.Cache))
				r.Get("/configmaps/{uid}", getConfigMap(k8sResources.Cache))

				r.Get("/secrets", getSecrets(k8sResources.Cache))
				r.Get("/secrets/{uid}", getSecret(k8sResources.Cache))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", getMutatingWebhooks(k8sResources.Cache))
				r.Get("/mutatingwebhooks/{uid}", getMutatingWebhook(k8sResources.Cache))

				r.Get("/validatingwebhooks", getValidatingWebhooks(k8sResources.Cache))
				r.Get("/validatingwebhooks/{uid}", getValidatingWebhook(k8sResources.Cache))

				r.Get("/hpas", getHPAs(k8sResources.Cache))
				r.Get("/hpas/{uid}", getHPA(k8sResources.Cache))

				r.Get("/priority-classes", getPriorityClasses(k8sResources.Cache))
				r.Get("/priority-classes/{uid}", getPriorityClass(k8sResources.Cache))

				r.Get("/runtime-classes", getRuntimeClasses(k8sResources.Cache))
				r.Get("/runtime-classes/{uid}", getRuntimeClass(k8sResources.Cache))

				r.Get("/poddisruptionbudgets", getPodDisruptionBudgets(k8sResources.Cache))
				r.Get("/poddisruptionbudgets/{uid}", getPodDisruptionBudget(k8sResources.Cache))

				r.Get("/limit-ranges", getLimitRanges(k8sResources.Cache))
				r.Get("/limit-ranges/{uid}", getLimitRange(k8sResources.Cache))

				r.Get("/resource-quotas", getResourceQuotas(k8sResources.Cache))
				r.Get("/resource-quotas/{uid}", getResourceQuota(k8sResources.Cache))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", getServices(k8sResources.Cache))
				r.Get("/services/{uid}", getService(k8sResources.Cache))

				r.Get("/networkpolicies", getNetworkPolicies(k8sResources.Cache))
				r.Get("/networkpolicies/{uid}", getNetworkPolicy(k8sResources.Cache))

				r.Get("/endpoints", getEndpoints(k8sResources.Cache))
				r.Get("/endpoints/{uid}", getEndpoint(k8sResources.Cache))

				r.Get("/virtualservices", getVirtualServices(k8sResources.Cache))
				r.Get("/virtualservices/{uid}", getVirtualService(k8sResources.Cache))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", getPersistentVolumes(k8sResources.Cache))
				r.Get("/persistentvolumes/{uid}", getPersistentVolume(k8sResources.Cache))

				r.Get("/persistentvolumeclaims", getPersistentVolumeClaims(k8sResources.Cache))
				r.Get("/persistentvolumeclaims/{uid}", getPersistentVolumeClaim(k8sResources.Cache))

				r.Get("/storageclasses", getStorageClasses(k8sResources.Cache))
				r.Get("/storageclasses/{uid}", getStorageClass(k8sResources.Cache))
			})
		})
	})

	if apiAuth {
		colorYellow := "\033[33m"
		colorReset := "\033[0m"
		url := fmt.Sprintf("http://%s:%s/auth?token=%s", ip, port, token)
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

func serveAuthStatus(w http.ResponseWriter, _ *http.Request) {
	authStatus := map[string]string{
		"API_AUTH_DISABLED": os.Getenv("API_AUTH_DISABLED"),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(authStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
