// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package monitor

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/defenseunicorns/uds-runtime/pkg/pepr"
	"github.com/defenseunicorns/uds-runtime/pkg/stream"
	"github.com/defenseunicorns/zarf/src/pkg/message"
	"github.com/go-chi/chi/v5"
)

func Pepr(w http.ResponseWriter, r *http.Request) {
	streamFilter := chi.URLParam(r, "stream")

	if !pepr.IsValidStreamFilter(pepr.StreamKind(streamFilter)) {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set the headers for streaming
	rest.WriteSSEHeaders(w)

	// Create a new BufferWriter
	bufferWriter := newBufferWriter(w)

	// pass context to stream reader to clean up spawned goroutines that watch pepr pods
	peprReader := pepr.NewStreamReader("", "")
	peprReader.JSON = true
	peprReader.FilterStream = pepr.StreamKind(streamFilter)

	peprStream := stream.NewStream(bufferWriter, peprReader, "pepr-system")
	peprStream.Follow = true
	peprStream.Timestamps = true

	//nolint:errcheck
	// Start the stream in a goroutine
	go peprStream.Start(ctx)

	// Track if the first message has been seen
	seen := false

	// Create a timer to send keep-alive messages
	// The first message is sent after 2 seconds to detect empty streams
	keepAliveTimer := time.NewTimer(2 * time.Second)
	defer keepAliveTimer.Stop()

	// Create a ticker to flush the buffer
	flushTicker := time.NewTicker(time.Second)
	defer flushTicker.Stop()

	for {
		select {
		// Check if the client has disconnected
		case <-r.Context().Done():
			message.Info("Client disconnected")
			return

		// Handle keep-alive messages
		// See https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#examples
		case <-keepAliveTimer.C:
			// Set the keep-alive duration to 30 seconds after the first message
			if !seen {
				keepAliveTimer.Reset(30 * time.Second)
			}

			bufferWriter.KeepAlive()

		// Flush every second if there is data
		case <-flushTicker.C:
			if bufferWriter.buffer.Len() > 0 {
				if err := bufferWriter.Flush(w); err != nil {
					message.WarnErr(err, "Failed to flush buffer")
					return
				}
			}
		}
	}
}

// bufferWriter is a custom writer that aggregates data and writes it to an http.ResponseWriter
type bufferWriter struct {
	buffer  *bytes.Buffer
	mutex   sync.Mutex
	flusher http.Flusher
}

// newBufferWriter creates a new BufferWriter
func newBufferWriter(w http.ResponseWriter) *bufferWriter {
	// Ensure the ResponseWriter also implements http.Flusher
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("ResponseWriter does not implement http.Flusher")
	}
	return &bufferWriter{
		buffer:  new(bytes.Buffer),
		flusher: flusher,
	}
}

func (bw *bufferWriter) KeepAlive() {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()
	fmt.Fprintf(bw.buffer, ": \n\n")
}

// Write writes data to the buffer
func (bw *bufferWriter) Write(p []byte) (n int, err error) {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()

	event := fmt.Sprintf("data: %s\n\n", p)

	// Write data to the buffer
	return bw.buffer.WriteString(event)
}

// Flush writes the buffer content to the http.ResponseWriter and flushes it
func (bw *bufferWriter) Flush(w http.ResponseWriter) error {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()

	_, err := w.Write(bw.buffer.Bytes())
	if err != nil {
		return err
	}

	// Clear the buffer
	bw.buffer.Reset()

	// Flush the response
	bw.flusher.Flush()
	return nil
}
