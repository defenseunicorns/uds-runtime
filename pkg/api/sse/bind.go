// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"encoding/json"
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
)

// Bind is a helper function to bind a cache to an SSE handler
func Bind(resource *resources.ResourceList) func(w http.ResponseWriter, r *http.Request) {
	// Return a function that sends the data to the client
	return func(w http.ResponseWriter, r *http.Request) {
		// By default, send the data as a sparse stream
		once := r.URL.Query().Get("once") == "true"
		dense := r.URL.Query().Get("dense") == "true"

		// Get the data from the cache as sparse by default
		getData := resource.GetSparseResources
		if dense {
			getData = resource.GetResources
		}

		// If once is true, send the data once and close the connection
		if once {
			// Convert the data to JSON
			data, err := json.Marshal(getData())
			if err != nil {
				http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
				return
			}

			// Set the headers
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			// Attempt to compress the data if the client supports it
			supportsGzip := useCompression(w, r)
			if supportsGzip {
				compressedData, err := compressBinaryData(data)
				if err != nil {
					http.Error(w, "Failed to compress data", http.StatusInternalServerError)
					return
				}

				// Replace the data with the compressed data
				data = compressedData
			}

			// Write the data to the response
			if _, err := w.Write(data); err != nil {
				http.Error(w, "Failed to write data", http.StatusInternalServerError)
				return
			}
		} else {
			// Otherwise, send the data as an SSE stream
			Handler(w, r, getData, resource.Changes)
		}
	}
}
