// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type {
  ExemptionElement,
  Exemption as Resource,
} from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import { ResourceStore } from '../store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface, type ResourceWithTable } from '../types'
import ExemptionDetails from './exemption-details/component.svelte'

interface Row extends CommonRow {
  title: string
  details: {
    component: typeof ExemptionDetails
    sort: string
    props: {
      exemption: ExemptionElement
    }
  }
  policies: string
}

export type Columns = ColumnWrapper<Row>

/**
 * Create a new UDSExemptionStore for streaming exemption resources
 *
 * @returns A new UDSExemptionStore instance
 */
export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/uds/exemptions`

  const transform = (resources: Resource[]) =>
    resources
      // Breakout the nested spec.exemptions array into individual rows
      .flatMap((r) => (r.spec?.exemptions ?? []).map((e) => ({ ...e, resource: r })))
      // Transform the resource into a table row
      .map<ResourceWithTable<Resource, Row>>((e) => ({
        resource: e.resource,
        table: {
          name: e.resource.metadata?.name ?? '',
          namespace: e.resource.metadata?.namespace ?? '',
          creationTimestamp: new Date(e.resource.metadata?.creationTimestamp ?? ''),
          title: e.title ?? '',
          details: {
            component: ExemptionDetails,
            sort: e.title ?? '',
            props: {
              exemption: e,
            },
          },
          policies: e.policies.join(', '),
        },
      }))

  const store = new ResourceStore<Resource, Row>('name')

  return {
    ...store,
    start: () => store.start(url, transform),
    sortByKey: store.sortByKey.bind(store),
  }
}
