// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package customhandlers

import (
	"encoding/json"
	"strings"

	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/zarf-dev/zarf/src/pkg/message"
	zarfTypes "github.com/zarf-dev/zarf/src/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func CreateZarfStateHandler(cache *resources.Cache) func() []unstructured.Unstructured {
	return func() []unstructured.Unstructured {
		// get secrets from cache
		secrets := cache.Secrets.GetResources()

		// get Zarf package secrets
		var zarfPkgSecrets []v1.Secret
		for _, s := range secrets {
			if strings.HasPrefix(s.GetName(), "zarf-package-") {
				var secret v1.Secret
				err := runtime.DefaultUnstructuredConverter.FromUnstructured(s.Object, &secret)
				if err != nil {
					// handle error
					continue
				}
				zarfPkgSecrets = append(zarfPkgSecrets, secret)
			}
		}

		// get Zarf state data from package secret and track deployed component statuses
		var result []unstructured.Unstructured
		for _, secret := range zarfPkgSecrets {
			componentStatuses := make(map[string]string)
			for _, val := range secret.Data {
				var deployedPkg zarfTypes.DeployedPackage
				err := json.Unmarshal(val, &deployedPkg)
				if err != nil {
					message.Warnf("failed to unmarshal secret data for %s: %v", secret.Name, err)
					continue
				}
				deployedComponents := deployedPkg.DeployedComponents

				// get status of each component and determine if all components are deployed
				status := "Deploying"
				if len(deployedComponents) > 0 {
					status = "Succeeded"
					for _, comp := range deployedComponents {
						componentStatuses[comp.Name] = string(comp.Status)
						switch comp.Status {
						case zarfTypes.ComponentStatusDeploying:
							status = "Deploying"
						case zarfTypes.ComponentStatusFailed:
							status = "Failed"
						case zarfTypes.ComponentStatusRemoving:
							status = "Removing"
						}
						if status != "Succeeded" {
							break
						}
					}
				}
				// send back unstructured data with package status
				unstructuredSecret := &unstructured.Unstructured{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"name":              deployedPkg.Name,
							"namespace":         secret.Namespace,
							"creationTimestamp": secret.CreationTimestamp,
						},
						"components": componentStatuses,
						"status":     status,
					},
				}
				result = append(result, *unstructuredSecret)
			}
		}
		return result
	}
}
