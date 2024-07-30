// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1PersistentVolume } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PersistentVolume Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['storage_class'], ['capacity'], ['claim'], ['age'], ['status']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'PersistentVolume',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: '',
      storage_class: 'local-path',
      capacity: '10Gi',
      claim: 'data-loki-backend-0',
      status: 'Bound',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'PersistentVolume',
    mockData,
    expectedTables,
    `/api/v1/resources/storage/persistentvolumes?dense=true`,
    createStore,
  )
})
