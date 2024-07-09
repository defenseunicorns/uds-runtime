import type { ContainerMetric, V1ContainerStatus } from '@kubernetes/client-node'

import ContainerStatus from '$lib/components/data/ContainerStatus.svelte'
import PodMetrics from '$lib/features/pods/PodMetrics.svelte'
import { type ColumnWrapper, type CommonRow } from '$lib/stores/resources/common'

export interface Row extends CommonRow {
  containers: {
    component: typeof ContainerStatus
    sort: number
    props: {
      containers: V1ContainerStatus[]
    }
  }
  restarts: number
  controller: string
  node: string
  status: string
  metrics: {
    component: typeof PodMetrics
    sort: number
    props: {
      containers: ContainerMetric[]
    }
  }
}

export type Columns = ColumnWrapper<Row>
