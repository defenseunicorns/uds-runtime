// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
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

    const original: any = await importOriginal()
    return {
      ...original,
      ResourceStore: vi
        .fn()
        .mockImplementation((url, transform, ...args) => new MockResourceStore(url, transform, mockData)),
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
