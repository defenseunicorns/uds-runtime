// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1Secret } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['type'], ['keys', 'line-clamp-5 max-w-screen-md'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      data: { '.dockerconfigjson': null },
      kind: 'Secret',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'private-registry',
        namespace: 'loki',
      },
      type: 'kubernetes.io/dockerconfigjson',
    },
  ] as unknown as V1Secret[]

  const expectedTable = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      type: mockData[0].type,
      keys: '.dockerconfigjson',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore('Secrets', mockData, expectedTable, `/api/v1/resources/configs/secrets`, createStore)
})
