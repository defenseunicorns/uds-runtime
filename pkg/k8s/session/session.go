package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/client"
)

type K8sSession struct {
	Clients        *client.Clients
	Cache          *resources.Cache
	Cancel         context.CancelFunc
	CurrentCtx     string
	CurrentCluster string
	Status         chan string
	InCluster      bool
	ready          bool
	createCache    createCache
	createClient   createClient
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
		Status:         make(chan string),
		ready:          true,
		createCache:    resources.NewCache,
		createClient:   client.NewClient,
	}

	return session, nil
}

// StartClusterMonitoring is a goroutine that checks the connection to the cluster
func (ks *K8sSession) StartClusterMonitoring() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Skip if not ready, prevents race conditions using new cache
		if !ks.ready {
			continue
		}
		// Perform cluster health check
		_, err := ks.Clients.Clientset.ServerVersion()
		if err != nil {
			ks.Status <- "error"
			ks.HandleReconnection()
		} else {
			ks.Status <- "success"
		}
	}
}

// HandleReconnection infinitely retries to re-create the client and cache of the formerly connected cluster
func (ks *K8sSession) HandleReconnection() {
	log.Println("Disconnected error received")

	// Set ready to false to block cluster check ticker
	ks.ready = false

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

		k8sClient, err := ks.createClient()
		if err != nil {
			log.Printf("Retrying to create k8s client: %v\n", err)
			continue
		}

		// Create a new context and cache
		ctx, cancel := context.WithCancel(context.Background())
		cache, err := ks.createCache(ctx, k8sClient)
		if err != nil {
			log.Printf("Retrying to create cache: %v\n", err)
			continue
		}

		ks.Clients = k8sClient
		ks.Cache = cache
		ks.Cancel = cancel

		ks.ready = true
		log.Println("Successfully reconnected to cluster and recreated cache")

		break
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

// ServeConnStatus returns a handler function that streams the connection status to the client
func (ks *K8sSession) ServeConnStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers to keep connection alive
		rest.WriteHeaders(w)

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// If running in cluster don't check connection
		if ks.InCluster {
			fmt.Fprint(w, "event: close\ndata: in-cluster\n\n")
			flusher.Flush()
		}

		sendStatus := func(msg string) {
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}

		// To mitigate timing between connection start and getting status updates, immediately check cluster connection
		_, err := ks.Clients.Clientset.ServerVersion()
		if err != nil {
			sendStatus("error")
		} else {
			sendStatus("success")
		}

		// Listen for updates and send them to the client
		for {
			select {
			case msg := <-ks.Status:
				sendStatus(msg)

			case <-r.Context().Done():
				// Client disconnected
				return
			}
		}
	}
}
