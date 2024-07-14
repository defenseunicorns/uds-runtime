// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

import type { ContainerMetric, PodMetric, V1Pod as Resource, V1ContainerStatus } from '@kubernetes/client-node'
import { writable } from 'svelte/store'

import ContainerStatus from '$components/data/ContainerStatus.svelte'
import { parseCPU } from '$components/data/PodMetrics'
import PodMetrics from '$components/data/PodMetrics.svelte'
import {
  ResourceStore,
  type ColumnWrapper,
  type CommonRow,
  type ResourceStoreInterface,
  type ResourceWithTable,
} from '../store'

interface Row extends CommonRow {
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

/**
 * Create a new PodStore for streaming Pod resources
 *
 * @returns A new PodStore instance
 */
export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/pods`

  const metrics = new Map<string, PodMetric>()
  // Store to trigger updates
  const metricsStore = writable<number>()
  const metricsEvents = new EventSource(`/api/v1/resources/podmetrics`)

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

  const transform = (resources: Resource[]) =>
    resources.map<ResourceWithTable<Resource, Row>>((r) => ({
      resource: r,
      table: {
        name: r.metadata?.name ?? '',
        namespace: r.metadata?.namespace ?? '',
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
        controller: r.metadata?.ownerReferences?.at(0)?.kind ?? '',
        node: r.spec?.nodeName ?? '',
        creationTimestamp: new Date(r.metadata?.creationTimestamp ?? ''),
        status: r.status?.phase ?? '',
      },
    }))

  const store = new ResourceStore<Resource, Row>('name')

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
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
