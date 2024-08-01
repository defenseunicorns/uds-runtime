// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"testing"

	"github.com/defenseunicorns/uds-runtime/pkg/test"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestGetResource(t *testing.T) {
	resourceList := setupResourceList()

	// Test for existing resource
	resource, found := resourceList.GetResource("1")
	require.True(t, found)
	require.Equal(t, "mock-pod-1", resource.Object["metadata"].(map[string]interface{})["name"])

	// Test for non-existing resource
	_, found = resourceList.GetResource("non-existent")
	require.False(t, found)
}

func TestGetResources(t *testing.T) {
	resourceList := setupResourceList()

	// Test GetResources
	resources := resourceList.GetResources()
	require.Len(t, resources, 2)

	// Extract resource names
	resourceNames := make([]string, len(resources))
	for i, resource := range resources {
		resourceNames[i] = resource.Object["metadata"].(map[string]interface{})["name"].(string)
	}

	// Test check resource names
	require.Contains(t, resourceNames, "mock-pod-1")
	require.Contains(t, resourceNames, "mock-pod-2")
}

func setupResourceList() *ResourceList {
	resourceList := &ResourceList{
		Resources: make(map[string]*unstructured.Unstructured),
	}

	resourceList.Resources["1"] = test.CreateMockPod("mock-pod-1", "uds-dev-stack", "1")
	resourceList.Resources["2"] = test.CreateMockPod("mock-pod-2", "uds-dev-stack", "2")

	return resourceList
}
