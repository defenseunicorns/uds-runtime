// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

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

	"github.com/defenseunicorns/uds-engine/pkg/api/monitor"
	"github.com/defenseunicorns/uds-engine/pkg/api/resources"
	"github.com/defenseunicorns/uds-engine/pkg/api/sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
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

	cache.Namespaces.GetResources()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/monitor/pepr/", monitor.Pepr)
		r.Get("/monitor/pepr/{stream}", monitor.Pepr)

		r.Route("/resources", func(r chi.Router) {
			r.Get("/events", sse.Bind[*v1.Event](cache.Events.GetResources, cache.Events.Changes))
			r.Get("/namespaces", sse.Bind[*v1.Namespace](cache.Namespaces.GetResources, cache.Namespaces.Changes))
			r.Get("/pods", sse.Bind[*v1.Pod](cache.Pods.GetResources, cache.Pods.Changes))
			r.Get("/deployments", sse.Bind[*appsv1.Deployment](cache.Deployments.GetResources, cache.Deployments.Changes))
			r.Get("/daemonsets", sse.Bind[*appsv1.DaemonSet](cache.Daemonsets.GetResources, cache.Daemonsets.Changes))
			r.Get("/statefulsets", sse.Bind[*appsv1.StatefulSet](cache.Statefulsets.GetResources, cache.Statefulsets.Changes))

			// Metrics have their own cache and change channel that updates every 30 seconds
			// They do not support informers directly, so we need to poll the API
			r.Get("/podmetrics", sse.Bind(
				func() []*metricsv1beta1.PodMetrics {
					return cache.PodMetrics.GetAll()
				},
				cache.MetricsChanges,
			))

			r.Route("/uds", func(r chi.Router) {
				r.Get("/packages", sse.Bind[*unstructured.Unstructured](cache.UDSPackages.GetResources, cache.UDSPackages.Changes))
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
