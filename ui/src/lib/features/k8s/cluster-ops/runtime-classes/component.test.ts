// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// eslint-disable @typescript-eslint/no-explicit-any
import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1RuntimeClass } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('RuntimeClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'RuntimeClasses'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['handler'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'node.k8s.io/v1',
        kind: 'RuntimeClass',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
          name: 'slight',
        },
        handler: 'testHandler',
      },
    ] as unknown as V1RuntimeClass[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'slight',
      namespace: '',
      handler: 'testHandler',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1RuntimeClass, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/runtime-classes`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
