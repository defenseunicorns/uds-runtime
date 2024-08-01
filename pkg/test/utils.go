// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package test

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CreateMockPod(name, namespace, uid string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Pod",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
				"uid":       uid,
			},
		},
	}
}
