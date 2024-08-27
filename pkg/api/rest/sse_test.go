// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package rest

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/test"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestWriteHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	WriteHeaders(rr)

	require.Equal(t, "text/event-stream; charset=utf-8", rr.Header().Get("Content-Type"))
	require.Equal(t, "no-cache", rr.Header().Get("Cache-Control"))
	require.Equal(t, "keep-alive", rr.Header().Get("Connection"))
}

func TestHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	//mock pod data
	mockPodData := test.CreateMockPod("mock-pod-1", "namespace1", "1")

	// Mock getData function
	getData := func(string, string) []unstructured.Unstructured {
		return []unstructured.Unstructured{*mockPodData}
	}

	// Mock changes channel
	changes := make(chan struct{})
	defer close(changes)

	// Run the handler in a separate goroutine to simulate real-world usage
	go func() {
		Handler(rr, req, getData, changes, nil)
	}()

	// Simulate a change
	time.Sleep(100 * time.Millisecond)
	changes <- struct{}{}

	// Allow some time for the handler to process the change
	time.Sleep(100 * time.Millisecond)

	// Verify the response
	require.Equal(t, "text/event-stream; charset=utf-8", rr.Header().Get("Content-Type"))
	require.Equal(t, "no-cache", rr.Header().Get("Cache-Control"))
	require.Equal(t, "keep-alive", rr.Header().Get("Connection"))

	// Verify the response body
	expectedData := []unstructured.Unstructured{*mockPodData}

	expectedJSON, _ := json.Marshal(expectedData)
	expectedBody := "data: " + string(expectedJSON) + "\n\n"
	require.Contains(t, rr.Body.String(), expectedBody)
}
