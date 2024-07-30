// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { VirtualService } from 'uds-core-types/src/pepr/operator/crd/generated/istio/virtualservice-v1beta1'
import Component from './component.svelte'
import { createStore } from './store'

suite('VirtualServiceTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['gateways'], ['hosts'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'networking.istio.io/v1',
      kind: 'VirtualService',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: 'keycloak-admin-admin-access-with-optional-client-certificate',
      namespace: 'keycloak',
      gateways: 'istio-admin-gateway/admin-gateway',
      hosts: 'keycloak.admin.uds.dev',
      age: { sort: 1721923882000, text: '1 minute' },
    },
  ]

  testK8sResourceStore(
    'VirtualService',
    mockData,
    expectedTables,
    `/api/v1/resources/networks/virtualservices?dense=true`,
    createStore,
  )
})
