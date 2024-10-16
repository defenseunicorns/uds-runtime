// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1Namespace as Resource } from '@kubernetes/client-node'
import Status from '$components/k8s/Status/component.svelte'
import {
  type ColumnWrapper,
  type CommonRow,
  type K8StatusMapping,
  type ResourceStoreInterface,
} from '$features/k8s/types'

import { ResourceStore, transformResource } from '../store'

export interface Row extends CommonRow {
  status: { component: typeof Status; props: { type: keyof K8StatusMapping; status: string } }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/namespaces`

  const transform = transformResource<Resource, Row>((r) => ({
    status: { component: Status, props: { type: 'Namespaces', status: r.status?.phase ?? '' } },
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')
  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
