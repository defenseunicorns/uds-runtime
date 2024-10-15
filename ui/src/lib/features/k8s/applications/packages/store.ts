// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { V1Secret as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import { type ColumnWrapper, type CommonRow, type ResourceStoreInterface } from '$features/k8s/types'
import { addToast } from '$features/toast'

import type { DeployedPackage } from './types'

export interface Row extends CommonRow {
  arch: string
  flavor: string
  version: string
  description: string
  components: {
    list: string[]
  }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  // Load secrets from the zarf namespace that start with zarf-package-, use dense mode
  const url = `/api/v1/resources/configs/secrets?dense=true&namespace=zarf&name=zarf-package-`

  const transform = transformResource<Resource, Row>((r) => {
    try {
      // Base64 decode the data
      const { data, name, deployedComponents } = JSON.parse(atob(r.data?.data ?? '')) as DeployedPackage

      return {
        name,
        arch: data.build?.architecture ?? '',
        flavor: data.build?.flavor ?? '',
        version: data.metadata?.version ?? '',
        description: data.metadata?.description ?? '',
        components: {
          list: deployedComponents.map((c) => c.name).sort(),
        },
      }
    } catch (e) {
      addToast({
        timeoutSecs: 5,
        message: `Failed to decode package data: ${e.message}`,
        type: 'error',
      })

      return {}
    }
  })

  const store = new ResourceStore<Resource, Row>(url, transform, 'name', true)

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
