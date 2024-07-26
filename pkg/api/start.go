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

	"github.com/defenseunicorns/uds-runtime/pkg/api/monitor"
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start(assets embed.FS) error {
	r := chi.NewRouter()

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
			r.Get("/nodes", sse.Bind(cache.Nodes))
			r.Get("/nodes/{uid}", sse.Bind(cache.Nodes))

			r.Get("/events", sse.Bind(cache.Events))
			r.Get("/events/{uid}", sse.Bind(cache.Events))

			r.Get("/namespaces", sse.Bind(cache.Namespaces))
			r.Get("/namespaces/{uid}", sse.Bind(cache.Namespaces))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", sse.Bind(cache.Pods))
				r.Get("/pods/{uid}", sse.Bind(cache.Pods))

				r.Get("/deployments", sse.Bind(cache.Deployments))
				r.Get("/deployments/{uid}", sse.Bind(cache.Deployments))

				r.Get("/daemonsets", sse.Bind(cache.Daemonsets))
				r.Get("/daemonsets/{uid}", sse.Bind(cache.Daemonsets))

				r.Get("/statefulsets", sse.Bind(cache.Statefulsets))
				r.Get("/statefulsets/{uid}", sse.Bind(cache.Statefulsets))

				r.Get("/jobs", sse.Bind(cache.Jobs))
				r.Get("/jobs/{uid}", sse.Bind(cache.Jobs))

				r.Get("/cronjobs", sse.Bind(cache.CronJobs))
				r.Get("/cronjobs/{uid}", sse.Bind(cache.CronJobs))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					sse.Handler(w, r, cache.PodMetrics.GetAll, cache.MetricsChanges)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", sse.Bind(cache.UDSPackages))
				r.Get("/uds-packages/{uid}", sse.Bind(cache.UDSPackages))

				r.Get("/uds-exemptions", sse.Bind(cache.UDSExemptions))
				r.Get("/uds-exemptions/{uid}", sse.Bind(cache.UDSExemptions))

				r.Get("/configmaps", sse.Bind(cache.Configmaps))
				r.Get("/configmaps/{uid}", sse.Bind(cache.Configmaps))

				r.Get("/secrets", sse.Bind(cache.Secrets))
				r.Get("/secrets/{uid}", sse.Bind(cache.Secrets))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", sse.Bind(cache.MutatingWebhooks))
				r.Get("/mutatingwebhooks/{uid}", sse.Bind(cache.MutatingWebhooks))

				r.Get("/validatingwebhooks", sse.Bind(cache.ValidatingWebhooks))
				r.Get("/validatingwebhooks/{uid}", sse.Bind(cache.ValidatingWebhooks))

				r.Get("/hpas", sse.Bind(cache.HPAs))
				r.Get("/hpas/{uid}", sse.Bind(cache.HPAs))

				r.Get("/priority-classes", sse.Bind(cache.PriorityClasses))
				r.Get("/priority-classes/{uid}", sse.Bind(cache.PriorityClasses))

				r.Get("/runtime-classes", sse.Bind(cache.RuntimeClasses))
				r.Get("/runtime-classes/{uid}", sse.Bind(cache.RuntimeClasses))

				r.Get("/limit-ranges", sse.Bind(cache.LimitRanges))
				r.Get("/limit-ranges/{uid}", sse.Bind(cache.LimitRanges))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", sse.Bind(cache.Services))
				r.Get("/services/{uid}", sse.Bind(cache.Services))

				r.Get("/networkpolicies", sse.Bind(cache.NetworkPolicies))
				r.Get("/networkpolicies/{uid}", sse.Bind(cache.NetworkPolicies))

				r.Get("/endpoints", sse.Bind(cache.Endpoints))
				r.Get("/endpoints/{uid}", sse.Bind(cache.Endpoints))

				r.Get("/virtualservices", sse.Bind(cache.VirtualServices))
				r.Get("/virtualservices/{uid}", sse.Bind(cache.VirtualServices))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", sse.Bind(cache.PersistentVolumes))
				r.Get("/persistentvolumes/{uid}", sse.Bind(cache.PersistentVolumes))

				r.Get("/persistentvolumeclaims", sse.Bind(cache.PersistentVolumeClaims))
				r.Get("/persistentvolumeclaims/{uid}", sse.Bind(cache.PersistentVolumeClaims))

				r.Get("/storageclasses", sse.Bind(cache.StorageClasses))
				r.Get("/storageclasses/{uid}", sse.Bind(cache.StorageClasses))
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
