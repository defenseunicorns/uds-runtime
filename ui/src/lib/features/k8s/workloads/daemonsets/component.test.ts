// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1DaemonSet } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('DaemonsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['desired'],
      ['current'],
      ['ready'],
      ['up_to_date'],
      ['available'],
      ['node_selector'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'apps/v1',
      kind: 'DaemonSet',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'ensure-machine-id',
        namespace: 'uds-dev-stack',
      },
      spec: {
        template: {
          metadata: { creationTimestamp: null, labels: { name: 'ensure-machine-id' } },
          spec: {
            nodeSelector: { 'kubernetes.io/os': 'linux' },
          },
        },
        updateStrategy: { rollingUpdate: { maxSurge: 0, maxUnavailable: 1 }, type: 'RollingUpdate' },
        selector: { matchLabels: { name: 'ensure-machine-id' } },
      },
      status: {
        currentNumberScheduled: 1,
        desiredNumberScheduled: 1,
        numberAvailable: 1,
        numberMisscheduled: 0,
        numberReady: 1,
        observedGeneration: 1,
        updatedNumberScheduled: 1,
      },
    },
  ] as unknown as V1DaemonSet[]

  const expectedTables = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      current: 1,
      desired: 1,
      node_selector: 'kubernetes.io/os: linux',
      ready: 1,
      up_to_date: 1,
      available: 1,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore(
    'daemonset',
    mockData,
    expectedTables,
    `/api/v1/resources/workloads/daemonsets?dense=true`,
    createStore,
  )
})
