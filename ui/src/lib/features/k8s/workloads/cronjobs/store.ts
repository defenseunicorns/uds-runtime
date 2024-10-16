// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1CronJob as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  schedule: string
  suspend: boolean
  active: number
  last_scheduled: string | Date
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Using dense=true due to schedule & suspend being defined in spec
  const url = `/api/v1/resources/workloads/cronjobs?dense=true`

  const transform = transformResource<Resource, Row>((r) => ({
    schedule: r.spec?.schedule ?? '',
    suspend: r.spec?.suspend ?? false,
    active: r.status?.active?.length ?? 0,
    last_scheduled: r.status?.lastScheduleTime ?? '',
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
