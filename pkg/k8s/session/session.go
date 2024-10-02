package session

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/client"
)

type K8sSession struct {
	Clients        *client.Clients
	Cache          *resources.Cache
	CurrentCtx     string
	CurrentCluster string
	Cancel         context.CancelFunc
	InCluster      bool
}

type createClient func() (*client.Clients, error)
type createCache func(ctx context.Context, client *client.Clients) (*resources.Cache, error)

// CreateK8sSession creates a new k8s session
func CreateK8sSession() (*K8sSession, error) {
	k8sClient, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cache, err := resources.NewCache(ctx, k8sClient)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}

	inCluster, err := client.IsRunningInCluster()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to check if running in cluster: %w", err)
	}

	var currentCtx, currentCluster string
	if !inCluster { // get the current context and cluster
		currentCtx, currentCluster, err = client.GetCurrentContext()
		if err != nil {
			cancel()
			return nil, fmt.Errorf("failed to get current context: %w", err)
		}
	}

	// K8sResources struct to hold references
	session := &K8sSession{
		Clients:        k8sClient,
		Cache:          cache,
		CurrentCtx:     currentCtx,
		CurrentCluster: currentCluster,
		Cancel:         cancel,
		InCluster:      inCluster,
	}

	return session, nil
}

// HandleReconnection is a goroutine that handles reconnection to the k8s API
// passing createClient and createCache instead of calling clients.NewClient and resources.NewCache for testing purposes
func (ks *K8sSession) HandleReconnection(disconnected chan error, createClient createClient,
	createCache createCache) {
	for err := range disconnected {
		log.Printf("Disconnected error received: %v\n", err)
		for {
			// Cancel the previous context
			ks.Cancel()
			time.Sleep(getRetryInterval())

			currentCtx, currentCluster, err := client.GetCurrentContext()
			if err != nil {
				log.Printf("Error fetching current context: %v\n", err)
				continue
			}

			// If the current context or cluster is different from the original, skip reconnection
			if currentCtx != ks.CurrentCtx || currentCluster != ks.CurrentCluster {
				log.Println("Current context has changed. Skipping reconnection.")
				continue
			}

			k8sClient, err := createClient()
			if err != nil {
				log.Printf("Retrying to create k8s client: %v\n", err)
				continue
			}

			// Create a new context and cache
			ctx, cancel := context.WithCancel(context.Background())
			cache, err := createCache(ctx, k8sClient)
			if err != nil {
				log.Printf("Retrying to create cache: %v\n", err)
				continue
			}

			ks.Clients = k8sClient
			ks.Cache = cache
			ks.Cancel = cancel
			log.Println("Successfully reconnected to k8s and recreated cache")
			break
		}
	}
}

// getRetryInterval returns the interval to wait before retrying to connect to the k8s API
func getRetryInterval() time.Duration {
	if interval, exists := os.LookupEnv("CONNECTION_RETRY_MS"); exists {
		parsed, err := strconv.Atoi(interval)
		if err == nil {
			return time.Duration(parsed) * time.Millisecond
		}
	}
	return 5 * time.Second // Default to 5 seconds if not set
}
