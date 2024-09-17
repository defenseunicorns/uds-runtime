// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import type { V1Deployment } from '@kubernetes/client-node'
import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'

import '@testing-library/jest-dom'

import { resourceDescriptions } from '$lib/utils/descriptions'

import Component from './component.svelte'
import { createStore } from './store'

suite('DeploymentTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Deployments'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['ready'], ['up_to_date'], ['available'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        metadata: { name: 'test', namespace: 'default', creationTimestamp: '2024-09-29T20:00:00Z' },
        status: {
          availableReplicas: 2,
          readyReplicas: 1,
          replicas: 2,
          updatedReplicas: 1,
          conditions: [{ type: 'Available' }],
        },
      },
    ] as unknown as V1Deployment[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'test',
      namespace: 'default',
      ready: '1 / 2',
      up_to_date: 1,
      available: 2,
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1Deployment, any>[]
  expect(store.url).toEqual('/api/v1/resources/workloads/deployments')
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
