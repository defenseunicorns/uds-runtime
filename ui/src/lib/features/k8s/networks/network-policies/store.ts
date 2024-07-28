// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1NetworkPolicy as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  ingress_ports: string
  ingress_block: string
  egress_ports: string
  egress_block: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Using dense=true due to use of .spec
  const url = `/api/v1/resources/networks/networkpolicies?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    ingress_ports:
      r.spec?.ingress?.flatMap((i) => i.ports?.map((p) => `${p.protocol}:${p.port}`) ?? []).join(', ') ?? '-',
    ingress_block:
      r.spec?.ingress
        ?.map((i) =>
          i.from
            ?.map((f) => {
              const cidr = f.ipBlock?.cidr
              const excepts = f.ipBlock?.except?.map((e) => `[${e}]`).join(', ')
              return excepts ? `${cidr} ${excepts}` : cidr
            })
            .join(', '),
        )
        .join(', ') ?? '-',
    egress_ports:
      r.spec?.egress?.flatMap((e) => e.ports?.map((p) => `${p.protocol}:${p.port}`) ?? []).join(', ') ?? '-',
    egress_block:
      r.spec?.egress
        ?.map((e) =>
          e.to
            ?.map((t) => {
              const cidr = t.ipBlock?.cidr
              const excepts = t.ipBlock?.except?.map((e) => `[${e}]`).join(', ')
              return excepts ? `${cidr} ${excepts}` : cidr
            })
            .join(', '),
        )
        .join(', ') ?? '-',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
