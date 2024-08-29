// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

//go:build integration

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	tests := []struct {
		name    string
		url     string
		isDense bool
	}{
		{
			name: "once=true",
			url:  "/api/v1/resources/workloads/pods?once=true",
		},
		{
			name:    "once=true&dense=true",
			url:     "/api/v1/resources/workloads/pods?once=true&dense=true",
			isDense: true,
		},
		{
			name: "sse sparse",
			url:  "/api/v1/resources/workloads/pods",
		},
		{
			name:    "sse dense=true",
			url:     "/api/v1/resources/workloads/pods?dense=true",
			isDense: true,
		},
		{
			name: "sse namespace & name",
			url:  "/api/v1/resources/workloads/pods?namespace=podinfo&name=podinfo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyIndx := 5
			if strings.Contains(tt.url, "once=true") {
				keyIndx = 0
			}

			// Create a new context with a timeout
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.url, nil)

			// Start serving the request for 1 second
			go func(ctx context.Context) {
				r.ServeHTTP(rr, req)
			}(ctx)

			// wait for the context to be done
			<-ctx.Done()
			require.Equal(t, http.StatusOK, rr.Code)

			var data []map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(data), 1)

			// Assert dense versus sparse
			if tt.isDense {
				require.NotNil(t, data[0]["spec"])
			} else {
				require.Nil(t, data[0]["spec"])
			}

			// Assert namespace and name filtering
			if strings.Contains(tt.url, "namespace=podinfo&name=podinfo") {
				require.Equal(t, 1, len(data))
				require.Equal(t, data[0]["metadata"].(map[string]interface{})["namespace"], "podinfo")
				require.Contains(t, data[0]["metadata"].(map[string]interface{})["name"], "podinfo")
			}
		})
	}
}

// TODO: Add case for no ns with UID

func TestFieldSelectors(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	uid := ""

	tests := []struct {
		name    string
		url     string
		isDense bool
	}{
		{
			name: "once=true",
			url:  "/api/v1/resources/workloads/pods?once=true&fields=.metadata.name,.metadata.uid,spec.nodeName",
		},
		{
			name:    "sse",
			url:     "/api/v1/resources/workloads/pods?fields=metadata.name,metadata.uid,spec.nodeName",
			isDense: true,
		},
		{
			name: "sse namespace & name",
			url:  "/api/v1/resources/workloads/pods?namespace=podinfo&name=podinfo&fields=metadata.name,metadata.uid,spec.nodeName",
		},
		{
			name: "uid",
			url:  "/api/v1/resources/workloads/pods/{uid}?fields=metadata.name,metadata.uid,spec.nodeName",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyIndx := 5
			if strings.Contains(tt.url, "once=true") || strings.Contains(tt.url, "{uid}") {
				keyIndx = 0
			}

			// Create a new context with a timeout
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			rr := httptest.NewRecorder()
			tt.url = strings.Replace(tt.url, "{uid}", uid, 1)
			req := httptest.NewRequest("GET", tt.url, nil)

			// Start serving the request for 1 second
			go func(ctx context.Context) {
				r.ServeHTTP(rr, req)
			}(ctx)

			// wait for the context to be done
			<-ctx.Done()

			fmt.Println(rr.Body.String())

			var data []map[string]interface{}
			var resource map[string]interface{}

			if strings.Contains(tt.url, uid) && uid != "" {
				json.Unmarshal(rr.Body.Bytes()[keyIndx:], &resource)
			} else {
				json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)
				resource = data[0]
			}

			// Assert fields selection
			require.Equal(t, 2, len(resource))
			require.NotNil(t, resource["metadata"].(map[string]interface{})["name"])
			require.NotNil(t, resource["spec"].(map[string]interface{})["nodeName"])
			require.NotNil(t, resource["metadata"].(map[string]interface{})["uid"])

			uid = resource["metadata"].(map[string]interface{})["uid"].(string)

			// Assert namespace and name filtering
			if strings.Contains(tt.url, "namespace=podinfo&name=podinfo") {
				require.Equal(t, 1, len(data))
				require.Equal(t, 2, len(resource))
				require.Contains(t, resource["metadata"].(map[string]interface{})["name"], "podinfo")
				require.NotNil(t, resource["spec"].(map[string]interface{})["nodeName"])
				require.NotNil(t, resource["metadata"].(map[string]interface{})["uid"])
			}
		})
	}
}

type TestRoute struct {
	name         string
	url          string
	expectedKind string
}

func TestPodRoutes(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	// Map so can mutate reference between tests
	uidMap := map[string]string{"uid": ""}

	podTests := []TestRoute{
		{
			name:         "pods",
			url:          "/api/v1/resources/workloads/pods",
			expectedKind: "Pod",
		},
		{
			name:         "pods/{uid}",
			url:          "/api/v1/resources/workloads/pods/{uid}",
			expectedKind: "Pod",
		},
	}

	for _, tt := range podTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestPackageRoutes(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	// Map so can mutate reference between tests
	uidMap := map[string]string{"uid": ""}

	packageTests := []TestRoute{
		{
			name:         "uds-packages",
			url:          "/api/v1/resources/configs/uds-packages",
			expectedKind: "Package",
		},
		{
			name:         "uds-packages/{uid}",
			url:          "/api/v1/resources/configs/uds-packages/{uid}",
			expectedKind: "Package",
		},
	}

	for _, tt := range packageTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestPeprRoutes(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	peprTests := []TestRoute{
		{
			name: "uds-packages",
			url:  "/api/v1/monitor/pepr/",
		},
		{
			name: "pepr/{stream}",
			url:  "/api/v1/monitor/pepr/policies",
		},
	}

	for _, tt := range peprTests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new context with a timeout -- increase to 2s to aggregate data
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.url, nil)

			// Start serving the request for 2 second
			go func(ctx context.Context) {
				r.ServeHTTP(rr, req)
			}(ctx)

			// wait for the context to be done
			<-ctx.Done()
			require.Equal(t, http.StatusOK, rr.Code)
			require.NotEmpty(t, rr.Body.String())
			require.Contains(t, rr.Body.String(), "header")
		})
	}
}

func TestClusterOverview(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	t.Run("cluster-overview", func(t *testing.T) {
		// Create a new context with a timeout -- increase to 2s to aggregate data
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/v1/monitor/cluster-overview", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "totalPods")
	})
}

// testRoutesHelper handles logic for testing getResources and getResource routes (e.g. /pods and /pods/{uid})
func testRoutesHelper(t *testing.T, tt TestRoute, uidMap map[string]string, r *chi.Mux) {
	t.Run(tt.name, func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()

		tt.url = strings.Replace(tt.url, "{uid}", uidMap["uid"], 1)
		req := httptest.NewRequest("GET", tt.url, nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusOK, rr.Code)

		// If uidMap is empty, then this is the first test case and we need to store the UID
		if uidMap["uid"] == "" {
			keyIndx := 5
			var data []map[string]interface{}
			json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)
			require.Equal(t, tt.expectedKind, data[0]["kind"])
			uidMap["uid"] = data[0]["metadata"].(map[string]interface{})["uid"].(string)
		} else {
			keyIndx := 0
			var data map[string]interface{}
			json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)
			require.Equal(t, tt.expectedKind, data["Object"].(map[string]interface{})["kind"])
		}
	})
}
