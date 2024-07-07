import ContainerStatus from '$lib/components/data/ContainerStatus.svelte'
import type { V1Pod as Resource, V1ContainerStatus } from '@kubernetes/client-node'
import {
  ResourceStore,
  type ColumnWrapper,
  type CommonRow,
  type ResourceStoreInterface,
  type ResourceWithTable,
} from './common'

interface Row extends CommonRow {
  containers: {
    component: typeof ContainerStatus
    props: {
      containers: V1ContainerStatus[]
    }
    sort: number
  }
  restarts: number
  controller: string
  node: string
  status: string
}

export type Columns = ColumnWrapper<Row>

/**
 * Create a new PodStore for streaming Pod resources
 *
 * @returns A new PodStore instance
 */
export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/pods`

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
        restarts: r.status?.containerStatuses?.reduce((acc, curr) => acc + curr.restartCount, 0) ?? 0,
        controller: r.metadata?.ownerReferences?.at(0)?.kind ?? '',
        node: r.spec?.nodeName ?? '',
        creationTimestamp: new Date(r.metadata?.creationTimestamp ?? ''),
        status: r.status?.phase ?? '',
      },
    }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
