// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { V1NetworkPolicy } from '@kubernetes/client-node'
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

suite('NetworkPolicyTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'NetworkPolicies'
  const description = resourceDescriptions[name]

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
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'NetworkPolicy',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
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
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'allow-keycloak-ingress-keycloak-backchannel-access',
          namespace: 'keycloak',
        },
        spec: {
          ingress: [
            {
              _from: [{ ipBlock: { cidr: '0.0.0.0/0', except: ['169.254.169.254/32'] } }],
              ports: [{ port: 8080, protocol: 'TCP' }],
            },
            {
              from: [{ ipBlock: { cidr: '0.0.0.0/0', except: ['169.254.169.254/32'] } }],
              ports: [{ port: 8080, protocol: 'TCP' }],
            },
          ],
        },
      },
    ] as unknown as V1NetworkPolicy[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      egress_block: ', ',
      egress_ports: 'UDP:53',
      ingress_block: '-',
      ingress_ports: '-',
      name: 'allow-authservice-egress-dns-lookup-via-coredns',
      namespace: 'authservice',
    },
    {
      egress_block: '-',
      egress_ports: '-',
      ingress_block: '0.0.0.0/0 [169.254.169.254/32], 0.0.0.0/0 [169.254.169.254/32]',
      ingress_ports: 'TCP:8080, TCP:8080',
      name: 'allow-keycloak-ingress-keycloak-backchannel-access',
      namespace: 'keycloak',
    },
  ]
  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V1NetworkPolicy, any>[]
  expect(store.url).toEqual('/api/v1/resources/networks/networkpolicies?dense=true')
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
  expectEqualIgnoringFields(start()[1].table, expectedTables[1] as unknown, ['creationTimestamp'])
})
