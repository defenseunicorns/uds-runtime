// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1PriorityClass } from '@kubernetes/client-node'
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

  const name = 'PriorityClasses'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize w-4/12'],
      ['value', 'w-1/12'],
      ['global_default', 'w-2/12'],
      ['description', 'w-4/12'],
      ['age', 'w-1/12'],
    ],
    name,
    description,
    isNamespaced: false,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description, isNamespaced: false })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'scheduling.k8s.io/v1',
        kind: 'PriorityClass',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'system-cluster-critical',
        },
        description: 'testdescription',
        value: 1,
        globalDefault: true,
      },
    ] as unknown as V1PriorityClass[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'system-cluster-critical',
      namespace: '',
      value: 1,
      description: 'testdescription',
      global_default: true,
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1PriorityClass, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/priority-classes`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
