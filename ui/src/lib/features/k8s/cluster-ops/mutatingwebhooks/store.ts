// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1MutatingWebhookConfiguration as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

export interface Row extends CommonRow {
  webhooks: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/cluster-ops/mutatingwebhooks?fields=.metadata,.webhooks[].name`

  const transform = transformResource<Resource, Row>((r) => ({
    webhooks:
      r.webhooks
        ?.map((w) => w.name)
        .sort()
        .join(', ') ?? '',
  }))

  const store = new ResourceStore<Resource, Row>('name', true)

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
