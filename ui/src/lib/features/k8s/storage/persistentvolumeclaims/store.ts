// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { writable } from 'svelte/store'

import type { V1PersistentVolumeClaim as Resource, V1Pod } from '@kubernetes/client-node'
import Status from '$components/k8s/Status/component.svelte'
import { ResourceStore, transformResource } from '$features/k8s/store'
import {
  type ColumnWrapper,
  type CommonRow,
  type K8StatusMapping,
  type ResourceStoreInterface,
} from '$features/k8s/types'

interface Row extends CommonRow {
  storage_class: string
  capacity: string
  pods: string[]
  status: { component: typeof Status; props: { type: keyof K8StatusMapping; status: string } }
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/storage/persistentvolumeclaims?dense=true`

  // correlate pods with pvcs
  const pods = new Map<string, string[]>() // map of pvc name -> pod names
  const podStore = writable<number>()
  const jsonPathFields = 'metadata.name,spec.volumes,status.phase'
  const podEventsPath = `/api/v1/resources/workloads/pods?fields=${jsonPathFields}`
  const podEvents = new EventSource(podEventsPath)

  podEvents.onmessage = (event) => {
    const data = JSON.parse(event.data) as V1Pod[]
    data.forEach((p) => {
      // find the pvcs for each pod
      p.spec?.volumes?.forEach((v) => {
        const claimName = `${v.persistentVolumeClaim?.claimName}` || ''
        let podNames = pods.get(claimName) ?? []
        if (claimName && p.status?.phase === 'Running') {
          // add pod to state
          podNames.push(p.metadata?.name ?? '')
          podNames = Array.from(new Set(podNames)) // de-dup
          pods.set(claimName, podNames)
        } else if (claimName && p.status?.phase !== 'Running') {
          // remove terminated pods from  state
          podNames = podNames.filter((n) => n !== p.metadata?.name)
          pods.set(claimName, podNames)
        }
      })
    })

    // trigger an update
    podStore.set(event.timeStamp)
  }

  const transform = transformResource<Resource, Row>((r) => ({
    storage_class: r.spec?.storageClassName ?? '',
    capacity: r.spec?.resources?.requests?.storage ?? '',
    status: { component: Status, props: { type: 'PersistentVolumeClaims', status: r.status?.phase ?? '' } },
  }))

  const store = new ResourceStore<Resource, Row>(url, transform, 'name', true, [podStore])
  store.stopCallback = podEvents.close.bind(podEvents)
  store.filterCallback = (data) => {
    return data.map((d) => {
      const pvcName = d.table.name
      if (pods.has(pvcName)) {
        d.table.pods = pods.get(pvcName) ?? []
      }
      return d
    })
  }
  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
