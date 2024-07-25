// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

// toUnstructured converts an object to an Unstructured
func toUnstructured(obj interface{}) (*unstructured.Unstructured, error) {
	// If it's already Unstructured, just return it
	if u, ok := obj.(*unstructured.Unstructured); ok {
		return u, nil
	}

	// If it's a DeletedFinalStateUnknown, extract the object and convert that
	if deletedObj, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		return toUnstructured(deletedObj.Obj)
	}

	// Otherwise, convert to Unstructured
	unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: unstructuredMap}, nil
}
