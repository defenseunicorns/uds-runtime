// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1ResourceQuota } from '@kubernetes/client-node'
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

suite('ResourceQuotasTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'ResourceQuotas'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'ResourceQuota',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'quota',
          namespace: 'default',
        },
        spec: {
          hard: {
            limits: {
              cpu: '1',
              memory: '1Gi',
            },
            requests: {
              cpu: '1',
              memory: '1Gi',
            },
          },
        },
      },
    ] as unknown as V1ResourceQuota[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'quota',
      namespace: 'default',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1ResourceQuota, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/resource-quotas`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
