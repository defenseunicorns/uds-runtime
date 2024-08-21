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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

			// Assert dense versus sparse
			if tt.isDense {
				require.NotNil(t, data[0]["spec"].(map[string]interface{})["containers"])
			} else {
				require.Nil(t, data[0]["spec"].(map[string]interface{})["containers"])
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

type TestRoute struct {
	name         string
	url          string
	expectedKind runtime.Object
}

func TestRoutes(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	// Map so can mutate reference between tests
	uidMap := map[string]string{"uid": ""}

	podTests := []TestRoute{
		{
			name:         "pods",
			url:          "/api/v1/resources/workloads/pods",
			expectedKind: &v1.Pod{},
		},
		{
			name:         "pods/{uid}",
			url:          "/api/v1/resources/workloads/pods/{uid}",
			expectedKind: &v1.Pod{},
		},
	}

	for _, tt := range podTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func testRoutesHelper(t *testing.T, tt TestRoute, uidMap map[string]string, r *chi.Mux) {
	fmt.Println("uid " + uidMap["uid"])
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

		if uidMap["uid"] != "" {
			keyIndx := 0
			var data json.RawMessage
			err := json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)
			require.NoError(t, err)

			// Unmarshal the data into the expected kind
			err = json.Unmarshal(data, tt.expectedKind)
			require.NoError(t, err)
		} else {
			keyIndx := 5
			var data []json.RawMessage
			json.Unmarshal(rr.Body.Bytes()[keyIndx:], &data)

			// Unmarshal the first entry into the expected kind
			err := json.Unmarshal(data[0], tt.expectedKind)
			require.NoError(t, err)

			// Get the UID from the first entry for the next test
			var dataStruct []map[string]interface{}
			json.Unmarshal(rr.Body.Bytes()[keyIndx:], &dataStruct)
			uidMap["uid"] = dataStruct[0]["metadata"].(map[string]interface{})["uid"].(string)
		}
	})
}
