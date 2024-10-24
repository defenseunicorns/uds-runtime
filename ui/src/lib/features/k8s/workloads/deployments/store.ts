// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1Deployment as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

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
    available: r.status?.availableReplicas ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
