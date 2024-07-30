// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1PodDisruptionBudget } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PriorityClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['min_available'],
      ['max_unavailable'],
      ['current_healthy'],
      ['desired_healthy'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'policy/v1',
      kind: 'PodDisruptionBudget',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'zarf-docker-registry',
        namespace: 'zarf',
      },
      spec: {
        minAvailable: 1,
        maxUnavailable: 1,
        selector: { matchLabels: { app: 'docker-registry', release: 'zarf-docker-registry' } },
      },
      status: {
        currentHealthy: 1,
        desiredHealthy: 1,
        disruptionsAllowed: 0,
        expectedPods: 1,
        observedGeneration: 1,
      },
    },
  ] as unknown as V1PodDisruptionBudget[]

  const expectedTable = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      min_available: 1,
      max_unavailable: 1,
      current_healthy: 1,
      desired_healthy: 1,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'poddisruptionbudgets',
    mockData,
    expectedTable,
    `/api/v1/resources/cluster-ops/poddisruptionbudgets?dense=true`,
    createStore,
  )
})
