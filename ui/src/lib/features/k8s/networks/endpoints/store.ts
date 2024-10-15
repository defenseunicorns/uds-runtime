// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1Endpoints as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  endpoints: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/networks/endpoints?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    endpoints:
      r.subsets
        ?.map((subset) =>
          subset.addresses
            ?.map((address) => subset.ports?.map((port) => `${address.ip}:${port.port}`).join(', '))
            .join(', '),
        )
        .join(', ') ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),

    sortByKey: store.sortByKey.bind(store),
  }
}
