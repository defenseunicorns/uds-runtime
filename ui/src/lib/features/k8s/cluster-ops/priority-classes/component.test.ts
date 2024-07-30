// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1PriorityClass } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PriorityClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['value'], ['global_default'], ['description'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'scheduling.k8s.io/v1',
      kind: 'PriorityClass',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'system-cluster-critical',
      },
      description: 'testdescription',
      value: 1,
      globalDefault: true,
    },
  ] as unknown as V1PriorityClass[]

  const expectedTable = [
    {
      name: mockData[0].metadata!.name,
      namespace: '',
      value: mockData[0].value,
      description: mockData[0].description,
      global_default: true,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'PriorityClasses',
    mockData,
    expectedTable,
    `/api/v1/resources/cluster-ops/priority-classes`,
    createStore,
  )
})
