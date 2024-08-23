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
func Bind(resource *resources.ResourceList) func(w http.ResponseWriter, r *http.Request) {
	// Return a function that sends the data to the client
	return func(w http.ResponseWriter, r *http.Request) {
		// By default, send the data as a sparse stream
		once := r.URL.Query().Get("once") == "true"
		dense := r.URL.Query().Get("dense") == "true"
		namespace := r.URL.Query().Get("namespace")
		namePartial := r.URL.Query().Get("name")
		fields := r.URL.Query().Get("fields") // Comma-separated list of fields to include

		var fieldsList []string
		if fields != "" {
			fieldsList = strings.Split(fields, ",")
			// If fields are specified, dense data retrieval is required
			dense = true
		}

		// Get the UID from the URL if it exists
		uid := chi.URLParam(r, "uid")

		// Get the data from the cache as sparse by default
		getData := resource.GetSparseResources
		if dense {
			getData = resource.GetResources
		}

		// If a UID is provided, send the data for that UID
		// Streaming is not supported for single resources
		if uid != "" {
			// If a namespace is provided, return a 400
			if namespace != "" {
				http.Error(w, "Namespace and UID cannot be used together", http.StatusBadRequest)
				return
			}

			data, found := resource.GetResource(uid)
			// If the resource is not found, return a 404
			if !found {
				http.Error(w, "Resource not found", http.StatusNotFound)
				return
			}

			// Otherwise, write the data to the client
			writeData(w, data, fieldsList)
			return
		}

		// If once is true, send the list data once and close the connection
		if once {
			writeData(w, getData(namespace, namePartial), fieldsList)
			return
		}

		// Otherwise, send the data as an SSE stream
		Handler(w, r, getData, resource.Changes, fieldsList)
	}
}

// writeData writes the payload to the http.ResponseWriter
// It handles field filtering if specific fields are requested
func writeData(w http.ResponseWriter, payload any, fieldsList []string) {
	// Marshal the payload to JSON and filter the fields if specified
	data, err := jsonMarshal(payload, fieldsList)
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
