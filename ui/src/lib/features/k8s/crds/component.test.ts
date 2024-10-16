// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1CustomResourceDefinition } from '@kubernetes/client-node'
import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import { resourceDescriptions } from '$lib/utils/descriptions'

import type { ResourceWithTable } from '../types'
import Component from './component.svelte'
import { createStore } from './store'

suite('CustomResourceDefinitions Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'CustomResourceDefinitions'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['group'], ['kind'], ['versions'], ['scope'], ['age']],
    name,
    description,
    isNamespaced: false,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description, isNamespaced: false })

  vi.mock('../store.ts', async (importOriginal) => {
    const mockData = [
      {
        metadata: {
          creationTimestamp: '2024-10-08T13:55:09Z',
          name: 'helmchartconfigs.helm.cattle.io',
          uid: 'd280c873-4535-490f-b30f-670523cf3bb0',
        },
        spec: {
          group: 'helm.cattle.io',
          names: { kind: 'HelmChartConfig' },
          scope: 'Namespaced',
          versions: [{ name: 'v1' }],
        },
      },
    ] as unknown as V1CustomResourceDefinition[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      group: 'helm.cattle.io',
      kind: 'HelmChartConfig',
      name: 'helmchartconfigs.helm.cattle.io',
      scope: 'Namespaced',
      versions: 'v1',
      namespace: '',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1CustomResourceDefinition, any>[]
  expect(store.url).toEqual(
    `/api/v1/resources/custom-resource-definitions?fields=metadata,spec.group,spec.names.kind,spec.versions[].name,spec.scope`,
  )
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
