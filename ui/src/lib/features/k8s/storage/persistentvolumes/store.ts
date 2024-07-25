// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1PersistentVolume as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  storage_class: string
  capacity: string
  claim: string
  status: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/storage/persistentvolumes`

  const transform = transformResource<Resource, Row>((r) => ({
    storage_class: r.spec?.storageClassName ?? '',
    capacity: r.spec?.capacity?.storage ?? '',
    claim: r.spec?.claimRef?.name ?? '',
    status: r.status?.phase ?? '',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
