// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

// Write writes the response
func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Flush flushes the response
func (w *gzipResponseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		w.Writer.Flush()
		f.Flush()
	}
}

// ConditionalCompress compresses the response if the client supports it
func ConditionalCompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()

			gzipWriter := &gzipResponseWriter{ResponseWriter: w, Writer: gz}
			next.ServeHTTP(gzipWriter, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
