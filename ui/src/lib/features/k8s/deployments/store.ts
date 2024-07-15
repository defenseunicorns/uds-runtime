// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Deployment as Resource } from '@kubernetes/client-node'

import {
    ResourceStore,
    type ColumnWrapper,
    type CommonRow,
    type ResourceStoreInterface,
    type ResourceWithTable,
} from '../store'

interface Row extends CommonRow {
  ready: string
  up_to_date: number
  available: number
}

export type Columns = ColumnWrapper<Row>

/**
 * Create a new DeploymentStore for streaming deployment resources
 *
 * @returns A new DeploymentStore instance
 */
export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/deployments`

  const transform = (resources: Resource[]) =>
    resources.map<ResourceWithTable<Resource, Row>>((r) => ({
      resource: r,
      table: {
        name: r.metadata?.name ?? '',
        namespace: r.metadata?.namespace ?? '',
        ready: `${r.status?.readyReplicas ?? 0} / ${r.status?.replicas ?? 0}`,
        up_to_date: r.status?.updatedReplicas ?? 0,
        available: r.status?.conditions?.filter((c) => c.type === 'Available').length ?? 0,
        creationTimestamp: new Date(r.metadata?.creationTimestamp ?? ''),
      },
    }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
