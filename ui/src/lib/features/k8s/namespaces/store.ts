// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Namespace as Resource } from '@kubernetes/client-node'

import { apiAuthEnabled, authenticated } from '$features/api-auth/store'
import { get } from 'svelte/store'
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
  const resourceStoreInterface: ResourceStoreInterface<Resource, Row> = {
    ...store,
    start: () => {
      return () => {}
    },
    sortByKey: store.sortByKey.bind(store),
  }
  // If api auth is enabled we don't want to start the store until we are authenticated
  if (get(apiAuthEnabled) === false || (get(apiAuthEnabled) && get(authenticated))) {
    resourceStoreInterface.start = store.start.bind(store)
  }
  return resourceStoreInterface
}
