// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'
import { writable } from 'svelte/store'

import type { V1Namespace } from '@kubernetes/client-node'
import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

vi.mock('$app/stores', () => {
  const namespaces = writable(true)

  return {
    page: writable({ data: { namespaces } }),
  }
})

suite('NamespaceTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  // Use expect.any(Function), because createStore is wrapped
  testK8sTableWithDefaults(Component, {
    createStore: expect.any(Function),
    columns: [['name', 'emphasize'], ['status'], ['age']],
    isNamespaced: false,
  })

  testK8sTableWithCustomColumns(Component, { createStore: expect.any(Function), isNamespaced: false })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'Namespace',
      metadata: {
        annotations: { 'uds.dev/original-istio-injection': 'non-existent', 'uds.dev/pkg-promtail': 'true' },
        creationTimestamp: TestCreationTimestamp,
        labels: {
          'app.kubernetes.io/managed-by': 'zarf',
          'istio-injection': 'enabled',
          'kubernetes.io/metadata.name': 'promtail',
        },
        name: 'promtail',
        resourceVersion: '3833',
        uid: 'ea570f9f-aa86-4793-a718-4f92686c1c08',
      },
      status: { phase: 'Active' },
    },
  ] as unknown as V1Namespace[]

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      status: mockData[0].status?.phase,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      namespace: '',
    },
  ]

  testK8sResourceStore('namespaces', mockData, expectedTables, `/api/v1/resources/namespaces`, createStore)
})
