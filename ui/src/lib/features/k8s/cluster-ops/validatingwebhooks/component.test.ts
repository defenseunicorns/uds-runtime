// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import type { V1ValidatingWebhookConfiguration } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['webhooks'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'admissionregistration.k8s.io/v1',
        kind: 'ValidatingWebhookConfiguration',
        metadata: {
          creationTimestamp: '2024-07-31T19:23:27Z',
          name: 'metallb-webhook-configuration',
        },
        webhooks: [
          {
            name: 'bgppeervalidationwebhook.metallb.io',
          },
          {
            name: 'ipaddresspoolvalidationwebhook.metallb.io',
          },
          {
            name: 'bgpadvertisementvalidationwebhook.metallb.io',
          },
          {
            name: 'communityvalidationwebhook.metallb.io',
          },
          {
            name: 'bfdprofilevalidationwebhook.metallb.io',
          },
          {
            name: 'l2advertisementvalidationwebhook.metallb.io',
          },
        ],
      },
    ] as unknown as V1ValidatingWebhookConfiguration[]

    const original: any = await importOriginal()
    return {
      ...original,
      ResourceStore: vi
        .fn()
        .mockImplementation((url, transform, ...args) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'metallb-webhook-configuration',
      namespace: '',
      webhooks:
        'bfdprofilevalidationwebhook.metallb.io, bgpadvertisementvalidationwebhook.metallb.io, bgppeervalidationwebhook.metallb.io, communityvalidationwebhook.metallb.io, ipaddresspoolvalidationwebhook.metallb.io, l2advertisementvalidationwebhook.metallb.io',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1ValidatingWebhookConfiguration, any>[]
  expect(store.url).toEqual(`/api/v1/resources/cluster-ops/validatingwebhooks?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
