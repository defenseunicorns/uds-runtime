// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package monitor

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

// single instance of the pepr stream cache
var streamCache = NewCache()

type Cache struct {
	buffer *bytes.Buffer
	lock   sync.RWMutex
}

// NewCache creates a new Cache instance for caching pepr stream responses
func NewCache() *Cache {
	return &Cache{
		buffer: nil,
	}
}

// Get returns the cached buffer
func (c *Cache) Get() *bytes.Buffer {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.buffer
}

// Set sets the cached buffer
func (c *Cache) Set(buffer *bytes.Buffer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.buffer = buffer
}

// Serve attempts to serve a cached response if available.
// It returns true if a cached response was served, false otherwise.
func (c *Cache) Serve(w http.ResponseWriter) bool {
	cachedBuffer := c.Get()
	if cachedBuffer == nil || cachedBuffer.Len() == 0 {
		return false
	}

	rest.WriteHeaders(w)
	_, err := w.Write(cachedBuffer.Bytes())
	if err != nil {
		message.Warnf("Failed to write response: %v", err)
		return false
	}
	w.(http.Flusher).Flush()
	message.Debug("Used cached pepr stream buffer")
	return true
}
