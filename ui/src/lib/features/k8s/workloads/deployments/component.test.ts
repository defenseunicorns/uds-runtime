// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1Deployment } from '@kubernetes/client-node'
import '@testing-library/jest-dom'
import Component from './component.svelte'
import { createStore } from './store'

suite('DeploymentTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['ready'], ['up_to_date'], ['available'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      metadata: { name: 'test', namespace: 'default', creationTimestamp: TestCreationTimestamp },
      status: { readyReplicas: 1, replicas: 2, updatedReplicas: 1, conditions: [{ type: 'Available' }] },
    },
  ] as unknown as V1Deployment[]

  const expectedTables = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      creationTimestamp: '',
      ready: '1 / 2',
      up_to_date: 1,
      available: 1,
      age: { text: '1 minute', sort: 1721923882000 },
    },
  ]

  testK8sResourceStore('deployments', mockData, expectedTables, '/api/v1/resources/workloads/deployments', createStore)
})
