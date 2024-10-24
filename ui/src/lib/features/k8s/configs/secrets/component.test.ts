// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1Secret } from '@kubernetes/client-node'
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

  const name = 'Secrets'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['type'], ['keys', 'line-clamp-5 max-w-screen-md'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        data: { '.dockerconfigjson': null },
        kind: 'Secret',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'private-registry',
          namespace: 'loki',
        },
        type: 'kubernetes.io/dockerconfigjson',
      },
    ] as unknown as V1Secret[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'private-registry',
      namespace: 'loki',
      type: 'kubernetes.io/dockerconfigjson',
      keys: '.dockerconfigjson',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1Secret, any>[]
  expect(store.url).toEqual(`/api/v1/resources/configs/secrets`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
