// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1CronJob } from '@kubernetes/client-node'
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

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'CronJobs'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['schedule'], ['suspend'], ['last_scheduled'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'batch/v1',
        kind: 'CronJob',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
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

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'hello',
      namespace: 'default',
      active: 0,
      schedule: '*/1 * * * *',
      suspend: false,
      last_scheduled: '2024-03-25T14:24:42Z',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1CronJob, any>[]
  expect(store.url).toEqual(`/api/v1/resources/workloads/cronjobs?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
