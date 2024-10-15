// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package session

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/client"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes"
)

func TestHandleReconnection(t *testing.T) {
	os.Setenv("CONNECTION_RETRY_MS", "100")
	defer os.Unsetenv("CONNECTION_RETRY_MS")

	// Mock GetCurrentContext to return the same context and cluster as the original
	client.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	createClientMock := func() (*client.Clients, error) {
		return &client.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *client.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	k8sSession := &K8sSession{
		Clients:        &client.Clients{},
		Cache:          &resources.Cache{},
		Cancel:         func() {},
		CurrentCtx:     "original-context",
		CurrentCluster: "original-cluster",
		createCache:    createCacheMock,
		createClient:   createClientMock,
	}

	require.Nil(t, k8sSession.Clients.Clientset)
	require.Nil(t, k8sSession.Cache.Pods)

	// Run the handleReconnection function in a goroutine
	go k8sSession.HandleReconnection()

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was updated
	require.NotNil(t, k8sSession.Clients.Clientset)
	require.NotNil(t, k8sSession.Cache.Pods)
}

// Test createClient returns an error
func TestHandleReconnectionCreateClientError(t *testing.T) {
	os.Setenv("CONNECTION_RETRY_MS", "100")
	defer os.Unsetenv("CONNECTION_RETRY_MS")

	// Mock GetCurrentContext to return the same context and cluster as the original
	client.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	createClientMock := func() (*client.Clients, error) {

		return nil, fmt.Errorf("failed to create client")
	}

	createCacheMock := func(ctx context.Context, client *client.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	k8sSession := &K8sSession{
		Clients:        &client.Clients{},
		Cache:          &resources.Cache{},
		Cancel:         func() {},
		CurrentCtx:     "original-context",
		CurrentCluster: "original-cluster",
		createCache:    createCacheMock,
		createClient:   createClientMock,
	}

	// Run the handleReconnection function in a goroutine
	go k8sSession.HandleReconnection()

	// Wait for the reconnection logic to attempt creating the client
	time.Sleep(200 * time.Millisecond)

	require.Nil(t, k8sSession.Clients.Clientset)
	require.Nil(t, k8sSession.Cache.Pods)
}

// Test createCache returns an error
func TestHandleReconnectionCreateCacheError(t *testing.T) {
	os.Setenv("CONNECTION_RETRY_MS", "100")
	defer os.Unsetenv("CONNECTION_RETRY_MS")

	// Mock GetCurrentContext to return the same context and cluster as the original
	client.GetCurrentContext = func() (string, string, error) {
		return "original-context", "original-cluster", nil
	}

	createClientMock := func() (*client.Clients, error) {
		return &client.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *client.Clients) (*resources.Cache, error) {
		return nil, fmt.Errorf("failed to create cache")
	}

	k8sSession := &K8sSession{
		Clients:        &client.Clients{},
		Cache:          &resources.Cache{},
		Cancel:         func() {},
		CurrentCtx:     "original-context",
		CurrentCluster: "original-cluster",
		createCache:    createCacheMock,
		createClient:   createClientMock,
	}

	// Run the handleReconnection function in a goroutine
	go k8sSession.HandleReconnection()

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources cache was not updated since cache creation failed
	require.Nil(t, k8sSession.Clients.Clientset)
	require.Nil(t, k8sSession.Cache.Pods)
}

func TestHandleReconnectionContextChanged(t *testing.T) {
	os.Setenv("CONNECTION_RETRY_MS", "100")
	defer os.Unsetenv("CONNECTION_RETRY_MS")

	// Mock GetCurrentContext to return a different context and cluster
	client.GetCurrentContext = func() (string, string, error) {
		return "new-context", "new-cluster", nil
	}

	createClientMock := func() (*client.Clients, error) {
		return &client.Clients{Clientset: &kubernetes.Clientset{}}, nil
	}

	createCacheMock := func(ctx context.Context, client *client.Clients) (*resources.Cache, error) {
		return &resources.Cache{Pods: &resources.ResourceList{}}, nil
	}

	k8sSession := &K8sSession{
		Clients:        &client.Clients{},
		Cache:          &resources.Cache{},
		Cancel:         func() {},
		CurrentCtx:     "original-context",
		CurrentCluster: "original-cluster",
		createCache:    createCacheMock,
		createClient:   createClientMock,
	}

	// Run the handleReconnection function in a goroutine
	go k8sSession.HandleReconnection()

	// Wait for the reconnection logic to complete
	time.Sleep(200 * time.Millisecond)

	// Verify that the K8sResources struct was not updated since the context/cluster has changed
	require.Nil(t, k8sSession.Clients.Clientset)
	require.Nil(t, k8sSession.Cache.Pods)
}
