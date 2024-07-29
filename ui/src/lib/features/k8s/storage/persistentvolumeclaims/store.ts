// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1PersistentVolumeClaim as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  storage_class: string
  capacity: string
  pods: string[]
  status: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/storage/persistentvolumeclaims?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    storage_class: r.spec?.storageClassName ?? '',
    capacity: r.spec?.resources?.requests?.storage ?? '',
    status: r.status?.phase ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: () => store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
