// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Deployment as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '../types'

interface Row extends CommonRow {
  ready: string
  up_to_date: number
  available: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/deployments`

  const transform = transformResource<Resource, Row>((r) => ({
    ready: `${r.status?.readyReplicas ?? 0} / ${r.status?.replicas ?? 0}`,
    up_to_date: r.status?.updatedReplicas ?? 0,
    available: r.status?.conditions?.filter((c) => c.type === 'Available').length ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
