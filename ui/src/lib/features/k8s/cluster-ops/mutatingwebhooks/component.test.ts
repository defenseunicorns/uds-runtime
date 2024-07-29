// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1MutatingWebhookConfiguration } from '@kubernetes/client-node'
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

  const mockData = [
    {
      apiVersion: 'admissionregistration.k8s.io/v1',
      kind: 'MutatingWebhookConfiguration',
      metadata: {
        creationTimestamp: '2024-07-27T02:34:57Z',
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

  const expectedTable = {
    name: mockData[0].metadata!.name,
    namespace: '',
    webhooks:
      'namespace.sidecar-injector.istio.io, object.sidecar-injector.istio.io, rev.namespace.sidecar-injector.istio.io, rev.object.sidecar-injector.istio.io',
    age: {
      sort: 1721923822000,
      text: 'less than a minute',
    },
  }

  testK8sResourceStore(
    'mutatingwebhooks',
    mockData,
    expectedTable,
    `/api/v1/resources/cluster-ops/mutatingwebhooks?dense=true`,
    createStore,
  )
})
