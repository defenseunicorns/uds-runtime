// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1PodDisruptionBudget as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  min_available: number | string
  max_unavailable: number | string
  current_healthy: number
  desired_healthy: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/cluster-ops/poddisruptionbudgets?fields=.metadata,.status,.spec.minAvailable,.spec.maxUnavailable`

  const transform = transformResource<Resource, Row>((r) => ({
    min_available: r.spec?.minAvailable ?? 'N/A',
    max_unavailable: r.spec?.maxUnavailable ?? 'N/A',
    desired_healthy: r.status?.desiredHealthy ?? 0,
    current_healthy: r.status?.currentHealthy ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
