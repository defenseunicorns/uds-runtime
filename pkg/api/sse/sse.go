// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package sse

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	// Check if the client accepts gzip encoding
	supportsGzip := useCompression(w, r)

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

		// If the client supports gzip, compress the data
		if supportsGzip {
			compressedData, err := compressData("data: %s\n\n", data)
			if err != nil {
				http.Error(w, "Error compressing data", http.StatusInternalServerError)
				return
			}

			// Write the compressed data to the client
			if _, err := w.Write(compressedData); err != nil {
				http.Error(w, "Failed to write data", http.StatusInternalServerError)
				return
			}
		} else {
			// Otherwise, write the uncompressed data to the client
			fmt.Fprintf(w, "data: %s\n\n", data)
		}
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

// useCompression checks if the client supports gzip encoding and sets the appropriate headers
func useCompression(w http.ResponseWriter, r *http.Request) bool {
	// Check if the client accepts gzip encoding
	acceptEncoding := r.Header.Get("Accept-Encoding")
	supportsGzip := strings.Contains(acceptEncoding, "gzip")

	// Set the Content-Encoding header if the client supports gzip
	if supportsGzip {
		w.Header().Set("Content-Encoding", "gzip")
	}

	// Return whether the client supports gzip encoding
	return supportsGzip
}

// compressData compresses data using gzip
func compressData(format string, a ...any) ([]byte, error) {
	// Format the data
	data := fmt.Sprintf(format, a...)
	// Compress the data
	return compressBinaryData([]byte(data))
}

// compressBinaryData compresses binary data using gzip
func compressBinaryData(data []byte) ([]byte, error) {
	var b bytes.Buffer

	// Create a new gzip writer
	gz := gzip.NewWriter(&b)

	// Write the data to the gzip writer
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	// Close the gzip writer
	if err := gz.Close(); err != nil {
		return nil, err
	}

	// Return the compressed data
	return b.Bytes(), nil
}
