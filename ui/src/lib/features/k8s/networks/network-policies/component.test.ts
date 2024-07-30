// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1NetworkPolicy } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('NetworkPolicyTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['ingress_ports'],
      ['ingress_block'],
      ['egress_ports'],
      ['egress_block'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'NetworkPolicy',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'allow-authservice-egress-dns-lookup-via-coredns',
        namespace: 'authservice',
      },
      spec: {
        egress: [
          {
            ports: [{ port: 53, protocol: 'UDP' }],
            to: [
              { namespaceSelector: { matchLabels: { 'kubernetes.io/metadata.name': 'kube-system' } } },
              { podSelector: { matchLabels: { 'k8s-app': 'kube-dns' } } },
            ],
          },
        ],
      },
    },
    {
      apiVersion: 'v1',
      kind: 'NetworkPolicy',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'allow-keycloak-ingress-keycloak-backchannel-access',
        namespace: 'keycloak',
      },
      spec: {
        ingress: [
          {
            from: [{ ipBlock: { cidr: '0.0.0.0/0', except: ['169.254.169.254/32'] } }],
            ports: [{ port: 8080, protocol: 'TCP' }],
          },
        ],
      },
    },
  ] as unknown as V1NetworkPolicy[]

  const expectedTables = [
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      creationTimestamp: TestCreationTimestamp,
      egress_block: ', ',
      egress_ports: 'UDP:53',
      ingress_block: '-',
      ingress_ports: '-',
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
    },
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      creationTimestamp: TestCreationTimestamp,
      egress_block: '-',
      egress_ports: '-',
      ingress_block: '0.0.0.0/0 [169.254.169.254/32]',
      ingress_ports: 'TCP:8080',
      name: mockData[1].metadata?.name,
      namespace: mockData[1].metadata?.namespace,
    },
  ]

  testK8sResourceStore(
    'networkpolicies',
    mockData,
    expectedTables,
    '/api/v1/resources/networks/networkpolicies?dense=true',
    createStore,
  )
})
