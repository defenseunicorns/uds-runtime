// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import { writable } from 'svelte/store'

import type { V1Namespace } from '@kubernetes/client-node'
import StatusComponent from '$components/k8s/Status/component.svelte'
import { resourceDescriptions } from '$lib/utils/descriptions'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '../test-helper'
import type { ResourceWithTable } from '../types'
import Component from './component.svelte'
import { createStore } from './store'

vi.mock('$app/stores', () => {
  const namespaces = writable(true)

  return {
    page: writable({ data: { namespaces } }),
  }
})

vi.mock('svelte/store', () => {
  return {
    writable: vi.fn().mockImplementation(<T>(initialValue: T) => {
      return {
        subscribe: (callback: (value: T) => void) => {
          callback(initialValue)
          return () => {}
        },
        set: vi.fn(),
        update: vi.fn(),
      }
    }),
  }
})

suite('NamespaceTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Namespaces'
  const description = resourceDescriptions[name]

  // Use expect.any(Function), because createStore is wrapped
  testK8sTableWithDefaults(Component, {
    createStore: expect.any(Function),
    columns: [
      ['name', 'w-4/12'],
      ['status', 'w-4/12'],
      ['age', 'w-4/12'],
    ],
    isNamespaced: false,
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, {
    createStore: expect.any(Function),
    isNamespaced: false,
    name,
    description,
  })

  vi.mock('../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'Namespace',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          labels: {
            'app.kubernetes.io/managed-by': 'zarf',
            'istio-injection': 'enabled',
            'kubernetes.io/metadata.name': 'promtail',
          },
          name: 'promtail',
        },
        status: { phase: 'Active' },
      },
    ] as unknown as V1Namespace[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'promtail',
      status: { component: StatusComponent, props: { type: 'Namespaces', status: 'Active' } },
      namespace: '',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1Namespace, any>[]
  expect(store.url).toEqual(`/api/v1/resources/namespaces`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
