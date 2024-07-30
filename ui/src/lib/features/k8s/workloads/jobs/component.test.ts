// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1Job } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['completions'], ['durations'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'batch/v1',
      kind: 'Job',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      completions: '0/1',
      durations: '1 minute',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore('jobs', mockData, expectedTables, `/api/v1/resources/workloads/jobs`, createStore)
})
