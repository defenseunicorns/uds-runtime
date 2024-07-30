// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1ResourceQuota } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('ResourceQuotasTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'ResourceQuota',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'quota',
        namespace: 'default',
      },
      spec: {
        hard: {
          limits: {
            cpu: '1',
            memory: '1Gi',
          },
          requests: {
            cpu: '1',
            memory: '1Gi',
          },
        },
      },
    },
  ] as unknown as V1ResourceQuota[]

  const expectedTable = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'resource-quotas',
    mockData,
    expectedTable,
    `/api/v1/resources/cluster-ops/resource-quotas`,
    createStore,
  )
})
