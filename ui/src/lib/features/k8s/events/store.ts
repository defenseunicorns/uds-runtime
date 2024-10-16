// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { CoreV1Event as Resource } from '@kubernetes/client-node'
import Status from '$components/k8s/Status/component.svelte'
import { ResourceStore, transformResource } from '$features/k8s/store'
import {
  type ColumnWrapper,
  type CommonRow,
  type K8StatusMapping,
  type ResourceStoreInterface,
} from '$features/k8s/types'

export interface Row extends CommonRow {
  count: number
  message: string
  object_kind: string
  object_name: string
  reason: string
  type: { component: typeof Status; props: { type: keyof K8StatusMapping; status: string } }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Using dense=true because most of the fields are stripped out in the default spareResource stream
  const url = `/api/v1/resources/events?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    count: r.count ?? 0,
    message: r.message ?? '',
    object_kind: r.involvedObject?.kind ?? '',
    object_name: r.involvedObject?.name ?? '',
    reason: r.reason ?? '',
    type: { component: Status, props: { type: 'Logs', status: r.type ?? '' } },
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
