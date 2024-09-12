// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { ResourceStore } from '$features/k8s/store'
import {
  type ColumnWrapper,
  type CommonRow,
  type ResourceStoreInterface,
  type ResourceWithTable,
} from '$features/k8s/types'
import type { Expose, Package as Resource } from 'uds-core-types/src/pepr/operator/crd/generated/package-v1alpha1'

interface Row extends CommonRow {
  url: {
    href: string
    text: string
    sort: string
  }
  status: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/configs/uds-packages?dense=true`

  const endpointBuilder = (r: Resource, e: Expose) => {
    if (r.status?.phase === 'Pending') {
      return { href: '', sort: '', text: 'Pending' }
    }
    const url = r.status?.endpoints?.find((ep) => ep.includes(`${e.host}.`)) ?? ''
    return {
      href: `https://${url}`,
      sort: url,
      text: url,
    }
  }

  const transform = (resources: Resource[]) =>
    resources
      // Breakout the nested spec.network.expose array into individual rows
      .flatMap((r) => (r.spec?.network?.expose ?? []).map((e) => ({ ...e, resource: r })))

      // Transform the resource into a table row
      .map<ResourceWithTable<Resource, Row>>((e) => ({
        resource: e.resource,
        table: {
          name: e.resource.metadata?.name ?? '',
          namespace: e.resource.metadata?.namespace ?? '',
          creationTimestamp: new Date(e.resource.metadata?.creationTimestamp ?? ''),
          url: endpointBuilder(e.resource, e),
          status: e.resource.status?.phase ?? '',
        },
      }))

      // Remove duplicate URLs
      .filter((e, i, a) => a.findIndex((x) => x.table.url.sort === e.table.url.sort) === i)

  const store = new ResourceStore<Resource, Row>(url, transform, 'url')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
