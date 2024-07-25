// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Endpoints as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  endpoints: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/networks/endpoints`

  const transform = transformResource<Resource, Row>((r) => ({
    endpoints: r.subsets?.map((e) => e).join(', ') ?? '',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}

/**
 * Success state of a Service depends on the type of service
 * https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
 * ClusterIP:     ClusterIP is defined
 * NodePort:      ClusterIP is defined
 * LoadBalancer:  ClusterIP is defined __and__ external endpoints exist
 * ExternalName:  true
 */
function isInSuccessState(resource: Resource): boolean {
  const resourceType = resource.spec?.type
  switch (resourceType) {
    case 'ExternalName':
      return true
    case 'LoadBalancer':
      if (resource.status?.loadBalancer?.ingress?.length === 0) {
        return false
      }
      break
    case 'ClusterIP':
    case 'NodePort':
    default:
      break
  }
  return resource.spec?.clusterIPs?.length ? resource.spec?.clusterIPs?.length > 0 : false
}
