// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V2HorizontalPodAutoscaler as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  metrics?: string
  min_pods?: number
  max_pods?: number
  replicas?: number
  status: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/cluster-ops/hpas?dense=true`

  const transform = transformResource<Resource, Row>((r) => {
    const status = r.status?.conditions?.filter((c) => c.status === 'True')[0]?.type
    const HPAUtilization = `${r.status?.currentMetrics?.at(0)?.resource?.current.averageUtilization}%`

    return {
      metrics: `${HPAUtilization || 'unknown'} / ${r.spec?.metrics?.at(0)?.resource?.target?.averageUtilization}%`,
      min_pods: r.spec?.minReplicas,
      max_pods: r.spec?.maxReplicas,
      replicas: r.status?.currentReplicas,
      status,
    }
  })

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
