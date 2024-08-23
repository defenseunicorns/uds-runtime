// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/cache"
)

func TestToUnstructured(t *testing.T) {
	// Test case: Object is already Unstructured
	t.Run("Already Unstructured", func(t *testing.T) {
		// Create a mock UDS Package
		mockUDSPackage := &unstructured.Unstructured{}
		mockUDSPackageName := "test-uds-package"
		mockUDSPackage.SetName(mockUDSPackageName)
		mockUDSPackage.SetUID("123e4567-e89b-12d3-a456-426614174UD5P4CK4G3")

		result, err := toUnstructured(mockUDSPackage)
		require.NoError(t, err)
		require.Equal(t, mockUDSPackage, result)
	})

	// Test case: Object is DeletedFinalStateUnknown
	t.Run("DeletedFinalStateUnknown", func(t *testing.T) {
		// Create a mock UDS Package
		mockUDSPackage := &unstructured.Unstructured{}
		mockUDSPackageName := "test-uds-package"
		mockUDSPackage.SetName(mockUDSPackageName)
		mockUDSPackage.SetUID("123e4567-e89b-12d3-a456-426614174UD5P4CK4G3")

		deletedObj := cache.DeletedFinalStateUnknown{Obj: mockUDSPackage}
		result, err := toUnstructured(deletedObj)
		require.NoError(t, err)
		require.Equal(t, mockUDSPackage, result)
	})

	// Test case: Object is a regular struct
	t.Run("Regular Struct", func(t *testing.T) {
		// Create a mock Pod
		mockPod := &corev1.Pod{}
		mockPodName := "test-pod"
		mockPod.SetName(mockPodName)
		mockPod.SetUID("123e4567-e89b-12d3-a456-426614174P0D")

		result, err := toUnstructured(mockPod)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, mockPodName, result.Object["metadata"].(map[string]interface{})["name"])
	})

	// Test case: Conversion error
	t.Run("Conversion Error", func(t *testing.T) {
		// Create an object that cannot be converted
		obj := make(chan int)
		result, err := toUnstructured(obj)
		require.Error(t, err)
		require.Nil(t, result)
	})
}
