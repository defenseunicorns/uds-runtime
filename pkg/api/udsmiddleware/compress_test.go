// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package udsmiddleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConditionalCompress(t *testing.T) {
	tests := []struct {
		name             string
		acceptEncoding   string
		expectedEncoding string
		expectedBody     string
	}{
		{
			name:             "With gzip encoding",
			acceptEncoding:   "gzip",
			expectedEncoding: "gzip",
			expectedBody:     "Hello, World!",
		},
		{
			name:             "Without gzip encoding",
			acceptEncoding:   "",
			expectedEncoding: "",
			expectedBody:     "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := ConditionalCompress(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.expectedBody))
			}))

			req := httptest.NewRequest("GET", "/", nil)
			if tt.acceptEncoding != "" {
				req.Header.Set("Accept-Encoding", tt.acceptEncoding)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if encoding := resp.Header.Get("Content-Encoding"); encoding != tt.expectedEncoding {
				t.Errorf("Expected Content-Encoding %q, got %q", tt.expectedEncoding, encoding)
			}

			var body []byte
			var err error
			if tt.expectedEncoding == "gzip" {
				reader, err := gzip.NewReader(resp.Body)
				if err != nil {
					t.Fatalf("Failed to create gzip reader: %v", err)
				}
				defer reader.Close()
				body, err = io.ReadAll(reader)
				if err != nil {
					t.Fatalf("Failed to read gzipped body: %v", err)
				}
			} else {
				body, err = io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("Failed to read response body: %v", err)
				}
			}

			if string(body) != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}

func TestGzipResponseWriter(t *testing.T) {
	w := httptest.NewRecorder()
	gzw := gzip.NewWriter(w)
	grw := &gzipResponseWriter{ResponseWriter: w, Writer: gzw}

	testContent := "Hello, Gzip!"
	n, err := grw.Write([]byte(testContent))
	if err != nil {
		t.Fatalf("Failed to write to gzipResponseWriter: %v", err)
	}
	if n != len(testContent) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(testContent), n)
	}

	grw.Flush()
	gzw.Close()

	reader, err := gzip.NewReader(w.Body)
	if err != nil {
		t.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("Failed to read decompressed content: %v", err)
	}

	if string(decompressed) != testContent {
		t.Errorf("Expected decompressed content %q, got %q", testContent, string(decompressed))
	}
}
