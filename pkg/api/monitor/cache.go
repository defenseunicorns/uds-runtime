// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package monitor

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

// single instance of the pepr stream cache
var streamCache = NewCache()

// Cache is a simple cache for pepr stream data, it can be invalidated by a timer and max size
type Cache struct {
	buffer  *bytes.Buffer
	lock    sync.RWMutex
	timer   *time.Timer
	maxSize int
}

// NewCache creates a new cache for pepr stream data
func NewCache() *Cache {
	c := &Cache{
		buffer:  &bytes.Buffer{},
		maxSize: 1024 * 1024 * 10, // 10MB
	}
	c.startResetTimer()
	return c
}

// startResetTimer starts a timer that resets the cache after 5 minutes
func (c *Cache) startResetTimer() {
	c.timer = time.AfterFunc(5*time.Minute, func() {
		message.Debug("Pepr cache invalidated by timer, resetting cache")
		c.Reset()
		c.startResetTimer() // restart timer
	})
}

// Reset resets the cache
func (c *Cache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()
	message.Debug("Resetting pepr cache")
	c.buffer.Reset()
}

// Stop stops the cache timer
func (c *Cache) Stop() {
	if c.timer != nil {
		message.Debugf("Stopping pepr cache timer")
		c.timer.Stop()
	}
}

// Get returns a deep copy of cached buffer
func (c *Cache) Get() *bytes.Buffer {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.buffer == nil {
		return nil
	}

	return bytes.NewBuffer(c.buffer.Bytes())
}

// Set sets the cached buffer
func (c *Cache) Set(buffer *bytes.Buffer) {
	if buffer.Len() > c.maxSize {
		message.Debugf("Pepr cache size %d exceeds max size %d, resetting", buffer.Len(), c.maxSize)
		c.Reset()
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.buffer = buffer
}

// Serve attempts to serve a cached response if available.
func (c *Cache) Serve(w http.ResponseWriter) {
	cachedBuffer := c.Get()
	if cachedBuffer == nil || cachedBuffer.Len() == 0 {
		return
	}

	rest.WriteHeaders(w)
	_, err := w.Write(cachedBuffer.Bytes())
	if err != nil {
		message.Warnf("Pepr cache failed to write response: %v", err)
		return
	}
	w.(http.Flusher).Flush()
	message.Debug("Used pepr cache to serve response")
}
