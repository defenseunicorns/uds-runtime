// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1StorageClass as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  provisioner: string
  reclaim_policy: string
  default: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/storage/storageclasses?dense=true`

  const isDefault = (r: Resource) =>
    r.metadata?.annotations && r.metadata.annotations['storageclass.kubernetes.io/is-default-class']

  const transform = transformResource<Resource, Row>((r) => ({
    provisioner: r.provisioner ?? '',
    reclaim_policy: r.reclaimPolicy ?? '',
    default: isDefault(r) ? 'Yes' : 'No',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
