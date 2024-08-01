// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Job as Resource } from '@kubernetes/client-node'
import { formatDistance } from 'date-fns'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  completions: string
  durations: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/jobs`

  const transform = transformResource<Resource, Row>((r) => ({
    completions: `${r.status?.succeeded ?? 0}/${r.spec?.completions ?? 0}`,
    durations: formatDistance(r.status?.startTime ?? new Date(), r.status?.completionTime ?? new Date()),
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
