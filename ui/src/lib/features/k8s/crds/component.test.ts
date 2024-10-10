// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import { testK8sTableWithDefaults } from '$features/k8s/test-helper'
import { resourceDescriptions } from '$lib/utils/descriptions'

import Component from './component.svelte'
import { createStore } from './store'

suite('CustomResourceDefinitions Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'CustomResourceDefinitions'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['group'], ['kind'], ['versions'], ['scope'], ['age']],
    name,
    description,
    isNamespaced: false,
  })

  // testK8sTableWithCustomColumns(Component, { createStore, name, description, isNamespaced: false })

  // vi.mock('../../store.ts', async (importOriginal) => {
  //   const mockData = [
  //     {
  //       apiVersion: 'scheduling.k8s.io/v1',
  //       kind: 'PriorityClass',
  //       metadata: {
  //         creationTimestamp: '2024-09-29T20:00:00Z',
  //         name: 'system-cluster-critical',
  //       },
  //       description: 'testdescription',
  //       value: 1,
  //       globalDefault: true,
  //     },
  //   ] as unknown as V1CustomResourceDefinition[]

  //   const original: Record<string, unknown> = await importOriginal()
  //   return {
  //     ...original,
  //     ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
  //   }
  // })

  // const expectedTables = [
  //   {
  //     name: 'system-cluster-critical',
  //     namespace: '',
  //     value: 1,
  //     description: 'testdescription',
  //     global_default: true,
  //   },
  // ]

  // const store = createStore()
  // const start = store.start as unknown as () => ResourceWithTable<V1CustomResourceDefinition, any>[]
  // expect(store.url).toEqual(`/api/v1/resources/cluster-ops/priority-classes`)
  // // ignore creationTimestamp because age is not calculated at this point and added to the table
  // expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
