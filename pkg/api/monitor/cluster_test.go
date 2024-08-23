// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package monitor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestBindClusterOverviewHandler(t *testing.T) {
	// Sample data for metrics
	metrics := map[string]*unstructured.Unstructured{
		"pod1": {
			Object: map[string]interface{}{
				"cpu":    "100m",
				"memory": "200Mi",
			},
		},
		"pod2": {
			Object: map[string]interface{}{
				"cpu":    "150m",
				"memory": "250Mi",
			},
		},
	}

	podMetrics := resources.NewPodMetrics()
	podMetrics.Update("pod1", metrics["pod1"])
	podMetrics.Update("pod2", metrics["pod2"])

	// Create ResourceList for nodes
	resourceList := resources.ResourceList{
		Resources:       make(map[string]*unstructured.Unstructured),
		SparseResources: make(map[string]*unstructured.Unstructured),
	}

	// Create a test cache
	cache := &resources.Cache{}
	cache.PodMetrics = podMetrics
	cache.Nodes = &resourceList

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/cluster-overview", nil)
	require.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BindClusterOverviewHandler(cache))

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start serving the request for 1 second
	go func(ctx context.Context) {
		// Call the handler with our request and ResponseRecorder
		handler.ServeHTTP(rr, req)
	}(ctx)

	// wait for the context to be done
	<-ctx.Done()

	// Check the status code is what we expect
	status := rr.Code
	require.Equal(t, http.StatusOK, status)

	// Check if the response body contains expected data
	expectedSubstring := `"totalPods":2`
	require.Contains(t, rr.Body.String(), expectedSubstring)
}
