// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1PersistentVolume } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PersistentVolume Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'PersistentVolumes'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['storage_class'], ['capacity'], ['claim'], ['age'], ['status']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'PersistentVolume',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'local-path-pv',
        },
        spec: {
          capacity: { storage: '10Gi' },
          claimRef: { namespace: 'loki', name: 'data-loki-backend-0' },
          storageClassName: 'local-path',
        },
        status: { phase: 'Bound' },
      },
    ] as unknown as V1PersistentVolume[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'local-path-pv',
      namespace: '',
      storage_class: 'local-path',
      capacity: '10Gi',
      claim: 'data-loki-backend-0',
      status: 'Bound',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1PersistentVolume, any>[]
  expect(store.url).toEqual(`/api/v1/resources/storage/persistentvolumes?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
