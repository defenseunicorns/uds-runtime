// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// Package monitor provides mechanisms for interacting with Pepr data streams
package monitor

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

type Cache struct {
	buffer *bytes.Buffer
	lock   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		buffer: nil,
	}
}

func (c *Cache) Get() *bytes.Buffer {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.buffer
}

func (c *Cache) Set(buffer *bytes.Buffer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.buffer = buffer
}

// ServeCachedResponse attempts to serve a cached response if available.
// It returns true if a cached response was served, false otherwise.
func (c *Cache) ServeCachedResponse(w http.ResponseWriter) bool {
	cachedBuffer := c.Get()
	if cachedBuffer == nil {
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
