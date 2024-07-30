// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
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

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'LimitRange',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'cpu-resource-constraint',
        namespace: 'default',
      },
    },
  ] as unknown as V1LimitRange[]

  const expectedTables = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'limitranges',
    mockData,
    expectedTables,
    `/api/v1/resources/cluster-ops/limit-ranges`,
    createStore,
  )
})
