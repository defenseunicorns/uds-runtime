// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"encoding/json"
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/go-chi/chi/v5"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Option is a type for functional options for the Bind function
type Option func(*bindOptions)

type bindOptions struct {
	customDataHandler func() []unstructured.Unstructured
}

// WithCustomDataHandler sets a custom data handler
func WithCustomDataHandler(handler func() []unstructured.Unstructured) Option {
	return func(bo *bindOptions) {
		bo.customDataHandler = handler
	}
}

// Bind is a helper function to bind a cache to an SSE handler
func Bind(resource *resources.ResourceList, opts ...Option) func(w http.ResponseWriter, r *http.Request) {
	options := &bindOptions{}
	for _, opt := range opts {
		opt(options)
	}
	// Return a function that sends the data to the client
	return func(w http.ResponseWriter, r *http.Request) {
		// By default, send the data as a sparse stream
		once := r.URL.Query().Get("once") == "true"
		dense := r.URL.Query().Get("dense") == "true"

		// Get the UID from the URL if it exists
		uid := chi.URLParam(r, "uid")

		// Get the data from the cache as sparse by default
		getData := resource.GetSparseResources
		if dense {
			getData = resource.GetResources
		}

		// use custom data handler if provided
		if options.customDataHandler != nil {
			getData = options.customDataHandler
		}

		// If a UID is provided, send the data for that UID
		// Streaming is not supported for single resources
		if uid != "" {
			data, found := resource.GetResource(uid)
			// If the resource is not found, return a 404
			if !found {
				http.Error(w, "Resource not found", http.StatusNotFound)
				return
			}

			// Otherwise, write the data to the client
			writeData(w, data)
			return
		}

		// If once is true, send the list data once and close the connection
		if once {
			writeData(w, getData())
			return
		}

		// Otherwise, send the data as an SSE stream
		Handler(w, r, getData, resource.Changes)
	}
}

func writeData(w http.ResponseWriter, payload any) {
	// Convert the data to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set the headers
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Write the data to the response
	if _, err := w.Write(data); err != nil {
		http.Error(w, "Failed to write data", http.StatusInternalServerError)
		return
	}
}
