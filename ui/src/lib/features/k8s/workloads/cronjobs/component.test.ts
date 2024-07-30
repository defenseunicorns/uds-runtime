// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1CronJob } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['schedule'], ['suspend'], ['last_scheduled'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'batch/v1',
      kind: 'CronJob',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'hello',
        namespace: 'default',
      },
      spec: {
        schedule: '*/1 * * * *',
        suspend: false,
      },
      status: {
        lastScheduleTime: '2024-03-25T14:24:42Z',
      },
    },
  ] as unknown as V1CronJob[]

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      active: 0,
      schedule: mockData[0].spec?.schedule,
      suspend: false,
      last_scheduled: '2024-03-25T14:24:42Z',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'cronjobs',
    mockData,
    expectedTables,
    `/api/v1/resources/workloads/cronjobs?dense=true`,
    createStore,
  )
})
