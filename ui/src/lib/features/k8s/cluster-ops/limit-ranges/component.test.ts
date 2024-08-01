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
import type { V1LimitRange } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('LimitRangesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'LimitRange',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
          name: 'cpu-resource-constraint',
          namespace: 'default',
        },
      },
    ] as unknown as V1LimitRange[]

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
      name: 'cpu-resource-constraint',
      namespace: 'default',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1LimitRange, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/limit-ranges`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
