// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Namespace as Resource } from '@kubernetes/client-node'

import { type CommonRow, type ResourceStoreInterface, type ResourceWithTable } from '$lib/types'
import { ResourceStore } from '../store'
import { type ColumnWrapper } from '../types'

export interface Row extends CommonRow {
  status: string
}

export type Columns = ColumnWrapper<Row>

/**
 * Create a new NamespaceStore for streaming namespaces
 *
 * @returns A new NamespaceStore instance
 */
export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/namespaces`

  const transform = (resources: Resource[]) =>
    resources.map<ResourceWithTable<Resource, Row>>((r) => ({
      resource: r,
      table: {
        name: r.metadata?.name ?? '',
        status: r.status?.phase ?? '',
        creationTimestamp: new Date(r.metadata?.creationTimestamp ?? ''),
      },
    }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
