// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import type { V1PersistentVolumeClaim } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PersistentVolumeClaim Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['storage_class'], ['capacity'], ['age'], ['status']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

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

    const original: any = await importOriginal()
    return {
      ...original,
      ResourceStore: vi
        .fn()
        .mockImplementation((url, transform, ...args) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'data-loki-backend-0',
      namespace: 'loki',
      storage_class: 'local-path',
      capacity: '10Gi',
      status: 'Bound',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1PersistentVolumeClaim, any>[]
  expect(store.url).toEqual(`/api/v1/resources/storage/persistentvolumeclaims?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
