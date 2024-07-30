// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1StatefulSet } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['ready'], ['up_to_date'], ['available'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'apps/v1',
      kind: 'StatefulSet',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'hello',
        namespace: 'default',
      },
      spec: {
        replicas: 1,
        serviceName: 'hello',
      },
      status: {
        replicas: 1,
        readyReplicas: 1,
        currentReplicas: 1,
        updatedReplicas: 1,
        currentRevision: 'hello',
        updateRevision: 'hello',
        availableReplicas: 1,
      },
    },
  ] as unknown as V1StatefulSet[]

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      ready: '1 / 1',
      up_to_date: 1,
      available: 1,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'statefulsets',
    mockData,
    expectedTables,
    `/api/v1/resources/workloads/statefulsets`,
    createStore,
  )
})
