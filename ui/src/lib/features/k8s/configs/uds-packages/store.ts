// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'
import type { Package as Resource } from 'uds-core-types/src/pepr/operator/crd/generated/package-v1alpha1'

import EndpointLinks from './links/component.svelte'

interface Row extends CommonRow {
  monitors: string
  endpoints: { component: typeof EndpointLinks; props: { endpoints: string[] } }
  ssoClients: string
  networkPolicies: number
  status: string
  retryAttempts: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/configs/uds-packages`

  const transform = transformResource<Resource, Row>((r) => ({
    monitors: r.status?.monitors?.join(', ') ?? '',
    endpoints: { component: EndpointLinks, props: { endpoints: r.status?.endpoints || [] } },
    ssoClients: r.status?.ssoClients?.join(', ') ?? '',
    networkPolicies: r.status?.networkPolicyCount ?? 0,
    status: r.status?.phase ?? '',
    retryAttempts: r.status?.retryAttempt ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
