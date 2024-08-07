package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestNodesOnce(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start serving the request for 1 second
	var testServer *httptest.Server
	go func(ctx context.Context) {
		testServer = httptest.NewServer(r)
	}(ctx)

	// wait for the context to be done
	<-ctx.Done()

	resp, err := testServer.Client().Get(testServer.URL + "/api/v1/resources/nodes?once=true")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body) // Read the response body
	require.NoError(t, err)

	fmt.Println("***********" + string(body))
}

func TestPodsSSE(t *testing.T) {
	r, err := api.Setup(nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req := httptest.NewRequest("GET", "/api/v1/resources/workloads/pods", nil)

	// Start serving the request for 1 second
	go func(ctx context.Context) {
		r.ServeHTTP(rr, req)
	}(ctx)

	// wait for the context to be done
	<-ctx.Done()

	// unstructured to pod
	var data []map[string]interface{}
	var pod v1.Pod
	jsonStr := string(rr.Body.Bytes()[5:])
	err = json.Unmarshal([]byte(jsonStr), &data)
	require.NoError(t, err)
	unstr := unstructured.Unstructured{Object: data[0]}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstr.Object, &pod)

	fmt.Println(pod)
}
