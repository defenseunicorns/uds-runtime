// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// eslint-disable @typescript-eslint/no-explicit-any
import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1StorageClass } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('StorageClass Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'StorageClasses'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['provisioner'], ['reclaim_policy'], ['default'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'storage.k8s.io/v1',
        kind: 'StorageClass',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'local-path',
        },
        provisioner: 'kubernetes.io/no-provisioner',
        reclaimPolicy: 'Delete',
        volumeBindingMode: 'Immediate',
        allowVolumeExpansion: true,
        age: '1 minute',
      },
    ] as unknown as V1StorageClass[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'local-path',
      namespace: '',
      provisioner: 'kubernetes.io/no-provisioner',
      reclaim_policy: 'Delete',
      default: 'No',
    },
  ]
  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1StorageClass, any>[]
  expect(store.url).toEqual('/api/v1/resources/storage/storageclasses')
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
