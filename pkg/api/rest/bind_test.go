// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

//go:build unit

package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/test"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestBind(t *testing.T) {
	resourceList := &resources.ResourceList{
		Resources:       make(map[string]*unstructured.Unstructured),
		SparseResources: make(map[string]*unstructured.Unstructured),
	}
	resourceList.Resources["1"] = test.CreateMockPod("mock-pod-1", "uds-dev-stack", "1")
	resourceList.Resources["2"] = test.CreateMockPod("mock-pod-2", "uds-dev-stack", "2")
	resourceList.SparseResources["1"] = test.CreateMockPod("mock-pod-1-sparse", "uds-dev-stack", "1")
	resourceList.SparseResources["2"] = test.CreateMockPod("mock-pod-2-sparse", "uds-dev-stack", "2")

	// Create a new router
	r := chi.NewRouter()
	r.Get("/resources/workloads/pods", Bind(resourceList))
	r.Get("/resources/workloads/pods/{uid}", Bind(resourceList))

	tests := []struct {
		name             string
		url              string
		expectedStatus   int
		expectedResponse []string
	}{
		{
			name:             "Get single resource",
			url:              "/resources/workloads/pods/1",
			expectedStatus:   http.StatusOK,
			expectedResponse: []string{`"Object":{"apiVersion":"v1","kind":"Pod","metadata":{"name":"mock-pod-1","namespace":"uds-dev-stack","uid":"1"}}`},
		},
		{
			name:             "Get non-existent resource",
			url:              "/resources/workloads/pods/3",
			expectedStatus:   http.StatusNotFound,
			expectedResponse: []string{"Resource not found\n"},
		},
		{
			name:           "Get all resources (sparse)",
			url:            "/resources/workloads/pods?once=true",
			expectedStatus: http.StatusOK,
			expectedResponse: []string{
				`{"apiVersion":"v1","kind":"Pod","metadata":{"name":"mock-pod-1-sparse","namespace":"uds-dev-stack","uid":"1"}}`,
				`{"apiVersion":"v1","kind":"Pod","metadata":{"name":"mock-pod-2-sparse","namespace":"uds-dev-stack","uid":"2"}}`,
			},
		},
		{
			name:           "Get all resources (dense)",
			url:            "/resources/workloads/pods?once=true&dense=true",
			expectedStatus: http.StatusOK,
			expectedResponse: []string{
				`{"apiVersion":"v1","kind":"Pod","metadata":{"name":"mock-pod-1","namespace":"uds-dev-stack","uid":"1"}}`,
				`{"apiVersion":"v1","kind":"Pod","metadata":{"name":"mock-pod-2","namespace":"uds-dev-stack","uid":"2"}}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.url, nil)
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code)

			for _, expectedValue := range tt.expectedResponse {
				require.Contains(t, rr.Body.String(), expectedValue)
			}
		})
	}

	// Test SSE functionality
	t.Run("SSE Stream", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/resources/workloads/pods", nil)
		rr := httptest.NewRecorder()

		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()

		// Ensure the response is a text/event-stream
		require.Contains(t, rr.Header().Get("Content-Type"), "text/event-stream")

		// Check if the response starts with "data:"
		require.True(t, len(rr.Body.Bytes()) > 5)
		require.Equal(t, "data:", string(rr.Body.Bytes()[:5]))

		// Parse the JSON data
		var data []map[string]interface{}
		jsonStr := string(rr.Body.Bytes()[5:])
		err := json.Unmarshal([]byte(jsonStr), &data)
		require.NoError(t, err)

		// Check if we received the correct number of resources
		require.Len(t, data, 2)
	})
}

func TestWriteData(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"key": "value"}

	writeData(rr, payload, nil, false)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("writeData returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type
	if contentType := rr.Header().Get("Content-Type"); contentType != "text/json; charset=utf-8" {
		t.Errorf("writeData returned wrong content type: got %v want %v", contentType, "text/json; charset=utf-8")
	}

	// Check the response body
	expected, _ := json.Marshal(payload)
	if rr.Body.String() != string(expected) {
		t.Errorf("writeData returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}
