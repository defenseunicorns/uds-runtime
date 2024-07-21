// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { Package as Resource } from 'uds-core-types/src/pepr/operator/crd/generated/package-v1alpha1'

import { ResourceStore, transformResource } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '../types'

interface Row extends CommonRow {
  monitors: string
  endpoints: string
  ssoClients: string
  networkPolicies: number
  status: string
  retryAttempts: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/config/uds-packages`

  const transform = transformResource<Resource, Row>((r) => ({
    monitors: r.status?.monitors?.join(', ') ?? '',
    endpoints: r.status?.endpoints?.join(', ') ?? '',
    ssoClients: r.status?.ssoClients?.join(', ') ?? '',
    networkPolicies: r.status?.networkPolicyCount ?? 0,
    status: r.status?.phase ?? '',
    retryAttempts: r.status?.retryAttempt ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
