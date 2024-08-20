// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { CoreV1Event as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

export interface Row extends CommonRow {
  count: number
  message: string
  object_kind: string
  object_name: string
  reason: string
  type: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Using dense=true because most of the fields are stripped out in the default spareResource stream
  let url = `/api/v1/resources/events?dense=true`

  // Check if API AUTH is enabled
  const apiAuthSet: boolean = import.meta.env.VITE_API_AUTH
    ? import.meta.env.VITE_API_AUTH.toLowerCase() === 'true'
    : false

  if (apiAuthSet) {
    // need to handle the multiple url search params
    url = url + `&`
  }

  const transform = transformResource<Resource, Row>((r) => ({
    count: r.count ?? 0,
    message: r.message ?? '',
    object_kind: r.involvedObject?.kind ?? '',
    object_name: r.involvedObject?.name ?? '',
    reason: r.reason ?? '',
    type: r.type ?? '',
    // A bit of a hack, but use the last seen timestamp to track age
    creationTimestamp: new Date(r.metadata.creationTimestamp ?? ''),
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'age', false)

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
