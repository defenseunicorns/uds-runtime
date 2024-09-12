// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { writable } from 'svelte/store'

import type { ContainerMetric, PodMetric, V1Pod as Resource, V1ContainerStatus } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

import ContainerStatus from './containers/component.svelte'
import PodMetrics from './metrics/component.svelte'
import { parseCPU } from './metrics/utils'
import Status from './status/component.svelte'

interface Row extends CommonRow {
  containers: {
    component: typeof ContainerStatus
    sort: number
    props: {
      containers: V1ContainerStatus[]
    }
  }
  restarts: number
  controlled_by: string
  node: string
  status: { component: typeof Status; props: { status: string } }
  metrics: {
    component: typeof PodMetrics
    sort: number
    props: {
      containers: ContainerMetric[]
    }
  }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/pods?fields=.metadata,.spec.nodeName,.status`

  const metrics = new Map<string, PodMetric>()
  // Store to trigger updates
  const metricsStore = writable<number>()
  const metricsEvents = new EventSource(`/api/v1/resources/workloads/podmetrics`)

  // Listen for new metrics
  metricsEvents.onmessage = (event) => {
    const data = JSON.parse(event.data) as PodMetric[]

    // Update the metrics map
    data.forEach((m) => {
      const key = `${m.metadata.namespace}/${m.metadata.name}`
      metrics.set(key, m)
    })

    // Trigger an update
    metricsStore.set(event.timeStamp)
  }

  const transform = transformResource<Resource, Row>((r) => ({
    containers: {
      component: ContainerStatus,
      props: {
        // Combine all containers
        containers: [
          r.status?.containerStatuses ?? [],
          r.status?.initContainerStatuses ?? [],
          r.status?.ephemeralContainerStatuses ?? [],
        ].flat(),
      },
      sort:
        // Sort by the total number of containers
        (r.status?.initContainerStatuses?.length ?? 0) +
        (r.status?.containerStatuses?.length ?? 0) +
        (r.status?.ephemeralContainerStatuses?.length ?? 0),
    },
    metrics: {
      component: PodMetrics,
      sort: 0,
      props: {
        containers: [],
      },
    },
    restarts: r.status?.containerStatuses?.reduce((acc, curr) => acc + curr.restartCount, 0) ?? 0,
    controlled_by: r.metadata?.ownerReferences?.at(0)?.kind ?? '',
    status: { component: Status, props: { status: r.status?.phase ?? '' } },
    node: r.spec?.nodeName ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  // Close the EventSource when the store is stopped
  store.stopCallback = metricsEvents.close.bind(metricsEvents)

  // Add the metrics data to the table
  store.filterCallback = (data) =>
    data.map((d) => {
      const key = `${d.resource.metadata?.namespace}/${d.resource.metadata?.name}`
      const metric = metrics.get(key)

      if (metric?.containers) {
        d.table.metrics.sort = metric.containers.reduce((sum, container) => sum + parseCPU(container.usage.cpu), 0)
        d.table.metrics.props.containers = metric.containers
      }

      return d
    })

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
