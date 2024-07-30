// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1StorageClass } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('StorageClass Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['provisioner'], ['reclaim_policy'], ['default'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'storage.k8s.io/v1',
      kind: 'StorageClass',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'local-path',
      },
      provisioner: 'kubernetes.io/no-provisioner',
      reclaimPolicy: 'Delete',
      volumeBindingMode: 'Immediate',
      allowVolumeExpansion: true,
      age: '1 minute',
    },
  ] as unknown as V1StorageClass[]

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: '',
      provisioner: mockData[0].provisioner,
      reclaim_policy: mockData[0].reclaimPolicy,
      default: 'No',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'storageclasses',
    mockData,
    expectedTables,
    '/api/v1/resources/storage/storageclasses',
    createStore,
  )
})
