// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1PersistentVolumeClaim } from '@kubernetes/client-node'
import StatusComponent from '$components/k8s/Status/component.svelte'
import {
  expectEqualIgnoringFields,
  MockEventSource,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import { vi } from 'vitest'

import Component from './component.svelte'
import { createStore } from './store'

suite('PersistentVolumeClaim Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'PersistentVolumeClaims'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize w-3/12'],
      ['namespace', 'w-2/12'],
      ['storage_class', 'w-2/12'],
      ['capacity', 'w-2/12'],
      ['pods', 'w-3/12'],
      ['age', 'w-1/12'],
      ['status', 'w-1/12'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  const urlAssertionMock = vi.fn().mockImplementation((url: string) => {
    console.log(url)
  })

  vi.stubGlobal(
    'EventSource',
    vi
      .fn()
      // pods EventSource is created first in createStore()
      .mockImplementationOnce((url: string) => new MockEventSource(url, urlAssertionMock)),
  )

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'PersistentVolumeClaim',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'data-loki-backend-0',
          namespace: 'loki',
        },
        spec: {
          accessModes: ['ReadWriteOnce'],
          resources: { requests: { storage: '10Gi' } },
          storageClassName: 'local-path',
          volumeMode: 'Filesystem',
          volumeName: 'pvc-eba4c087-a7ad-43a5-a32a-fce07be0404b',
        },
        status: { accessModes: ['ReadWriteOnce'], capacity: { storage: '10Gi' }, phase: 'Bound' },
      },
    ] as unknown as V1PersistentVolumeClaim[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'data-loki-backend-0',
      namespace: 'loki',
      storage_class: 'local-path',
      capacity: '10Gi',
      status: { component: StatusComponent, props: { type: 'PersistentVolumeClaims', status: 'Bound' } },
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1PersistentVolumeClaim, any>[]
  expect(store.url).toEqual(`/api/v1/resources/storage/persistentvolumeclaims?dense=true`)

  // ignore creationTimestamp and pods because neither are calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp', 'status.component'])
  vi.unstubAllGlobals()
})
