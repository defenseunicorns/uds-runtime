// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// WriteHeaders sets the headers for an SSE connection
func WriteHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

// Handler is a generic SSE handler that sends data to the client
func Handler(w http.ResponseWriter, r *http.Request, getData func() []unstructured.Unstructured, changes <-chan struct{}) {
	WriteHeaders(w)

	// Ensure the ResponseWriter supports flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	sendData := func() {
		defer flusher.Flush()

		// Convert the data to JSON
		data, err := json.Marshal(getData())
		if err != nil {
			fmt.Fprintf(w, "data: Error: %v\n\n", err)
			return
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
	}

	// Send the initial data
	sendData()

	for {
		select {
		// Check if the client has disconnected
		case <-r.Context().Done():
			return

		// Send data to the client when there are changes
		case <-changes:
			sendData()
		}
	}
}
