// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1Endpoints } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EndpointTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['endpoints'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'Endpoints',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      endpoints: '10.42.0.52:9094, 10.42.0.52:9094, 10.42.0.52:9093',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'endpoints',
    mockData,
    expectedTables,
    `/api/v1/resources/networks/endpoints?dense=true`,
    createStore,
  )
})
