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
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { VirtualService } from 'uds-core-types/src/pepr/operator/crd/generated/istio/virtualservice-v1beta1'
import Component from './component.svelte'
import { createStore } from './store'

suite('VirtualServiceTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'VirtualServices'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['gateways'], ['hosts'], ['age']],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'networking.istio.io/v1',
        kind: 'VirtualService',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'keycloak-admin-admin-access-with-optional-client-certificate',
          namespace: 'keycloak',
        },
        spec: {
          gateways: ['istio-admin-gateway/admin-gateway'],
          hosts: ['keycloak.admin.uds.dev'],
          http: [
            {
              headers: {
                request: {
                  add: { 'istio-mtls-client-certificate': '%DOWNSTREAM_PEER_CERT%' },
                  remove: ['istio-mtls-client-certificate'],
                },
              },
              route: [{ destination: { host: 'keycloak-http.keycloak.svc.cluster.local', port: { number: 8080 } } }],
            },
          ],
        },
      },
    ] as unknown as VirtualService[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'keycloak-admin-admin-access-with-optional-client-certificate',
      namespace: 'keycloak',
      gateways: 'istio-admin-gateway/admin-gateway',
      hosts: 'keycloak.admin.uds.dev',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<VirtualService, any>[]
  expect(store.url).toEqual(`/api/v1/resources/networks/virtualservices?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
