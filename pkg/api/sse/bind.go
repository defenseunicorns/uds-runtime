// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"encoding/json"
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Bind is a helper function to bind a cache to an SSE handler
func Bind[T metav1.Object](resource *resources.ResourceList[T]) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		once := r.URL.Query().Get("once")

		// If once is true, send the data once and close the connection
		if once == "true" {
			// Convert the data to JSON
			data, err := json.Marshal(resource.GetResources())
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
		} else {
			Handler(w, r, resource.GetResources, resource.Changes)
		}
	}
}
