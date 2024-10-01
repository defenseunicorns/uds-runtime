package k8s_session

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/client"
	"k8s.io/client-go/rest"
)

type K8sSessionCTX struct {
	Client         *client.Clients
	Cache          *resources.Cache
	CurrentCtx     string
	CurrentCluster string
	Cancel         context.CancelFunc
	InCluster      bool
}

type createClient func() (*client.Clients, error)
type createCache func(ctx context.Context, client *client.Clients) (*resources.Cache, error)

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

// isRunningInCluster checks if the application is running in cluster
func IsRunningInCluster() (bool, error) {
	_, err := rest.InClusterConfig()

	if err == rest.ErrNotInCluster {
		return false, nil
	} else if err != nil {
		return true, err
	}

	return true, nil
}

// handleReconnection is a goroutine that handles reconnection to the k8s API
// passing createClient and createCache instead of calling clients.NewClient and resources.NewCache for testing purposes
func HandleReconnection(disconnected chan error, k8sResources *K8sSessionCTX, createClient createClient,
	createCache createCache) {
	for err := range disconnected {
		log.Printf("Disconnected error received: %v\n", err)
		for {
			// Cancel the previous context
			k8sResources.Cancel()
			time.Sleep(getRetryInterval())

			currentCtx, currentCluster, err := client.GetCurrentContext()
			if err != nil {
				log.Printf("Error fetching current context: %v\n", err)
				continue
			}

			// If the current context or cluster is different from the original, skip reconnection
			if currentCtx != k8sResources.CurrentCtx || currentCluster != k8sResources.CurrentCluster {
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

			k8sResources.Client = k8sClient
			k8sResources.Cache = cache
			k8sResources.Cancel = cancel
			log.Println("Successfully reconnected to k8s and recreated cache")
			break
		}
	}
}

func CreateK8sSession() (*K8sSessionCTX, error) {
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

	// K8sResources struct to hold references
	k8sResources := &K8sSessionCTX{
		Client: k8sClient,
		Cache:  cache,
		Cancel: cancel,
	}

	return k8sResources, nil
}
