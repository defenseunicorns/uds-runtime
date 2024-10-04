package session

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/client"
)

type K8sSession struct {
	Clients        *client.Clients
	Cache          *resources.Cache
	CurrentCtx     string
	CurrentCluster string
	Cancel         context.CancelFunc
	InCluster      bool
	Status         chan string
	disconnected   chan error
	Ready          bool         // Readiness flag
	ReadyMutex     sync.RWMutex // Mutex to guard access to the readiness flag
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
		disconnected:   make(chan error),
		Ready:          true,
	}

	return session, nil
}

// HandleReconnection is a goroutine that handles reconnection to the k8s API
// passing createClient and createCache instead of calling clients.NewClient and resources.NewCache for testing purposes
func (ks *K8sSession) HandleReconnection(createClient createClient,
	createCache createCache) {
	for err := range ks.disconnected {
		log.Printf("Disconnected error received: %v\n", err)

		// Set readiness to false since we are about to reconnect
		ks.ReadyMutex.Lock()
		ks.Ready = false
		ks.ReadyMutex.Unlock()

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

			// Set readiness to true
			ks.ReadyMutex.Lock()
			ks.Ready = true
			ks.ReadyMutex.Unlock()

			log.Println("Successfully reconnected to k8s and recreated cache")

			// Purge the disconnected channel in case more errors came in while reconnecting
			for len(ks.disconnected) > 0 {
				<-ks.disconnected
			}

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

// StartClusterMonitoring is a goroutine that checks the connection to the cluster
func (ks *K8sSession) StartClusterMonitoring() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	recovering := false

	for range ticker.C {
		// Perform cluster health check
		_, err := ks.Clients.Clientset.ServerVersion()
		if err != nil {
			ks.Status <- "error"
			ks.disconnected <- err
			recovering = true
		} else if recovering {
			// if errors are resolved, send a reconnected message
			ks.Status <- "reconnected"
			// reset recovering flag
			recovering = false
		} else {
			ks.Status <- "success"
		}
	}
}

// SSE Handler for /health
func ServeConnStatus(k8sSession *K8sSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers to keep connection alive
		rest.WriteHeaders(w)

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// If running in cluster don't check for version and send error or reconnected events
		if k8sSession.InCluster {
			response := map[string]string{
				"success": "in-cluster",
			}
			data, err := json.Marshal(response)
			if err != nil {
				http.Error(w, fmt.Sprintf("data: Error: %v\n\n", err), http.StatusInternalServerError)
				return
			}
			// Write the data to the response
			fmt.Fprintf(w, "data: %s\n\n", data)

			flusher.Flush()

			return
		}

		sendData := func(msg string) {
			data, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, fmt.Sprintf("data: Error: %v\n\n", err), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		}

		// to mitigate timing between connection start and getting status updates
		// immediately check cluster connection and return error if not connected
		_, err := k8sSession.Clients.Clientset.ServerVersion()
		if err != nil {
			sendData("error")
		}

		// Listen for updates and send them to the client
		for {
			select {
			case msg := <-k8sSession.Status:
				sendData(msg)

			case <-r.Context().Done():
				// Client disconnected
				return
			}
		}
	}
}
