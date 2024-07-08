// resources/metrics.go

package resources

import (
	"context"
	"fmt"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type PodMetrics struct {
	sync.RWMutex
	metrics map[string]*metricsv1beta1.PodMetrics
}

func NewPodMetrics() *PodMetrics {
	return &PodMetrics{
		metrics: make(map[string]*metricsv1beta1.PodMetrics),
	}
}

func (pm *PodMetrics) GetAll() []*metricsv1beta1.PodMetrics {
	pm.RLock()
	defer pm.RUnlock()
	result := make([]*metricsv1beta1.PodMetrics, 0, len(pm.metrics))
	for _, metric := range pm.metrics {
		result = append(result, metric)
	}
	return result
}

func (pm *PodMetrics) Update(podUID string, metrics *metricsv1beta1.PodMetrics) {
	pm.Lock()
	defer pm.Unlock()
	pm.metrics[podUID] = metrics
}

func (pm *PodMetrics) Get(podUID string) *metricsv1beta1.PodMetrics {
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
	pods := c.Pods.GetResources()

	// Fetch metrics for each pod
	for _, pod := range pods {
		// Only collect metrics for running pods
		if pod.Status.Phase != "Running" {
			continue
		}

		// Fetch metrics for the pod
		metrics, err := metricsClient.MetricsV1beta1().PodMetricses(pod.Namespace).Get(ctx, pod.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error fetching metrics for pod %s/%s: %v\n", pod.Namespace, pod.Name, err)
			continue
		}

		// Update the cache with the new metrics
		c.PodMetrics.Update(string(pod.UID), metrics)
	}

	// Notify subscribers of the change
	select {
	case c.MetricsChanges <- struct{}{}:
	default:
	}
}
