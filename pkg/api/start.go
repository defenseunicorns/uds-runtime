// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package api

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"

	"strings"

	_ "github.com/defenseunicorns/uds-runtime/pkg/api/docs" //nolint:staticcheck
	"github.com/defenseunicorns/uds-runtime/pkg/api/monitor"
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/udsmiddleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title UDS Runtime API
// @version 0.0.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @schemes http https
func Setup(assets *embed.FS) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(udsmiddleware.ConditionalCompress)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ctx := context.Background()
	cache, err := resources.NewCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	// Add Swagger UI route
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/monitor", func(r chi.Router) {
			r.Get("/pepr/", monitor.Pepr)
			r.Get("/pepr/{stream}", monitor.Pepr)
			r.Get("/cluster-overview", monitor.BindClusterOverviewHandler(cache))
		})

		r.Route("/resources", func(r chi.Router) {
			r.Get("/nodes", getNodes(cache))
			r.Get("/nodes/{uid}", getNode(cache))

			r.Get("/events", getEvents(cache))
			r.Get("/events/{uid}", getEvent(cache))

			r.Get("/namespaces", getNamespaces(cache))
			r.Get("/namespaces/{uid}", getNamespace(cache))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", getPods(cache))
				r.Get("/pods/{uid}", getPod(cache))

				r.Get("/deployments", getDeployments(cache))
				r.Get("/deployments/{uid}", getDeployment(cache))

				r.Get("/daemonsets", getDaemonsets(cache))
				r.Get("/daemonsets/{uid}", getDaemonset(cache))

				r.Get("/statefulsets", getStatefulsets(cache))
				r.Get("/statefulsets/{uid}", getStatefulset(cache))

				r.Get("/jobs", getJobs(cache))
				r.Get("/jobs/{uid}", getJob(cache))

				r.Get("/cronjobs", getCronJobs(cache))
				r.Get("/cronjobs/{uid}", getCronJob(cache))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					getPodMetrics(w, r, cache)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", getUDSPackages(cache))
				r.Get("/uds-packages/{uid}", getUDSPackage(cache))

				r.Get("/uds-exemptions", getUDSExemptions(cache))
				r.Get("/uds-exemptions/{uid}", getUDSExemption(cache))

				r.Get("/configmaps", getConfigMaps(cache))
				r.Get("/configmaps/{uid}", getConfigMap(cache))

				r.Get("/secrets", getSecrets(cache))
				r.Get("/secrets/{uid}", getSecret(cache))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", getMutatingWebhooks(cache))
				r.Get("/mutatingwebhooks/{uid}", getMutatingWebhook(cache))

				r.Get("/validatingwebhooks", getValidatingWebhooks(cache))
				r.Get("/validatingwebhooks/{uid}", getValidatingWebhook(cache))

				r.Get("/hpas", getHPAs(cache))
				r.Get("/hpas/{uid}", getHPA(cache))

				r.Get("/priority-classes", getPriorityClasses(cache))
				r.Get("/priority-classes/{uid}", getPriorityClass(cache))

				r.Get("/runtime-classes", getRuntimeClasses(cache))
				r.Get("/runtime-classes/{uid}", getRuntimeClass(cache))

				r.Get("/poddisruptionbudgets", getPodDisruptionBudgets(cache))
				r.Get("/poddisruptionbudgets/{uid}", getPodDisruptionBudget(cache))

				r.Get("/limit-ranges", getLimitRanges(cache))
				r.Get("/limit-ranges/{uid}", getLimitRange(cache))

				r.Get("/resource-quotas", getResourceQuotas(cache))
				r.Get("/resource-quotas/{uid}", getResourceQuota(cache))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", getServices(cache))
				r.Get("/services/{uid}", getService(cache))

				r.Get("/networkpolicies", getNetworkPolicies(cache))
				r.Get("/networkpolicies/{uid}", getNetworkPolicy(cache))

				r.Get("/endpoints", getEndpoints(cache))
				r.Get("/endpoints/{uid}", getEndpoint(cache))

				r.Get("/virtualservices", getVirtualServices(cache))
				r.Get("/virtualservices/{uid}", getVirtualService(cache))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", getPersistentVolumes(cache))
				r.Get("/persistentvolumes/{uid}", getPersistentVolume(cache))

				r.Get("/persistentvolumeclaims", getPersistentVolumeClaims(cache))
				r.Get("/persistentvolumeclaims/{uid}", getPersistentVolumeClaim(cache))

				r.Get("/storageclasses", getStorageClasses(cache))
				r.Get("/storageclasses/{uid}", getStorageClass(cache))
			})
		})
	})

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
