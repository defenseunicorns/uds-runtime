// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package resources

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

func TestPodMetrics(t *testing.T) {
	pm := NewPodMetrics()

	t.Run("Initial State", func(t *testing.T) {
		require.Equal(t, 0, pm.GetCount())
		cpu, mem := pm.GetUsage()
		require.Equal(t, 0.0, cpu)
		require.Equal(t, 0.0, mem)
		require.Empty(t, pm.GetHistoricalUsage())
	})

	t.Run("Add Metrics", func(t *testing.T) {
		metric1 := &unstructured.Unstructured{}
		metric1.SetNamespace("default")
		metric1.SetName("metric1")
		pm.metrics["metric1"] = metric1

		metric2 := &unstructured.Unstructured{}
		metric2.SetNamespace("kube-system")
		metric2.SetName("metric2")
		pm.metrics["metric2"] = metric2

		require.Equal(t, 2, pm.GetCount())
	})

	t.Run("GetAll Metrics", func(t *testing.T) {
		allMetrics := pm.GetAll("", "")
		require.Len(t, allMetrics, 2)

		defaultMetrics := pm.GetAll("default", "")
		require.Len(t, defaultMetrics, 1)
		require.Equal(t, "metric1", defaultMetrics[0].GetName())

		kubeSystemMetrics := pm.GetAll("kube-system", "")
		require.Len(t, kubeSystemMetrics, 1)
		require.Equal(t, "metric2", kubeSystemMetrics[0].GetName())
	})
}

// CustomFakeNodeMetricsInterface implements a fake NodeMetricsInterface
type CustomFakeNodeMetricsInterface struct {
	Err error
}

func (f *CustomFakeNodeMetricsInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.NodeMetricsList, error) {
	return nil, f.Err
}

// We don't need these methods for this test, so we'll leave them unimplemented
func (f *CustomFakeNodeMetricsInterface) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.NodeMetrics, error) {
	return nil, nil
}

func (f *CustomFakeNodeMetricsInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nil, nil
}

// CustomFakeMetricsV1beta1Client implements a fake MetricsV1beta1Interface
type CustomFakeMetricsV1beta1Client struct {
	FakeNodeMetrics *CustomFakeNodeMetricsInterface
}

func (f *CustomFakeMetricsV1beta1Client) NodeMetricses() metricsv1beta1.NodeMetricsInterface {
	return f.FakeNodeMetrics
}

// We don't need these methods for this test, so we'll leave them unimplemented
func (f *CustomFakeMetricsV1beta1Client) PodMetricses(namespace string) metricsv1beta1.PodMetricsInterface {
	return nil
}

func (f *CustomFakeMetricsV1beta1Client) RESTClient() rest.Interface {
	return nil
}

func TestCollectMetrics(t *testing.T) {

	expectedError := fmt.Errorf("custom error: unable to list node metrics")

	fakeNodeMetrics := &CustomFakeNodeMetricsInterface{Err: expectedError}
	fakeMetricsClient := &CustomFakeMetricsV1beta1Client{FakeNodeMetrics: fakeNodeMetrics}

	// Create a test Pods
	podGVK := coreV1.SchemeGroupVersion.WithKind("Pod")
	pods := &ResourceList{
		Resources:       make(map[string]*unstructured.Unstructured),
		SparseResources: make(map[string]*unstructured.Unstructured),
		Changes:         make(chan struct{}, 1),
		HasSynced:       nil,
		gvk:             podGVK,
		CRDExists:       true,
	}
	podMetrics := NewPodMetrics()

	// Create a Cache instance with the mock Pods
	cache := &Cache{
		Pods:       pods,
		PodMetrics: podMetrics,
	}

	ctx := context.TODO()

	logOutput := &logCapture{}
	log.SetOutput(logOutput)

	cache.collectMetrics(ctx, fakeMetricsClient)

	require.Equal(t, cache.PodMetrics.current.CPU, float64(-1))
	require.Equal(t, cache.PodMetrics.current.Memory, float64(-1))
	require.Equal(t, cache.PodMetrics.historical[0].CPU, float64(0))
	require.Equal(t, cache.PodMetrics.historical[0].Memory, float64(0))

	require.Contains(t, logOutput.String(), expectedError.Error())
}

type logCapture struct {
	logs []string
}

func (lc *logCapture) Write(p []byte) (n int, err error) {
	lc.logs = append(lc.logs, string(p))
	return len(p), nil
}

func (lc *logCapture) String() string {
	return strings.Join(lc.logs, "")
}
