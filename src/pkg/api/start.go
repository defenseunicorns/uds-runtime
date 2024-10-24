// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package api

import (
	"crypto/tls"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/defenseunicorns/pkg/exec"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	_ "github.com/defenseunicorns/uds-runtime/src/pkg/api/docs" //nolint:staticcheck
	udsMiddleware "github.com/defenseunicorns/uds-runtime/src/pkg/api/middleware"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/monitor"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
	"github.com/defenseunicorns/uds-runtime/src/pkg/k8s/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

// @title UDS Runtime API
// @version 0.0.0
// @BasePath /api/v1
// @schemes http https
func Setup(assets *embed.FS) (*chi.Mux, bool, error) {
	// configure config vars for local or in-cluster auth
	err := auth.Configure()
	if err != nil {
		return nil, false, fmt.Errorf("failed to configure auth: %w", err)
	}

	// Create a k8s session
	k8sSession, err := session.CreateK8sSession()
	if err != nil {
		return nil, false, fmt.Errorf("failed to setup k8s session: %w", err)
	}

	inCluster := k8sSession.InCluster

	if !inCluster {
		// Start the cluster monitoring goroutine
		go k8sSession.StartClusterMonitoring()
	}

	r := chi.NewRouter()

	// Add middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(udsMiddleware.Auth)
	r.Use(udsMiddleware.ConditionalCompress)

	// add routes
	r.Get("/healthz", healthz)
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/cluster-check", checkClusterConnection(k8sSession))
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/auth", authHandler)
		r.Route("/monitor", func(r chi.Router) {
			r.Get("/pepr/", monitor.Pepr)
			r.Get("/pepr/{stream}", monitor.Pepr)
			r.Get("/cluster-overview", monitor.BindClusterOverviewHandler(k8sSession.Cache))
		})

		r.Route("/resources", func(r chi.Router) {
			r.Get("/nodes", withLatestCache(k8sSession, getNodes))
			r.Get("/nodes/{uid}", withLatestCache(k8sSession, getNode))

			r.Get("/events", withLatestCache(k8sSession, getEvents))
			r.Get("/events/{uid}", withLatestCache(k8sSession, getEvent))

			r.Get("/namespaces", withLatestCache(k8sSession, getNamespaces))
			r.Get("/namespaces/{uid}", withLatestCache(k8sSession, getNamespace))

			r.Get("/custom-resource-definitions", withLatestCache(k8sSession, getCRDs))
			r.Get("/custom-resource-definitions/{uid}", withLatestCache(k8sSession, getCRD))

			// Workload resources
			r.Route("/workloads", func(r chi.Router) {
				r.Get("/pods", withLatestCache(k8sSession, getPods))
				r.Get("/pods/{uid}", withLatestCache(k8sSession, getPod))

				r.Get("/deployments", withLatestCache(k8sSession, getDeployments))
				r.Get("/deployments/{uid}", withLatestCache(k8sSession, getDeployment))

				r.Get("/daemonsets", withLatestCache(k8sSession, getDaemonsets))
				r.Get("/daemonsets/{uid}", withLatestCache(k8sSession, getDaemonset))

				r.Get("/statefulsets", withLatestCache(k8sSession, getStatefulsets))
				r.Get("/statefulsets/{uid}", withLatestCache(k8sSession, getStatefulset))

				r.Get("/jobs", withLatestCache(k8sSession, getJobs))
				r.Get("/jobs/{uid}", withLatestCache(k8sSession, getJob))

				r.Get("/cronjobs", withLatestCache(k8sSession, getCronJobs))
				r.Get("/cronjobs/{uid}", withLatestCache(k8sSession, getCronJob))

				// Metrics have their own cache and change channel that updates every 30 seconds
				// They do not support informers directly, so we need to poll the API
				r.Get("/podmetrics", func(w http.ResponseWriter, r *http.Request) {
					getPodMetrics(w, r, k8sSession.Cache)
				})
			})

			// Config resources
			r.Route("/configs", func(r chi.Router) {
				r.Get("/uds-packages", withLatestCache(k8sSession, getUDSPackages))
				r.Get("/uds-packages/{uid}", withLatestCache(k8sSession, getUDSPackage))

				r.Get("/uds-exemptions", withLatestCache(k8sSession, getUDSExemptions))
				r.Get("/uds-exemptions/{uid}", withLatestCache(k8sSession, getUDSExemption))

				r.Get("/configmaps", withLatestCache(k8sSession, getConfigMaps))
				r.Get("/configmaps/{uid}", withLatestCache(k8sSession, getConfigMap))

				r.Get("/secrets", withLatestCache(k8sSession, getSecrets))
				r.Get("/secrets/{uid}", withLatestCache(k8sSession, getSecret))
			})

			// Cluster ops resources
			r.Route("/cluster-ops", func(r chi.Router) {
				r.Get("/mutatingwebhooks", withLatestCache(k8sSession, getMutatingWebhooks))
				r.Get("/mutatingwebhooks/{uid}", withLatestCache(k8sSession, getMutatingWebhook))

				r.Get("/validatingwebhooks", withLatestCache(k8sSession, getValidatingWebhooks))
				r.Get("/validatingwebhooks/{uid}", withLatestCache(k8sSession, getValidatingWebhook))

				r.Get("/hpas", withLatestCache(k8sSession, getHPAs))
				r.Get("/hpas/{uid}", withLatestCache(k8sSession, getHPA))

				r.Get("/priority-classes", withLatestCache(k8sSession, getPriorityClasses))
				r.Get("/priority-classes/{uid}", withLatestCache(k8sSession, getPriorityClass))

				r.Get("/runtime-classes", withLatestCache(k8sSession, getRuntimeClasses))
				r.Get("/runtime-classes/{uid}", withLatestCache(k8sSession, getRuntimeClass))

				r.Get("/poddisruptionbudgets", withLatestCache(k8sSession, getPodDisruptionBudgets))
				r.Get("/poddisruptionbudgets/{uid}", withLatestCache(k8sSession, getPodDisruptionBudget))

				r.Get("/limit-ranges", withLatestCache(k8sSession, getLimitRanges))
				r.Get("/limit-ranges/{uid}", withLatestCache(k8sSession, getLimitRange))

				r.Get("/resource-quotas", withLatestCache(k8sSession, getResourceQuotas))
				r.Get("/resource-quotas/{uid}", withLatestCache(k8sSession, getResourceQuota))
			})

			// Network resources
			r.Route("/networks", func(r chi.Router) {
				r.Get("/services", withLatestCache(k8sSession, getServices))
				r.Get("/services/{uid}", withLatestCache(k8sSession, getService))

				r.Get("/networkpolicies", withLatestCache(k8sSession, getNetworkPolicies))
				r.Get("/networkpolicies/{uid}", withLatestCache(k8sSession, getNetworkPolicy))

				r.Get("/endpoints", withLatestCache(k8sSession, getEndpoints))
				r.Get("/endpoints/{uid}", withLatestCache(k8sSession, getEndpoint))

				r.Get("/virtualservices", withLatestCache(k8sSession, getVirtualServices))
				r.Get("/virtualservices/{uid}", withLatestCache(k8sSession, getVirtualService))
			})

			// Storage resources
			r.Route("/storage", func(r chi.Router) {
				r.Get("/persistentvolumes", withLatestCache(k8sSession, getPersistentVolumes))
				r.Get("/persistentvolumes/{uid}", withLatestCache(k8sSession, getPersistentVolume))

				r.Get("/persistentvolumeclaims", withLatestCache(k8sSession, getPersistentVolumeClaims))
				r.Get("/persistentvolumeclaims/{uid}", withLatestCache(k8sSession, getPersistentVolumeClaim))

				r.Get("/storageclasses", withLatestCache(k8sSession, getStorageClasses))
				r.Get("/storageclasses/{uid}", withLatestCache(k8sSession, getStorageClass))
			})
		})
	})

	if config.LocalAuthEnabled {
		port := "8443"
		host := "runtime-local.uds.dev"
		colorYellow := "\033[33m"
		colorReset := "\033[0m"
		url := fmt.Sprintf("https://%s:%s?token=%s", host, port, local.AuthToken)
		log.Printf("%sRuntime API connection: %s%s", colorYellow, url, colorReset)
		err := exec.LaunchURL(url)
		if err != nil {
			return nil, inCluster, fmt.Errorf("failed to launch URL: %w", err)
		}
	}

	// Serve static files from embed.FS
	if assets != nil {
		staticFS, err := fs.Sub(assets, "ui/build")
		if err != nil {
			return nil, inCluster, fmt.Errorf("failed to create static file system: %w", err)
		}

		if err := fileServer(r, http.FS(staticFS)); err != nil {
			return nil, inCluster, fmt.Errorf("failed to serve static files: %w", err)
		}
	}
	return r, inCluster, nil
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
	fsHandler := http.FileServer(root)

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
		fsHandler.ServeHTTP(w, r)
	})

	return nil
}

func Serve(r *chi.Mux, localCert []byte, localKey []byte, inCluster bool) error {
	//nolint:gosec,govet
	if inCluster {
		slog.Info("Starting server in in-cluster mode on 0.0.0.0:8080")

		if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
			message.WarnErrf(err, "server failed to start: %s", err.Error())
			return err
		}
	} else {
		slog.Info("Starting server in local mode on 127.0.0.1:8443")
		// create tls config from embedded cert and key
		cert, err := tls.X509KeyPair(localCert, localKey)
		if err != nil {
			log.Fatalf("Failed to load embedded certificate: %v", err)
		}
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		// Create a server with TLS config
		server := &http.Server{
			Addr:      "127.0.0.1:8443",
			Handler:   r,
			TLSConfig: tlsConfig,
		}

		if err = server.ListenAndServeTLS("", ""); err != nil {
			message.WarnErrf(err, "server failed to start: %s", err.Error())
			return err
		}
	}

	return nil
}

// withLatestCache returns a wrapper lambda function, creating a closure that can dynamically access the latest cache
func withLatestCache(k8sSession *session.K8sSession, handler func(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(k8sSession.Cache)(w, r)
	}
}
