// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1StatefulSet } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'StatefulSets'
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
        apiVersion: 'apps/v1',
        kind: 'StatefulSet',
        metadata: {
          creationTimestamp: '2024-05-24T14:51:22Z',
          name: 'hello',
          namespace: 'default',
        },
        spec: {
          replicas: 1,
          serviceName: 'hello',
        },
        status: {
          replicas: 1,
          readyReplicas: 1,
          currentReplicas: 1,
          updatedReplicas: 1,
          currentRevision: 'hello',
          updateRevision: 'hello',
          availableReplicas: 1,
        },
      },
    ] as unknown as V1StatefulSet[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'hello',
      namespace: 'default',
      ready: '1 / 1',
      up_to_date: 1,
      available: 1,
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1StatefulSet, any>[]
  expect(store.url).toEqual(`/api/v1/resources/workloads/statefulsets`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
