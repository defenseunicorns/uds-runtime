// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1ResourceQuota as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

export type Columns = ColumnWrapper<CommonRow>

export function createStore(): ResourceStoreInterface<Resource, CommonRow> {
  const url = `/api/v1/resources/cluster-ops/resource-quotas`

  const transform = transformResource<Resource, CommonRow>(() => ({}))

  const store = new ResourceStore<Resource, CommonRow>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
