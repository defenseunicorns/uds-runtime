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

// Mock function for createClient
func createMockClient() (*k8s.Clients, error) {
	return &k8s.Clients{}, nil
}

// Mock function for createCache
func createMockCache(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
	return &resources.Cache{}, nil
}

func TestHandleReconnection(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "1") // 1ms retry interval
	defer os.Unsetenv("RETRY_INTERVAL_MS")
	// Create a K8sResources struct with mock values
	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	// Mock the createClient function to return a new client or an error
	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{}, nil
	}

	// Mock the createCache function to return a new cache or an error
	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{}, nil
	}

	// Create the disconnected channel
	disconnected := make(chan error, 1)

	// Send a disconnection error to the channel to trigger reconnection logic
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(100 * time.Millisecond)

	// Verify that the K8sResources struct was updated
	assert.NotNil(t, k8sResources.Client)
	assert.NotNil(t, k8sResources.Cache)

	// Close the disconnected channel
	close(disconnected)
}

// Test for when createClient returns an error
func TestHandleReconnectionCreateClientError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100") // 1ms retry interval
	defer os.Unsetenv("RETRY_INTERVAL_MS")
	// Create a K8sResources struct with mock values
	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	// Mock the createClient function to return an error
	createClientMock := func() (*k8s.Clients, error) {
		return nil, fmt.Errorf("failed to create client")
	}

	// Mock the createCache function to return a new cache
	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return &resources.Cache{}, nil
	}

	// Create the disconnected channel
	disconnected := make(chan error, 1)

	// Send a disconnection error to the channel
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to attempt creating the client
	time.Sleep(100 * time.Millisecond)

	// Verify that the K8sResources struct was not updated since client creation failed
	assert.Nil(t, k8sResources.Client)

	// Close the disconnected channel
	close(disconnected)
}

// Test for when createCache returns an error
func TestHandleReconnectionCreateCacheError(t *testing.T) {
	os.Setenv("RETRY_INTERVAL_MS", "100") // 1ms retry interval
	defer os.Unsetenv("RETRY_INTERVAL_MS")
	// Create a K8sResources struct with mock values
	k8sResources := &K8sResources{
		Client: &k8s.Clients{},
		Cache:  &resources.Cache{},
		Cancel: func() {},
	}

	// Mock the createClient function to return a new client
	createClientMock := func() (*k8s.Clients, error) {
		return &k8s.Clients{}, nil
	}

	// Mock the createCache function to return an error
	createCacheMock := func(ctx context.Context, client *k8s.Clients) (*resources.Cache, error) {
		return nil, fmt.Errorf("failed to create cache")
	}

	// Create the disconnected channel
	disconnected := make(chan error, 1)

	// Send a disconnection error to the channel
	disconnected <- fmt.Errorf("simulated disconnection")

	// Run the handleReconnection function in a goroutine
	go handleReconnection(disconnected, k8sResources, createClientMock, createCacheMock)

	// Wait for the reconnection logic to complete
	time.Sleep(100 * time.Millisecond)

	// Verify that the K8sResources cache was not updated since cache creation failed
	assert.Nil(t, k8sResources.Cache)

	// Close the disconnected channel
	close(disconnected)
}
