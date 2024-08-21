// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package udsmiddleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
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

			encoding := resp.Header.Get("Content-Encoding")
			require.Equal(t, encoding, tt.expectedEncoding)

			var body []byte
			var err error
			if tt.expectedEncoding == "gzip" {
				reader, err := gzip.NewReader(resp.Body)
				require.NoError(t, err)
				defer reader.Close()

				body, err = io.ReadAll(reader)
				require.NoError(t, err)
			} else {
				body, err = io.ReadAll(resp.Body)
				require.NoError(t, err)
			}

			require.Equal(t, tt.expectedBody, string(body))
		})
	}
}

func TestGzipResponseWriter(t *testing.T) {
	w := httptest.NewRecorder()
	gzw := gzip.NewWriter(w)
	grw := &gzipResponseWriter{ResponseWriter: w, Writer: gzw}

	testContent := "Hello, Gzip!"
	n, err := grw.Write([]byte(testContent))
	require.NoError(t, err)
	require.Equal(t, len(testContent), n)

	grw.Flush()
	gzw.Close()

	reader, err := gzip.NewReader(w.Body)
	require.NoError(t, err)
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	require.NoError(t, err)

	require.Equal(t, testContent, string(decompressed))
}
