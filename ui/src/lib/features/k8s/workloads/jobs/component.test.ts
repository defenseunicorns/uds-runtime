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
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1Job } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Jobs'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['completions'], ['durations'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'batch/v1',
        kind: 'Job',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'hello',
          namespace: 'default',
        },
        spec: {
          completions: 1,
          parallelism: 1,
        },
        status: {
          startTime: '2024-03-25T14:24:42Z',
          completionTime: '2024-03-25T14:25:42Z',
        },
      },
    ] as unknown as V1Job[]

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
      name: 'hello',
      namespace: 'default',
      completions: '0/1',
      durations: '1 minute',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1Job, any>[]
  expect(store.url).toEqual(`/api/v1/resources/workloads/jobs`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
