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
	"github.com/go-chi/chi/v5"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

// @Description Get Pepr data
// @Tags monitor
// @Accept  html
// @Produce  text/event-stream
// @Success 200
// @Router /monitor/pepr/{stream} [get]
// @Param stream path string false "stream type to filter on, all streams by default" Enums(AnyStream, PolicyStream, OperatorStream, AllowStream, DenyStream, MutateStream, FailureStream)
func Pepr(w http.ResponseWriter, r *http.Request) {
	streamFilter := chi.URLParam(r, "stream")

	if !pepr.IsValidStreamFilter(pepr.StreamKind(streamFilter)) {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	// Only use cache for the default stream (empty streamFilter)
	if streamFilter == "" {
		streamCache.ServeCachedResponse(w)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set the headers for streaming
	rest.WriteHeaders(w)

	// Create a new BufferWriter
	bufferWriter := newBufferWriter(w)

	// pass context to stream reader to clean up spawned goroutines that watch pepr pods
	peprReader := pepr.NewStreamReader("", "")
	peprReader.JSON = true
	peprReader.FilterStream = pepr.StreamKind(streamFilter)

	peprStream := stream.NewStream(bufferWriter, peprReader, "pepr-system")
	peprStream.Follow = true
	peprStream.Timestamps = true

	// Start the stream in a goroutine
	message.Debug("Starting parent pepr stream goroutine")
	//nolint:errcheck
	go peprStream.Start(ctx)

	// Create a timer to send keep-alive messages
	keepAliveTimer := time.NewTimer(2 * time.Second)
	defer keepAliveTimer.Stop()

	// Create a ticker to flush the buffer
	flushTicker := time.NewTicker(time.Second)
	defer flushTicker.Stop()

	// create tmp cached buffer to hold stream data
	tmpCacheBuffer := &bytes.Buffer{}

	for {
		select {
		// Check if the client has disconnected
		case <-r.Context().Done():
			message.Info("Client disconnected")
			return

		// Handle keep-alive messages
		case <-keepAliveTimer.C:
			keepAliveTimer.Reset(30 * time.Second)
			bufferWriter.KeepAlive()

		// Flush every second if there is data
		case <-flushTicker.C:
			if bufferWriter.buffer.Len() > 0 {
				// Write to both the response and the cache buffer
				data := bufferWriter.buffer.Bytes()

				if err := bufferWriter.Flush(w); err != nil {
					message.WarnErr(err, "Failed to flush buffer")
					return
				}

				// Update the cached buffer if on default stream
				if streamFilter == "" {
					tmpCacheBuffer.Write(data)
					streamCache.Set(tmpCacheBuffer)
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
	_, err := fmt.Fprintf(bw.buffer, ": \n\n")
	if err != nil {
		message.Warnf("Failed to write keep-alive message: %v", err)
	}
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
