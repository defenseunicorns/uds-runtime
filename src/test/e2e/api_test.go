// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

type TestRoute struct {
	name         string
	url          string
	expectedKind string
}

func setup() (*chi.Mux, error) {
	os.Setenv("LOCAL_AUTH_ENABLED", "false")
	r, _, err := api.Setup(nil)
	return r, err
}

func teardown() {
	os.Setenv("LOCAL_AUTH_ENABLED", "true")
}

func TestQueryParams(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

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
			err = json.Unmarshal(processResponseBody(rr.Body.String())[keyIndx:], &data)
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

func TestFieldSelectors(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	uid := ""

	tests := []struct {
		name string
		url  string
	}{
		{
			name: "once=true",
			url:  "/api/v1/resources/workloads/pods?once=true&fields=.metadata.name,.metadata.uid,spec.nodeName",
		},
		{
			name: "sse",
			url:  "/api/v1/resources/workloads/pods?fields=metadata.name,metadata.uid,spec.nodeName",
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
			// Replace the UID in the URL if it exists; depends on at least 1 previous test case
			tt.url = strings.Replace(tt.url, "{uid}", uid, 1)
			req := httptest.NewRequest("GET", tt.url, nil)

			// Start serving the request for 1 second
			go func(ctx context.Context) {
				r.ServeHTTP(rr, req)
			}(ctx)

			// wait for the context to be done
			<-ctx.Done()

			var data []map[string]interface{}
			var resource map[string]interface{}

			body := processResponseBody(rr.Body.String())

			if strings.Contains(tt.url, uid) && uid != "" {
				json.Unmarshal(body[keyIndx:], &resource)
			} else {
				json.Unmarshal(body[keyIndx:], &data)
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

func TestUIDFail(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	t.Run("no_uid_with_query_params", func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/v1/resources/workloads/pods/123?namespace=podinfo&name=podinfo", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestPeprRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

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
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

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

// processResponseBody removes potential second occurrence of "data: [...]" from the response body when making SSE calls
// these occurrences are intermittent and happen most likely because of network latency allowing a second sendData() event
func processResponseBody(body string) []byte {
	// Check if "data:" exists in the body
	if !strings.Contains(body, "data:") {
		// If no "data:" is detected, return the body as is
		return []byte(body)
	}

	// Split the body by newlines
	lines := strings.Split(body, "\n")

	// Check if data: [] is empty
	emptyData := regexp.MustCompile(`data:\s*\[\s*\]`)
	if emptyData.MatchString(lines[0]) {
		return []byte{}
	}
	// return the first data:[] event
	return []byte(lines[0])
}

func TestClusterHealth(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	t.Run("healthz", func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/healthz", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "\"status\":\"UP\"")
	})

	t.Run("cluster connected", func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/cluster-check", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "success")
	})
}

func TestGetClusters(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	t.Run("cluster connected", func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/v1/clusters", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()
		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "k3d-uds")
	})
}

func TestSwagger(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	t.Run("swagger", func(t *testing.T) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/swagger", nil)

		// Start serving the request for 1 second
		go func(ctx context.Context) {
			r.ServeHTTP(rr, req)
		}(ctx)

		// wait for the context to be done
		<-ctx.Done()

		// follow redirect
		require.Equal(t, http.StatusMovedPermanently, rr.Code)
		req = httptest.NewRequest("GET", rr.Header().Get("Location"), nil)
		rr = httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "Swagger")
	})
}

func TestTopLevelResourceRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	uidMap := map[string]string{}

	resource := []TestRoute{
		// Nodes
		{
			name:         "nodes",
			url:          "/api/v1/resources/nodes",
			expectedKind: "Node",
		},
		{
			name:         "nodes/{uid}",
			url:          "/api/v1/resources/nodes/{uid}",
			expectedKind: "Node",
		},

		// Events
		{
			name:         "events",
			url:          "/api/v1/resources/events",
			expectedKind: "Event",
		},
		{
			name:         "events/{uid}",
			url:          "/api/v1/resources/events/{uid}",
			expectedKind: "Event",
		},

		// Namespaces
		{
			name:         "namespaces",
			url:          "/api/v1/resources/namespaces",
			expectedKind: "Namespace",
		},
		{
			name:         "namespaces/{uid}",
			url:          "/api/v1/resources/namespaces/{uid}",
			expectedKind: "Namespace",
		},
		// Custom Resource Definitions
		{
			name:         "crds",
			url:          "/api/v1/resources/custom-resource-definitions",
			expectedKind: "CustomResourceDefinition",
		},
		{
			name:         "crds/{uid}",
			url:          "/api/v1/resources/custom-resource-definitions/{uid}",
			expectedKind: "CustomResourceDefinition",
		},
	}

	for _, tt := range resource {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestWorkloadRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	uidMap := map[string]string{}

	workloadTests := []TestRoute{
		// Workloads - Pods
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

		// Workloads - Deployments
		{
			name:         "deployments",
			url:          "/api/v1/resources/workloads/deployments",
			expectedKind: "Deployment",
		},
		{
			name:         "deployments/{uid}",
			url:          "/api/v1/resources/workloads/deployments/{uid}",
			expectedKind: "Deployment",
		},

		// Workloads - Daemonsets
		{
			name:         "daemonsets",
			url:          "/api/v1/resources/workloads/daemonsets",
			expectedKind: "DaemonSet",
		},
		{
			name:         "daemonsets/{uid}",
			url:          "/api/v1/resources/workloads/daemonsets/{uid}",
			expectedKind: "DaemonSet",
		},

		// Workloads - Statefulsets
		{
			name:         "statefulsets",
			url:          "/api/v1/resources/workloads/statefulsets",
			expectedKind: "StatefulSet",
		},
		{
			name:         "statefulsets/{uid}",
			url:          "/api/v1/resources/workloads/statefulsets/{uid}",
			expectedKind: "StatefulSet",
		},

		// Workloads - Jobs
		{
			name:         "jobs",
			url:          "/api/v1/resources/workloads/jobs",
			expectedKind: "Job",
		},
		{
			name:         "jobs/{uid}",
			url:          "/api/v1/resources/workloads/jobs/{uid}",
			expectedKind: "Job",
		},

		// Workloads - Cronjobs
		{
			name:         "cronjobs",
			url:          "/api/v1/resources/workloads/cronjobs",
			expectedKind: "CronJob",
		},
		{
			name:         "cronjobs/{uid}",
			url:          "/api/v1/resources/workloads/cronjobs/{uid}",
			expectedKind: "CronJob",
		},
	}

	for _, tt := range workloadTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestConfigRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	uidMap := map[string]string{}

	configTests := []TestRoute{
		// Configs - UDS Packages
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

		// Configs - UDS Exemptions
		{
			name:         "uds-exemptions",
			url:          "/api/v1/resources/configs/uds-exemptions",
			expectedKind: "Exemption",
		},
		{
			name:         "uds-exemptions/{uid}",
			url:          "/api/v1/resources/configs/uds-exemptions/{uid}",
			expectedKind: "Exemption",
		},

		// Configs - ConfigMaps
		{
			name:         "configmaps",
			url:          "/api/v1/resources/configs/configmaps",
			expectedKind: "ConfigMap",
		},
		{
			name:         "configmaps/{uid}",
			url:          "/api/v1/resources/configs/configmaps/{uid}",
			expectedKind: "ConfigMap",
		},

		// Configs - Secrets
		{
			name:         "secrets",
			url:          "/api/v1/resources/configs/secrets",
			expectedKind: "Secret",
		},
		{
			name:         "secrets/{uid}",
			url:          "/api/v1/resources/configs/secrets/{uid}",
			expectedKind: "Secret",
		},
	}

	for _, tt := range configTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestNetworkAndStorageRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	uidMap := map[string]string{}

	networkTests := []TestRoute{
		// Network - Services
		{
			name:         "services",
			url:          "/api/v1/resources/networks/services",
			expectedKind: "Service",
		},
		{
			name:         "services/{uid}",
			url:          "/api/v1/resources/networks/services/{uid}",
			expectedKind: "Service",
		},
	}

	storageTests := []TestRoute{
		// Storage - PersistentVolumes
		{
			name:         "persistentvolumes",
			url:          "/api/v1/resources/storage/persistentvolumes",
			expectedKind: "PersistentVolume",
		},
		{
			name:         "persistentvolumes/{uid}",
			url:          "/api/v1/resources/storage/persistentvolumes/{uid}",
			expectedKind: "PersistentVolume",
		},

		// Storage - PersistentVolumeClaims
		{
			name:         "persistentvolumeclaims",
			url:          "/api/v1/resources/storage/persistentvolumeclaims",
			expectedKind: "PersistentVolumeClaim",
		},
		{
			name:         "persistentvolumeclaims/{uid}",
			url:          "/api/v1/resources/storage/persistentvolumeclaims/{uid}",
			expectedKind: "PersistentVolumeClaim",
		},
	}

	for _, tt := range networkTests {
		testRoutesHelper(t, tt, uidMap, r)
	}

	for _, tt := range storageTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

func TestClusterOpsRoutes(t *testing.T) {
	r, err := setup()
	require.NoError(t, err)

	defer teardown()

	uidMap := map[string]string{}

	clusterOpsTests := []TestRoute{

		// Cluster Ops - Mutating Webhooks
		{
			name:         "mutatingwebhooks",
			url:          "/api/v1/resources/cluster-ops/mutatingwebhooks",
			expectedKind: "MutatingWebhookConfiguration",
		},
		{
			name:         "mutatingwebhooks/{uid}",
			url:          "/api/v1/resources/cluster-ops/mutatingwebhooks/{uid}",
			expectedKind: "MutatingWebhookConfiguration",
		},

		// Cluster Ops - Validating Webhooks
		{
			name:         "validatingwebhooks",
			url:          "/api/v1/resources/cluster-ops/validatingwebhooks",
			expectedKind: "ValidatingWebhookConfiguration",
		},
		{
			name:         "validatingwebhooks/{uid}",
			url:          "/api/v1/resources/cluster-ops/validatingwebhooks/{uid}",
			expectedKind: "ValidatingWebhookConfiguration",
		},

		// Cluster Ops - HPAs
		{
			name:         "hpas",
			url:          "/api/v1/resources/cluster-ops/hpas",
			expectedKind: "HorizontalPodAutoscaler",
		},
		{
			name:         "hpas/{uid}",
			url:          "/api/v1/resources/cluster-ops/hpas/{uid}",
			expectedKind: "HorizontalPodAutoscaler",
		},

		// Cluster Ops - Priority Classes
		{
			name:         "priority-classes",
			url:          "/api/v1/resources/cluster-ops/priority-classes",
			expectedKind: "PriorityClass",
		},
		{
			name:         "priority-classes/{uid}",
			url:          "/api/v1/resources/cluster-ops/priority-classes/{uid}",
			expectedKind: "PriorityClass",
		},

		// Cluster Ops - Runtime Classes
		{
			name:         "runtime-classes",
			url:          "/api/v1/resources/cluster-ops/runtime-classes",
			expectedKind: "RuntimeClass",
		},
		{
			name:         "runtime-classes/{uid}",
			url:          "/api/v1/resources/cluster-ops/runtime-classes/{uid}",
			expectedKind: "RuntimeClass",
		},
	}

	for _, tt := range clusterOpsTests {
		testRoutesHelper(t, tt, uidMap, r)
	}
}

// testRoutesHelper handles logic for testing getResources and getResource routes (e.g. /pods and /pods/{uid})
func testRoutesHelper(t *testing.T, tt TestRoute, uidMap map[string]string, r *chi.Mux) {
	t.Run(tt.name, func(t *testing.T) {
		maxRetries := 3
		baseTimeout := 1 * time.Second
		backoffFactor := 2

		for attempt := 1; attempt <= maxRetries; attempt++ {
			timeout := time.Duration(attempt) * baseTimeout
			// Create a new context with a timeout (timeout increases with each attempt)
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			rr := httptest.NewRecorder()

			tt.url = strings.Replace(tt.url, "{uid}", uidMap[tt.expectedKind], 1)
			req := httptest.NewRequest("GET", tt.url, nil)

			// Start serving the request for 1 second
			go func(ctx context.Context) {
				r.ServeHTTP(rr, req)
			}(ctx)

			// wait for the context to be done
			<-ctx.Done()

			require.Equal(t, http.StatusOK, rr.Code)

			body := processResponseBody(rr.Body.String())

			// if data is not empty then perform assertions
			if len(body) > 0 {
				// If uidMap.expectedKind is empty (eg. {Pod: ""}), then we need to store a UID for this kind
				if uidMap[tt.expectedKind] == "" {
					keyIndx := 5
					var data []map[string]interface{}
					json.Unmarshal(body[keyIndx:], &data)
					require.Equal(t, tt.expectedKind, data[0]["kind"])
					uidMap[tt.expectedKind] = data[0]["metadata"].(map[string]interface{})["uid"].(string)
				} else {
					keyIndx := 0
					var data map[string]interface{}
					json.Unmarshal(body[keyIndx:], &data)
					require.Equal(t, tt.expectedKind, data["Object"].(map[string]interface{})["kind"])
				}

				return
			}

			// if body was empty retry for maxRetries then fail
			if attempt < maxRetries {
				sleepDuration := time.Duration(attempt*backoffFactor) * time.Second
				fmt.Printf("Attempt %d failed, retrying in %v...\n", attempt, sleepDuration)
				fmt.Println("Body: ", rr.Body.String())
				time.Sleep(sleepDuration) // Exponential backoff
			} else {
				// After the last retry attempt, fail the test
				require.Fail(t, "Test failed after maximum retry attempts")
			}
		}
	})
}
