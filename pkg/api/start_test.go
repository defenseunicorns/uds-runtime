//go:build unit

package api

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes"
)

func TestHandleReconnection(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	// Mock GetCurrentContext to return the same context and cluster as the original
	k8s.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	k8sResources := &K8sResources{
		Client:          &k8s.Clients{},
		Cache:           &resources.Cache{},
		Cancel:          func() {},
		OriginalCtx:     "original-context",
		OriginalCluster: "original-cluster",
	}

	assert.Nil(t, k8sResources.Client.Clientset)
	assert.Nil(t, k8sResources.Cache.Pods)

	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	disconnected := make(chan error, 1)

	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was updated
	assert.NotNil(t, k8sResources.Client.Clientset)
	assert.NotNil(t, k8sResources.Cache.Pods)

	close(disconnected)
}

// Test createClient returns an error
func TestHandleReconnectionCreateClientError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	k8sResources := &K8sResources{
		Client:          &k8s.Clients{},
		Cache:           &resources.Cache{},
		Cancel:          func() {},
		OriginalCtx:     "original-context",
		OriginalCluster: "original-cluster",
	}

	// Mock GetCurrentContext to return the same context and cluster as the original
	k8s.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	createClientMock := func() (*k8s.Clients, error) {

		return nil, fmt.Errorf("failed to create client")
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	disconnected := make(chan error, 1)
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to attempt creating the client
	time.Sleep(200 * time.Millisecond)

	assert.Nil(t, k8sResources.Client.Clientset)
	assert.Nil(t, k8sResources.Cache.Pods)

	close(disconnected)
}

// Test createCache returns an error
func TestHandleReconnectionCreateCacheError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	k8sResources := &K8sResources{
		Client:          &k8s.Clients{},
		Cache:           &resources.Cache{},
		Cancel:          func() {},
		OriginalCtx:     "original-context",
		OriginalCluster: "original-cluster",
	}

	// Mock GetCurrentContext to return the same context and cluster as the original
	k8s.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return nil, fmt.Errorf("failed to create cache")
	}

	disconnected := make(chan error, 1)

	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources cache was not updated since cache creation failed
	assert.Nil(t, k8sResources.Client.Clientset)
	assert.Nil(t, k8sResources.Cache.Pods)

	close(disconnected)
}

func TestHandleReconnectionContextChanged(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	// Mock GetCurrentContext to return a different context and cluster
	k8s.GetCurrentContext = func() (string, string, error) {
		return "new-context", "new-cluster", nil
	}

	k8sResources := &K8sResources{
		Client:          &k8s.Clients{},
		Cache:           &resources.Cache{},
		Cancel:          func() {},
		OriginalCtx:     "original-context",
		OriginalCluster: "original-cluster",
	}

	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	disconnected := make(chan error, 1)

	// Simulate a disconnection
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was not updated since the context/cluster has changed
	assert.Nil(t, k8sResources.Client.Clientset)
	assert.Nil(t, k8sResources.Cache.Pods)

	close(disconnected)
}
