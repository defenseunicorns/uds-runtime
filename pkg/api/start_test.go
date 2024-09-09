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
)

func TestHandleReconnection(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{}, nil
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{}, nil
	}

	disconnected := make(chan error, 1)

	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was updated
	assert.NotNil(t, k8sResources.Client)
	assert.NotNil(t, k8sResources.Cache)

	close(disconnected)
}

// Test createClient returns an error
func TestHandleReconnectionCreateClientError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	createClientMock := func() (*k8s.Clients, error) {
		return nil, fmt.Errorf("failed to create client")
	}

	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{}, nil
	}

	disconnected := make(chan error, 1)
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to attempt creating the client
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was not updated since client creation failed
	assert.Nil(t, k8sResources.Client)

	close(disconnected)
}

// Test createCache returns an error
func TestHandleReconnectionCreateCacheError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100")
	defer os.Unsetenv("RETRY_INTERVAL_MS")

	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{}, nil
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
	assert.Nil(t, k8sResources.Cache)

	close(disconnected)
}
