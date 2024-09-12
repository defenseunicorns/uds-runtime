// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { ResourceStore } from '$features/k8s/store'
import {
  type ColumnWrapper,
  type CommonRow,
  type ResourceStoreInterface,
  type ResourceWithTable,
} from '$features/k8s/types'
import type {
  ExemptionElement,
  Matcher,
  Policy,
  Exemption as Resource,
} from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import ExemptionDetails from './exemption-details/component.svelte'
import ExemptionMatcher from './exemption-matcher/component.svelte'

interface Row extends CommonRow {
  title: string
  details: {
    component: typeof ExemptionDetails
    sort: string
    props: {
      exemption: ExemptionElement
    }
  }
  matcher: {
    component: typeof ExemptionMatcher
    props: {
      matcher: Matcher
    }
  }
  policies: {
    list: Policy[]
  }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Using dense=true because this CR does not use the status field
  const url = `/api/v1/resources/configs/uds-exemptions?dense=true`

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
          matcher: {
            component: ExemptionMatcher,
            props: {
              matcher: e.matcher,
            },
          },
          policies: {
            list: e.policies.sort(),
          },
        },
      }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
