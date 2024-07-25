// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// resources/metrics.go

package resources

import (
	"context"
	"fmt"
	"sync"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type PodMetrics struct {
	sync.RWMutex
	metrics map[string]*unstructured.Unstructured
}

func NewPodMetrics() *PodMetrics {
	return &PodMetrics{
		metrics: make(map[string]*unstructured.Unstructured),
	}
}

func (pm *PodMetrics) GetAll() []unstructured.Unstructured {
	pm.RLock()
	defer pm.RUnlock()
	result := make([]unstructured.Unstructured, 0, len(pm.metrics))
	for _, metric := range pm.metrics {
		result = append(result, *metric)
	}
	return result
}

func (pm *PodMetrics) Update(podUID string, metrics *unstructured.Unstructured) {
	pm.Lock()
	defer pm.Unlock()
	pm.metrics[podUID] = metrics
}

func (pm *PodMetrics) Get(podUID string) *unstructured.Unstructured {
	pm.RLock()
	defer pm.RUnlock()
	return pm.metrics[podUID]
}

func (pm *PodMetrics) Delete(podUID string) {
	pm.Lock()
	defer pm.Unlock()
	delete(pm.metrics, podUID)
}

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

func (c *Cache) collectMetrics(ctx context.Context, metricsClient *versioned.Clientset) {
	// Fetch all pods
	pods := c.Pods.GetSparseResources()

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

		// Convert the metrics to unstructured
		converted, err := toUnstructured(metrics)
		if err != nil {
			fmt.Printf("Error converting metrics for pod %s/%s: %v\n", pod.GetNamespace(), pod.GetName(), err)
			continue
		}

		// Update the cache with the new metrics
		c.PodMetrics.Update(string(pod.GetUID()), converted)
	}

	// Notify subscribers of the change
	select {
	case c.MetricsChanges <- struct{}{}:
	default:
	}
}
