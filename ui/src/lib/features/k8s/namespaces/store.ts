// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Namespace as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '../types'

export interface Row extends CommonRow {
  status: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/namespaces`

  const transform = transformResource<Resource, Row>((r) => ({
    status: r.status?.phase ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
