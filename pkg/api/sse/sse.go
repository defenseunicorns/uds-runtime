// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

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

	// Track the last sent time
	var lastSent time.Time
	// Use a mutex to prevent concurrent access to the last sent time and pending flag
	var mu sync.Mutex
	// Track if there is a pending update
	var pendingUpdate bool
	// Set the debounce interval
	debounceInterval := time.Second

	// Function to send the data
	sendData := func(immediate bool) {
		// Lock the mutex to prevent concurrent access
		mu.Lock()
		defer mu.Unlock()

		// Check if within the debounce interval and set the pending flag if so
		now := time.Now()
		if !immediate && now.Sub(lastSent) < debounceInterval {
			pendingUpdate = true
			return
		}

		// Flush the headers at the end
		defer flusher.Flush()

		// Convert the data to JSON
		data, err := json.Marshal(getData())
		if err != nil {
			fmt.Fprintf(w, "data: Error: %v\n\n", err)
			return
		}

		// Write the data to the response
		fmt.Fprintf(w, "data: %s\n\n", data)

		// Update the last sent time and reset the pending flag
		lastSent = now
		pendingUpdate = false
	}

	// Send the initial data
	sendData(true)

	// Setup a ticker to check for pending updates
	ticker := time.NewTicker(debounceInterval)
	defer ticker.Stop()

	for {
		select {
		// If the context is done, return
		case <-r.Context().Done():
			return

		// If there is a change, send the data
		case <-changes:
			sendData(false)

		// If there is a pending update, send the data immediately
		case <-ticker.C:
			mu.Lock()
			if pendingUpdate {
				mu.Unlock()
				sendData(true)
			} else {
				mu.Unlock()
			}
		}
	}
}
