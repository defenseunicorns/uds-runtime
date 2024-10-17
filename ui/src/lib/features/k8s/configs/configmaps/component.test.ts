// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1ConfigMap } from '@kubernetes/client-node'
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

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'ConfigMaps'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'w-2/12'],
      ['namespace', 'w-2/12'],
      ['keys', 'w-7/12 max-w-screen-md truncate'],
      ['age', 'w-1/12'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        data: {
          'add-policy': null,
          'add-svcacct': null,
          'add-user': null,
          'custom-command': null,
          initialize: null,
          'policy_0.json': null,
        },
        kind: 'ConfigMap',
        metadata: {
          creationTimestamp: '2024-05-24T14:51:22Z',
          name: 'minio',
          namespace: 'uds-dev-stack',
        },
      },
    ] as unknown as V1ConfigMap[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'minio',
      namespace: 'uds-dev-stack',
      keys: 'add-policy, add-svcacct, add-user, custom-command, initialize, policy_0.json',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1ConfigMap, any>[]
  expect(store.url).toEqual(`/api/v1/resources/configs/configmaps`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
