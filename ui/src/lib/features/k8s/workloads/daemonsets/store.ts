// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1DaemonSet as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  desired: number
  current: number
  ready: number
  up_to_date: number
  available: number
  node_selector: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/daemonsets?fields=.metadata,.status,.spec.template.spec.nodeSelector`

  const transform = transformResource<Resource, Row>((r) => ({
    desired: r.status?.desiredNumberScheduled ?? 0,
    current: r.status?.currentNumberScheduled ?? 0,
    ready: r.status?.numberReady ?? 0,
    up_to_date: r.status?.updatedNumberScheduled ?? 0,
    available: r.status?.numberAvailable ?? 0,
    node_selector: r.spec?.template.spec?.nodeSelector
      ? Object.entries(r.spec?.template.spec?.nodeSelector ?? {})
          .map(([key, value]) => `${key}: ${value}`)
          .join(', ')
      : '-',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
