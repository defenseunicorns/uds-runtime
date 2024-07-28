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
	"strings"

	udsMiddleware "github.com/defenseunicorns/uds-runtime/pkg/api/middleware"
	"github.com/defenseunicorns/uds-runtime/pkg/api/monitor"
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start(assets embed.FS) error {
	r := chi.NewRouter()

	r.Use(udsMiddleware.ConditionalCompress)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	ctx := context.Background()
	cache, err := resources.NewCache(ctx)
	if err != nil {
		return fmt.Errorf("failed to create cache: %w", err)
	}

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/monitor/pepr/", monitor.Pepr)
		r.Get("/monitor/pepr/{stream}", monitor.Pepr)

		r.Route("/resources", func(r chi.Router) {
			r.Get("/nodes", rest.Bind(cache.Nodes))
			r.Get("/nodes/{uid}", rest.Bind(cache.Nodes))

			r.Get("/events", rest.Bind(cache.Events))
			r.Get("/events/{uid}", rest.Bind(cache.Events))

			r.Get("/namespaces", rest.Bind(cache.Namespaces))
			r.Get("/namespaces/{uid}", rest.Bind(cache.Namespaces))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", rest.Bind(cache.Pods))
				r.Get("/pods/{uid}", rest.Bind(cache.Pods))

				r.Get("/deployments", rest.Bind(cache.Deployments))
				r.Get("/deployments/{uid}", rest.Bind(cache.Deployments))

				r.Get("/daemonsets", rest.Bind(cache.Daemonsets))
				r.Get("/daemonsets/{uid}", rest.Bind(cache.Daemonsets))

				r.Get("/statefulsets", rest.Bind(cache.Statefulsets))
				r.Get("/statefulsets/{uid}", rest.Bind(cache.Statefulsets))

				r.Get("/jobs", rest.Bind(cache.Jobs))
				r.Get("/jobs/{uid}", rest.Bind(cache.Jobs))

				r.Get("/cronjobs", rest.Bind(cache.CronJobs))
				r.Get("/cronjobs/{uid}", rest.Bind(cache.CronJobs))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					rest.SSEHandler(w, r, cache.PodMetrics.GetAll, cache.MetricsChanges, nil)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", rest.Bind(cache.UDSPackages))
				r.Get("/uds-packages/{uid}", rest.Bind(cache.UDSPackages))

				r.Get("/uds-exemptions", rest.Bind(cache.UDSExemptions))
				r.Get("/uds-exemptions/{uid}", rest.Bind(cache.UDSExemptions))

				r.Get("/configmaps", rest.Bind(cache.Configmaps))
				r.Get("/configmaps/{uid}", rest.Bind(cache.Configmaps))

				r.Get("/secrets", rest.Bind(cache.Secrets))
				r.Get("/secrets/{uid}", rest.Bind(cache.Secrets))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", rest.Bind(cache.MutatingWebhooks))
				r.Get("/mutatingwebhooks/{uid}", rest.Bind(cache.MutatingWebhooks))

				r.Get("/validatingwebhooks", rest.Bind(cache.ValidatingWebhooks))
				r.Get("/validatingwebhooks/{uid}", rest.Bind(cache.ValidatingWebhooks))

				r.Get("/hpas", rest.Bind(cache.HPAs))
				r.Get("/hpas/{uid}", rest.Bind(cache.HPAs))

				r.Get("/priority-classes", rest.Bind(cache.PriorityClasses))
				r.Get("/priority-classes/{uid}", rest.Bind(cache.PriorityClasses))

				r.Get("/runtime-classes", rest.Bind(cache.RuntimeClasses))
				r.Get("/runtime-classes/{uid}", rest.Bind(cache.RuntimeClasses))

				r.Get("/poddisruptionbudgets", rest.Bind(cache.PodDisruptionBudgets))
				r.Get("/poddisruptionbudgets/{uid}", rest.Bind(cache.PodDisruptionBudgets))

				r.Get("/limit-ranges", rest.Bind(cache.LimitRanges))
				r.Get("/limit-ranges/{uid}", rest.Bind(cache.LimitRanges))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", rest.Bind(cache.Services))
				r.Get("/services/{uid}", rest.Bind(cache.Services))

				r.Get("/networkpolicies", rest.Bind(cache.NetworkPolicies))
				r.Get("/networkpolicies/{uid}", rest.Bind(cache.NetworkPolicies))

				r.Get("/endpoints", rest.Bind(cache.Endpoints))
				r.Get("/endpoints/{uid}", rest.Bind(cache.Endpoints))

				r.Get("/virtualservices", rest.Bind(cache.VirtualServices))
				r.Get("/virtualservices/{uid}", rest.Bind(cache.VirtualServices))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", rest.Bind(cache.PersistentVolumes))
				r.Get("/persistentvolumes/{uid}", rest.Bind(cache.PersistentVolumes))

				r.Get("/persistentvolumeclaims", rest.Bind(cache.PersistentVolumeClaims))
				r.Get("/persistentvolumeclaims/{uid}", rest.Bind(cache.PersistentVolumeClaims))

				r.Get("/storageclasses", rest.Bind(cache.StorageClasses))
				r.Get("/storageclasses/{uid}", rest.Bind(cache.StorageClasses))
			})
		})
	})

	// Serve static files from embed.FS
	staticFS, err := fs.Sub(assets, "ui/build")
	if err != nil {
		return fmt.Errorf("failed to create static file system: %w", err)
	}

	if err := fileServer(r, http.FS(staticFS)); err != nil {
		return fmt.Errorf("failed to serve static files: %w", err)
	}

	log.Println("Starting server on :8080")
	//nolint:gosec
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}

	return nil
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
