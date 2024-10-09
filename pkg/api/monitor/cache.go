// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

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

type Cache struct {
	buffer  *bytes.Buffer
	lock    sync.RWMutex
	timer   *time.Timer
	maxSize int
}

func NewCache() *Cache {
	c := &Cache{
		buffer:  &bytes.Buffer{},
		maxSize: 1024 * 1024 * 10, // 10MB
	}
	c.startResetTimer()
	return c
}

func (c *Cache) startResetTimer() {
	c.timer = time.AfterFunc(5*time.Minute, func() {
		message.Debug("Cache invalidated by timer, resetting cache")
		c.Reset()
		c.startResetTimer() // restart timer
	})
}

func (c *Cache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()
	message.Debug("Resetting cache")
	c.buffer.Reset()
}

func (c *Cache) Stop() {
	if c.timer != nil {
		message.Debugf("Stopping cache timer")
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
		message.Debugf("Buffer size %d exceeds max size %d, resetting cache", buffer.Len(), c.maxSize)
		c.Reset()
		return
	}
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
