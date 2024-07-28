// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package rest

import (
	"net/http"
	"strings"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/go-chi/chi/v5"
)

// Bind is a helper function to bind a cache to an SSE handler
// It returns an http.HandlerFunc that can be used in a router
func Bind(resource *resources.ResourceList) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		once := r.URL.Query().Get("once") == "true"   // If true, send data only once
		dense := r.URL.Query().Get("dense") == "true" // If true, send full resource data
		fields := r.URL.Query().Get("fields")         // Comma-separated list of fields to include

		var fieldsList []string
		if fields != "" {
			fieldsList = strings.Split(fields, ",")
			// If fields are specified, dense data retrieval is required
			dense = true
		}

		// Get the UID from the URL if it exists
		uid := chi.URLParam(r, "uid")

		// If a UID is provided, send data for that specific resource
		if uid != "" {
			data, found := resource.GetResource(uid)
			if !found {
				http.Error(w, "Resource not found", http.StatusNotFound)
				return
			}
			writeData(w, data, fieldsList)
			return
		}

		// Choose between sparse and dense data retrieval
		getData := resource.GetSparseResources
		if dense {
			getData = resource.GetResources
		}

		// If 'once' is true, send data once and close the connection
		if once {
			writeData(w, getData(), fieldsList)
			return
		}

		// Otherwise, set up an SSE stream
		SSEHandler(w, r, getData, resource.Changes, fieldsList)
	}
}

// writeData writes the payload to the http.ResponseWriter
// It handles field filtering if specific fields are requested
func writeData(w http.ResponseWriter, payload any, fieldsList []string) {
	// Marshal the payload to JSON and filter the fields if specified
	data, err := jsonMarshal(payload, fieldsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Write data to the response
	if _, err := w.Write(data); err != nil {
		http.Error(w, "Failed to write data", http.StatusInternalServerError)
		return
	}
}
