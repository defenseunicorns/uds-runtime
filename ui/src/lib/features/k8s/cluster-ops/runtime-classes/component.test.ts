// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1RuntimeClass } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('RuntimeClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['handler'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'node.k8s.io/v1',
      kind: 'RuntimeClass',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'slight',
      },
      handler: 'testHandler',
    },
  ] as unknown as V1RuntimeClass[]

  const expectedTables = [
    {
      name: mockData[0].metadata!.name,
      namespace: '',
      handler: 'testHandler',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'RuntimeClasses',
    mockData,
    expectedTables,
    `/api/v1/resources/cluster-ops/runtime-classes`,
    createStore,
  )
})
