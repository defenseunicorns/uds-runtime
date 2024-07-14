// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteHeaders sets the headers for an SSE connection
func WriteHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

// Handler is a generic SSE handler that sends data to the client
func Handler[T any](w http.ResponseWriter, r *http.Request, getData func() T, changes <-chan struct{}) {
	WriteHeaders(w)

	// Ensure the ResponseWriter supports flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	sendData := func() {
		data, err := json.Marshal(getData())
		if err != nil {
			fmt.Fprintf(w, "data: Error: %v\n\n", err)
			flusher.Flush()
			return
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
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
