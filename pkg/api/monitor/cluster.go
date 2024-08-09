// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/sse"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ClusterData struct {
	TotalPods       int               `json:"totalPods"`
	TotalNodes      int               `json:"totalNodes"`
	CPUCapacity     float64           `json:"cpuCapacity"`
	MemoryCapacity  float64           `json:"memoryCapacity"`
	CurrentUsage    resources.Usage   `json:"currentUsage"`
	HistoricalUsage []resources.Usage `json:"historicalUsage"`
}

func BindClusterOverviewHandler(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	// Return a function that sends the data to the client
	return func(w http.ResponseWriter, r *http.Request) {
		sse.WriteHeaders(w)
		// Ensure the ResponseWriter supports flushing
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// Get the current data
		getUsage := func(cache *resources.Cache) {
			var clusterData ClusterData

			// Timestamp the data
			clusterData.CurrentUsage.Timestamp = time.Now()
			// Get all available pod metrics
			clusterData.TotalPods = cache.PodMetrics.GetCount()
			// Get the current usage
			clusterData.CurrentUsage.CPU, clusterData.CurrentUsage.Memory = cache.PodMetrics.GetUsage()
			// Get the historical usage
			clusterData.HistoricalUsage = cache.PodMetrics.GetHistoricalUsage()

			// Load node data
			nodes := cache.Nodes.GetSparseResources("", "")

			// Calculate the total number of nodes
			clusterData.TotalNodes = len(nodes)

			// Calculate the total capacity of the cluster
			clusterData.CPUCapacity = 0.0
			clusterData.MemoryCapacity = 0.0

			// Get the capacity of all nodes in the cluster
			for _, node := range nodes {
				// Get the CPU capacity for the node
				cpu, found, err := unstructured.NestedString(node.Object, "status", "capacity", "cpu")
				if !found || err != nil {
					continue
				}
				parsed, err := strconv.ParseFloat(cpu, 64)
				if err != nil {
					continue
				}

				// Convert from cores to milli-cores
				clusterData.CPUCapacity += (parsed * 1000)

				// Get the memory capacity for the node
				mem, found, err := unstructured.NestedString(node.Object, "status", "capacity", "memory")
				if !found || err != nil {
					continue
				}
				parsed, err = parseMemory(mem)
				if err != nil {
					continue
				}

				clusterData.MemoryCapacity += parsed
			}

			// Flush the headers at the end
			defer flusher.Flush()

			// Convert the data to JSON
			data, err := json.Marshal(clusterData)
			if err != nil {
				fmt.Fprintf(w, "data: Error: %v\n\n", err)
				return
			}

			// Write the data to the response
			fmt.Fprintf(w, "data: %s\n\n", data)
		}

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		// Send the initial data
		getUsage(cache)

		for {
			select {
			// If the context is done, return
			case <-r.Context().Done():
				return

			// If there is a pending update, send the data immediately
			case <-cache.MetricsChanges:
				getUsage(cache)

			// Respond to node changes
			case <-cache.Nodes.Changes:
				getUsage(cache)
			}
		}
	}
}

func parseMemory(memoryString string) (float64, error) {
	// Remove the 'i' suffix if present
	memoryString = strings.TrimSuffix(memoryString, "i")

	// Extract the numeric part and the unit
	var value float64
	var unit string
	var err error

	if strings.HasSuffix(memoryString, "K") {
		value, err = strconv.ParseFloat(strings.TrimSuffix(memoryString, "K"), 64)
		unit = "K"
	} else if strings.HasSuffix(memoryString, "M") {
		value, err = strconv.ParseFloat(strings.TrimSuffix(memoryString, "M"), 64)
		unit = "M"
	} else if strings.HasSuffix(memoryString, "G") {
		value, err = strconv.ParseFloat(strings.TrimSuffix(memoryString, "G"), 64)
		unit = "G"
	} else {
		value, err = strconv.ParseFloat(memoryString, 64)
		unit = ""
	}

	if err != nil {
		return 0, fmt.Errorf("failed to parse memory value: %v", err)
	}

	// Convert to bytes
	switch unit {
	case "K":
		value *= 1024
	case "M":
		value *= 1024 * 1024
	case "G":
		value *= 1024 * 1024 * 1024
	}

	return value, nil
}
