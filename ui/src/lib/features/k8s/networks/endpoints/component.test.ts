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
import type { V1Endpoints } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EndpointTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Endpoints'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['endpoints'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'Endpoints',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'alertmanager-operated',
          namespace: 'monitoring',
        },
        subsets: [
          {
            addresses: [
              {
                hostname: 'alertmanager-kube-prometheus-stack-alertmanager-0',
                ip: '10.42.0.52',
                nodeName: 'k3d-uds-server-0',
                targetRef: {
                  kind: 'Pod',
                  name: 'alertmanager-kube-prometheus-stack-alertmanager-0',
                  namespace: 'monitoring',
                  uid: '2d7bfb5e-2a97-4db6-b279-50168ace54a9',
                },
              },
            ],
            ports: [
              { name: 'udp-mesh', port: 9094, protocol: 'UDP' },
              { name: 'tcp-mesh', port: 9094, protocol: 'TCP' },
              { name: 'http-web', port: 9093, protocol: 'TCP' },
            ],
          },
        ],
      },
    ] as unknown as V1Endpoints[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'alertmanager-operated',
      namespace: 'monitoring',
      endpoints: '10.42.0.52:9094, 10.42.0.52:9094, 10.42.0.52:9093',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1Endpoints, any>[]
  expect(store.url).toEqual(`/api/v1/resources/networks/endpoints?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
