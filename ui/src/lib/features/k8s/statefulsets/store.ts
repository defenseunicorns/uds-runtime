// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1StatefulSet as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '../types'

interface Row extends CommonRow {
  ready: string
  service: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/statefulsets`

  const transform = transformResource<Resource, Row>((r) => ({
    ready: `${r.status?.readyReplicas ?? 0} / ${r.status?.replicas ?? 0}`,
    service: r.spec?.serviceName ?? '',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
