// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { V1CronJob as Resource } from '@kubernetes/client-node'

import { ResourceStore, transformResource } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '../types'

interface Row extends CommonRow {
  schedule: string
  suspend: boolean
  active: number
  last_scheduled: string | Date
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/workloads/cronjobs`

  const transform = transformResource<Resource, Row>((r) => ({
    schedule: r.spec?.schedule ?? '',
    suspend: r.spec?.suspend ?? false,
    active: r.status?.active?.length ?? 0,
    last_scheduled: r.status?.lastScheduleTime ?? '',
  }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
