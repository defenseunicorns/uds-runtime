// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1Service as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  type: string
  cluster_ip: string
  external_ip: string
  ports: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/networks/services?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    type: r.spec?.type ?? '',
    cluster_ip: r.spec?.clusterIP ?? '',
    external_ip: r.status?.loadBalancer?.ingress?.map((p) => `${p.ip}`).join(', ') ?? '-',
    ports:
      r.spec?.ports
        ?.map((p) => (p.nodePort ? `${p.port}:${p.nodePort}/${p.protocol}` : `${p.port}/${p.protocol}`))
        .join(', ') ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
