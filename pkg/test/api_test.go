package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		isDense        bool
	}{
		{
			name:           "once=true",
			url:            "/api/v1/resources/workloads/pods?once=true",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "sse sparse",
			url:            "/api/v1/resources/workloads/pods",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "sse dense=true",
			url:            "/api/v1/resources/workloads/pods?dense=true",
			expectedStatus: http.StatusOK,
			isDense:        true,
		},
		{
			name:           "sse namespace & name",
			url:            "/api/v1/resources/workloads/pods?namespace=podinfo&name=podinfo",
			expectedStatus: http.StatusOK,
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
			require.Equal(t, tt.expectedStatus, rr.Code)

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

// func TestRoutes(t *testing.T) {
// 	r, err := api.Setup(nil)
// 	require.NoError(t, err)

// 	tests := []struct {
// 		name             string
// 		url              string
// 		expectedStatus   int
// 		expectedKind     runtime.Object
// 		expectedResponse []string
// 	}{
// 		{
// 			name:           "once=true",
// 			url:            "/api/v1/resources/workloads/pods?once=true",
// 			expectedStatus: http.StatusOK,
// 			expectedKind:   &v1.Pod{},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a new context with a timeout
// 			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 			defer cancel()

// 			// Start serving the request for 1 second
// 			var testServer *httptest.Server
// 			go func(ctx context.Context) {
// 				testServer = httptest.NewServer(r)
// 			}(ctx)

// 			// wait for the context to be done
// 			<-ctx.Done()

// 			resp, err := testServer.Client().Get(testServer.URL + tt.url)
// 			require.NoError(t, err)
// 			require.Equal(t, 200, resp.StatusCode)

// 			body, err := io.ReadAll(resp.Body) // Read the response body
// 			require.NoError(t, err)
// 			defer resp.Body.Close()

// 			var responseArray []json.RawMessage
// 			err = json.Unmarshal(body, &responseArray) // Unmarshal the response body into an array of json.RawMessage
// 			require.NoError(t, err)
// 			require.NotEmpty(t, responseArray)

// 			err = json.Unmarshal(responseArray[0], tt.expectedKind) // Unmarshal the first entry into the expected kind
// 			require.NoError(t, err)
// 		})
// 	}

// }
