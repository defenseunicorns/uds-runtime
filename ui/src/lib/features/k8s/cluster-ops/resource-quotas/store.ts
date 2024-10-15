// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1ResourceQuota as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

export type Columns = ColumnWrapper<CommonRow>

export function createStore(): ResourceStoreInterface<Resource, CommonRow> {
  const url = `/api/v1/resources/cluster-ops/resource-quotas`

  const transform = transformResource<Resource, CommonRow>(() => ({}))

  const store = new ResourceStore<Resource, CommonRow>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
