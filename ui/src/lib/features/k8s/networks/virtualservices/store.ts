// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { VirtualService as Resource } from 'uds-core-types/src/pepr/operator/crd/generated/istio/virtualservice-v1beta1'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  gateways: string
  hosts: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/networks/virtualservices?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    gateways: r.spec?.gateways?.join(', ') ?? '',
    hosts: r.spec?.hosts?.join(', ') ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: () => store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
