// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"context"
	"fmt"
	"sync"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

const MAX_HISTORY_LENGTH = 200

type Usage struct {
	Timestamp time.Time
	CPU       float64
	Memory    float64
}

type PodMetrics struct {
	sync.RWMutex
	metrics map[string]*unstructured.Unstructured
	current struct {
		CPU    float64
		Memory float64
	}
	historical []Usage
}

func NewPodMetrics() *PodMetrics {
	return &PodMetrics{
		metrics: make(map[string]*unstructured.Unstructured),
	}
}

// GetCount returns the number of metrics in the cache
func (pm *PodMetrics) GetCount() int {
	pm.RLock()
	defer pm.RUnlock()
	return len(pm.metrics)
}

// GetUsage returns the current CPU and memory usage
func (pm *PodMetrics) GetUsage() (cpu float64, mem float64) {
	pm.RLock()
	defer pm.RUnlock()
	return pm.current.CPU, pm.current.Memory
}

// GetHistoricalUsage returns the historical usage data
func (pm *PodMetrics) GetHistoricalUsage() []Usage {
	pm.RLock()
	defer pm.RUnlock()
	return pm.historical
}

// GetAll returns all metrics in the cache with optional filtering by namespace, second argument is ignored
func (pm *PodMetrics) GetAll(namespace string, _ string) []unstructured.Unstructured {
	pm.RLock()
	defer pm.RUnlock()
	result := make([]unstructured.Unstructured, 0, len(pm.metrics))
	for _, metric := range pm.metrics {
		// Filter by namespace
		if namespace == "" || metric.GetNamespace() == namespace {
			result = append(result, *metric)
		}
	}

	return result
}

// Update updates the metrics for a pod in the cache
func (pm *PodMetrics) Update(podUID string, metrics *unstructured.Unstructured) {
	pm.Lock()
	defer pm.Unlock()
	pm.metrics[podUID] = metrics
}

// Get returns the metrics for a pod in the cache
func (pm *PodMetrics) Get(podUID string) *unstructured.Unstructured {
	pm.RLock()
	defer pm.RUnlock()
	return pm.metrics[podUID]
}

// Delete removes the metrics for a pod from the cache
func (pm *PodMetrics) Delete(podUID string) {
	pm.Lock()
	defer pm.Unlock()
	delete(pm.metrics, podUID)
}

// StartMetricsCollection starts a goroutine to collect metrics for all pods in the cache
func (c *Cache) StartMetricsCollection(ctx context.Context, metricsClient *versioned.Clientset) {
	// Collect metrics immediately
	c.collectMetrics(ctx, metricsClient)

	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.collectMetrics(ctx, metricsClient)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

// Update the CalculateUsage function
func (c *Cache) CalculateUsage(metrics *v1beta1.PodMetrics) (float64, float64) {
	var totalCPU, totalMemory float64
	for _, container := range metrics.Containers {
		totalCPU += float64(container.Usage.Cpu().MilliValue())
		totalMemory += float64(container.Usage.Memory().Value())
	}

	// CPU in millicores, memory in bytes
	return totalCPU, totalMemory
}

func (c *Cache) collectMetrics(ctx context.Context, metricsClient *versioned.Clientset) {
	// Fetch all pods
	pods := c.Pods.GetSparseResources("", "")

	var totalCPU, totalMemory float64

	// Fetch metrics for each pod
	for _, pod := range pods {
		// Only collect metrics for running pods
		phase, _, _ := unstructured.NestedString(pod.Object, "status", "phase")
		if phase != "Running" {
			continue
		}

		// Fetch metrics for the pod
		metrics, err := metricsClient.MetricsV1beta1().PodMetricses(pod.GetNamespace()).Get(ctx, pod.GetName(), metaV1.GetOptions{})
		if err != nil {
			fmt.Printf("Error fetching metrics for pod %s/%s: %v\n", pod.GetNamespace(), pod.GetName(), err)
			continue
		}

		// Calculate the total CPU and memory usage
		cpu, mem := c.CalculateUsage(metrics)
		totalCPU += cpu
		totalMemory += mem

		// Convert the metrics to unstructured
		converted, err := toUnstructured(metrics)
		if err != nil {
			fmt.Printf("Error converting metrics for pod %s/%s: %v\n", pod.GetNamespace(), pod.GetName(), err)
			continue
		}

		// Update the cache with the new metrics
		c.PodMetrics.Update(string(pod.GetUID()), converted)
	}

	// Add the metrics to the cache and historical usage
	c.PodMetrics.current.CPU = totalCPU
	c.PodMetrics.current.Memory = totalMemory
	c.PodMetrics.historical = append(c.PodMetrics.historical, Usage{
		Timestamp: time.Now(),
		CPU:       totalCPU,
		Memory:    totalMemory,
	})

	// Limit the historical usage to the maximum length
	if len(c.PodMetrics.historical) > MAX_HISTORY_LENGTH {
		c.PodMetrics.historical = c.PodMetrics.historical[len(c.PodMetrics.historical)-MAX_HISTORY_LENGTH:]
	}

	// Notify subscribers of the change
	select {
	case c.MetricsChanges <- struct{}{}:
	default:
	}
}
