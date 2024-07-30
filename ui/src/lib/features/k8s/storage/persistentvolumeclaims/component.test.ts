// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
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

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'PersistentVolumeClaim',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      storage_class: 'local-path',
      capacity: '10Gi',
      status: 'Bound',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'PersistentVolumeClaim',
    mockData,
    expectedTables,
    `/api/v1/resources/storage/persistentvolumeclaims?dense=true`,
    createStore,
  )
})
