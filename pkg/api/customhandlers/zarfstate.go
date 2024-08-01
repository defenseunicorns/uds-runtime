// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package customhandlers

import (
	"strings"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CreateZarfStateHandler(cache *resources.Cache) func() []unstructured.Unstructured {
	return func() []unstructured.Unstructured {
		// get secrets from cache
		secrets := cache.Secrets.GetResources()

		// get Zarf package secrets
		var zarfPkgSecrets []unstructured.Unstructured
		for _, s := range secrets {
			if strings.HasPrefix(s.GetName(), "zarf-package-") {
				zarfPkgSecrets = append(zarfPkgSecrets, s)
			}
		}

		return zarfPkgSecrets
	}
}
