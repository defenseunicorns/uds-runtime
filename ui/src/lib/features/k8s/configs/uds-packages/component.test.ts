// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { Package } from 'uds-core-types/src/pepr/operator/crd/generated/package-v1alpha1'
import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['status'],
      ['endpoints'],
      ['monitors'],
      ['ssoClients'],
      ['networkPolicies'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'uds.dev/v1alpha1',
      kind: 'Package',
      metadata: {
        annotations: {
          'meta.helm.sh/release-name': 'uds-neuvector-config',
          'meta.helm.sh/release-namespace': 'neuvector',
        },
        creationTimestamp: TestCreationTimestamp,
        generation: 1,
        labels: { 'app.kubernetes.io/managed-by': 'Helm' },
        name: 'neuvector',
        namespace: 'neuvector',
        resourceVersion: '2451',
        uid: '963b2b20-a5f4-4b70-8562-3245c696913e',
      },
      status: {
        authserviceClients: [],
        endpoints: ['neuvector.admin.uds.dev', '2.admin.uds.dev'],
        monitors: ['testMonitor', 'testMonitor2'],
        networkPolicyCount: 13,
        observedGeneration: 1,
        phase: 'Ready',
        retryAttempt: 0,
        ssoClients: ['uds-core-admin-neuvector'],
      },
    },
  ] as unknown as Package[]

  const expectedTable = [
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      creationTimestamp: '2024-07-25T16:10:22.000Z',
      endpoints: 'neuvector.admin.uds.dev, 2.admin.uds.dev',
      monitors: 'testMonitor, testMonitor2',
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata?.namespace,
      networkPolicies: mockData[0].status?.networkPolicyCount,
      retryAttempts: mockData[0].status?.retryAttempt,
      ssoClients: 'uds-core-admin-neuvector',
      status: 'Ready',
    },
  ]

  testK8sResourceStore('uds-packages', mockData, expectedTable, `/api/v1/resources/configs/uds-packages`, createStore)
})
