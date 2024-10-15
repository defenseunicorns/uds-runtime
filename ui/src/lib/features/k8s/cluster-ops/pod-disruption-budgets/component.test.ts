// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1PodDisruptionBudget } from '@kubernetes/client-node'
import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'

import Component from './component.svelte'
import { createStore } from './store'

suite('PriorityClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'PodDisruptionBudgets'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize w-3/12'],
      ['namespace', 'w-1/12'],
      ['min_available', 'w-2/12'],
      ['max_unavailable', 'w-2/12'],
      ['current_healthy', 'w-2/12'],
      ['desired_healthy', 'w-2/12'],
      ['age', 'w-1/12'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'policy/v1',
        kind: 'PodDisruptionBudget',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
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

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'zarf-docker-registry',
      namespace: 'zarf',
      min_available: 1,
      max_unavailable: 1,
      current_healthy: 1,
      desired_healthy: 1,
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1PodDisruptionBudget, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/poddisruptionbudgets?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
