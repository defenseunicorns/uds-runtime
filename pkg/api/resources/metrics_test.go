// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestPodMetrics(t *testing.T) {
	pm := NewPodMetrics()

	// Test GetCount
	require.Equal(t, 0, pm.GetCount())

	// Test GetUsage
	cpu, mem := pm.GetUsage()
	require.Equal(t, 0.0, cpu)
	require.Equal(t, 0.0, mem)

	// Test GetHistoricalUsage
	historical := pm.GetHistoricalUsage()
	require.Empty(t, historical)

	// Add some metrics
	metric1 := &unstructured.Unstructured{}
	metric1.SetNamespace("default")
	metric1.SetName("metric1")
	pm.metrics["metric1"] = metric1

	metric2 := &unstructured.Unstructured{}
	metric2.SetNamespace("kube-system")
	metric2.SetName("metric2")
	pm.metrics["metric2"] = metric2

	// Test GetCount after adding metrics
	require.Equal(t, 2, pm.GetCount())

	// Test GetAll without namespace filter
	allMetrics := pm.GetAll("", "")
	require.Len(t, allMetrics, 2)

	// Test GetAll with namespace filter
	defaultMetrics := pm.GetAll("default", "")
	require.Len(t, defaultMetrics, 1)
	require.Equal(t, "metric1", defaultMetrics[0].GetName())

	kubeSystemMetrics := pm.GetAll("kube-system", "")
	require.Len(t, kubeSystemMetrics, 1)
	require.Equal(t, "metric2", kubeSystemMetrics[0].GetName())
}
