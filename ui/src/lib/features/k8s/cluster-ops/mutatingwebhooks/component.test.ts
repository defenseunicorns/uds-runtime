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
import type { V1MutatingWebhookConfiguration } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'MutatingWebhookConfigurations'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['webhooks'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'admissionregistration.k8s.io/v1',
        kind: 'MutatingWebhookConfiguration',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
          name: 'istio-sidecar-injector',
        },
        webhooks: [
          {
            name: 'rev.namespace.sidecar-injector.istio.io',
          },
          {
            name: 'rev.object.sidecar-injector.istio.io',
          },
          {
            name: 'namespace.sidecar-injector.istio.io',
          },
          {
            name: 'object.sidecar-injector.istio.io',
          },
        ],
      },
    ] as unknown as V1MutatingWebhookConfiguration[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'istio-sidecar-injector',
      namespace: '',
      webhooks:
        'namespace.sidecar-injector.istio.io, object.sidecar-injector.istio.io, rev.namespace.sidecar-injector.istio.io, rev.object.sidecar-injector.istio.io',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1MutatingWebhookConfiguration, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/mutatingwebhooks?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
