// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Node as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  status: string
  roles: string
  version: string
  pods: number
  taints: number
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/nodes`

  const transform = transformResource<Resource, Row>((r) => ({
    status:
      r.status?.conditions
        ?.filter((c) => c.type === 'Ready')
        .map((c) => c.status)
        .join(' / ') ?? '',
    // iterate over all labels and check if they are node-role.kubernetes.io/<value> = "true" and return <value>
    roles: Object.entries(r.metadata?.labels ?? {})
      .map(([key, value]) => {
        if (key.startsWith('node-role.kubernetes.io/') && value === 'true') {
          return key.replace('node-role.kubernetes.io/', '')
        }
      })
      .filter(Boolean)
      .join(', '),
    version: r.status?.nodeInfo?.kubeletVersion ?? '',
    pods: parseInt(r.status?.capacity?.pods ?? '0'),
    taints: r.spec?.taints?.length ?? 0,
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
