// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

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

  const transform = (resources: Resource[]) => {
    // Additional CRD error check needed here since the exemption resouces are broken out to pull the
    // data needed for each table row vs using the generic transformResource function
    if (!Array.isArray(resources)) {
      // Check if the resources contain an error
      const containsError = Object.keys(resources)[0] === 'error'
      if (containsError) {
        return [
          {
            resource: Object.values(resources)[0] as Resource,
            table: {} as Row,
          },
        ]
      }
      return []
    }

    return (
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
    )
  }
  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
