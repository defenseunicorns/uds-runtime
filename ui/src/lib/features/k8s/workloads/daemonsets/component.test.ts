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
import type { V1DaemonSet } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('DaemonsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'DaemonSets'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['desired'],
      ['current'],
      ['ready'],
      ['up_to_date'],
      ['available'],
      ['node_selector'],
      ['age'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'apps/v1',
        kind: 'DaemonSet',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'ensure-machine-id',
          namespace: 'uds-dev-stack',
        },
        spec: {
          template: {
            metadata: { creationTimestamp: null, labels: { name: 'ensure-machine-id' } },
            spec: {
              nodeSelector: { 'kubernetes.io/os': 'linux' },
            },
          },
          updateStrategy: { rollingUpdate: { maxSurge: 0, maxUnavailable: 1 }, type: 'RollingUpdate' },
          selector: { matchLabels: { name: 'ensure-machine-id' } },
        },
        status: {
          currentNumberScheduled: 1,
          desiredNumberScheduled: 1,
          numberAvailable: 1,
          numberMisscheduled: 0,
          numberReady: 1,
          observedGeneration: 1,
          updatedNumberScheduled: 1,
        },
      },
    ] as unknown as V1DaemonSet[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'ensure-machine-id',
      namespace: 'uds-dev-stack',
      current: 1,
      desired: 1,
      node_selector: 'kubernetes.io/os: linux',
      ready: 1,
      up_to_date: 1,
      available: 1,
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1DaemonSet, any>[]
  expect(store.url).toEqual(`/api/v1/resources/workloads/daemonsets?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
