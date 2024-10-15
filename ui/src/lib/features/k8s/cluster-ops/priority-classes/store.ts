// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1PriorityClass as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  description: string
  global_default: boolean
  value: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/cluster-ops/priority-classes`

  const transform = transformResource<Resource, Row>((r) => ({
    description: r.description ?? '',
    global_default: r.globalDefault ?? false,
    value: r.value ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
