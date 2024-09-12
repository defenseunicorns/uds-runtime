// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { Package } from 'uds-core-types/src/pepr/operator/crd/generated/package-v1alpha1'
import { beforeEach, vi } from 'vitest'

import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Packages'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['status'],
      ['endpoints'],
      ['monitors'],
      ['ssoClients'],
      ['networkPolicies'],
      ['age'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'uds.dev/v1alpha1',
        kind: 'Package',
        metadata: {
          creationTimestamp: '2024-07-25T16:10:22Z',
          name: 'neuvector',
          namespace: 'neuvector',
        },
        status: {
          authserviceClients: [],
          endpoints: ['neuvector.admin.uds.dev', '2.admin.uds.dev'],
          monitors: ['testMonitor', 'testMonitor2'],
          networkPolicyCount: 13,
          observedGeneration: 1,
          phase: 'Ready',
          retryAttempt: 0,
          ssoClients: ['uds-core-admin-neuvector'],
        },
      },
    ] as unknown as Package[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      creationTimestamp: '2024-07-25T16:10:22.000Z',
      endpoints: { component: {}, props: { endpoints: ['neuvector.admin.uds.dev', '2.admin.uds.dev'] } },
      monitors: 'testMonitor, testMonitor2',
      name: 'neuvector',
      namespace: 'neuvector',
      networkPolicies: 13,
      retryAttempts: 0,
      ssoClients: 'uds-core-admin-neuvector',
      status: 'Ready',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<Package, any>[]
  expect(store.url).toEqual(`/api/v1/resources/configs/uds-packages`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, [
    'creationTimestamp',
    'endpoints.component',
  ])
})
