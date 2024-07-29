// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import {
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
      metadata: { name: 'test', namespace: 'default', creationTimestamp: new Date() },
      status: { readyReplicas: 1, replicas: 2, updatedReplicas: 1, conditions: [{ type: 'Available' }] },
    },
  ] as V1Deployment[]

  const expectedTable = {
    name: mockData[0].metadata!.name,
    namespace: mockData[0].metadata!.namespace,
    creationTimestamp: '',
    ready: '1 / 2',
    up_to_date: 1,
    available: 1,
    age: { text: 'less than a minute', sort: 1721923822000 },
  }

  testK8sResourceStore('deployments', mockData, expectedTable, '/api/v1/resources/workloads/deployments', createStore)
})
