// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V2HorizontalPodAutoscaler } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('PriorityClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['metrics'],
      ['min_pods'],
      ['max_pods'],
      ['replicas'],
      ['age'],
      ['status'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'autoscaling/v2',
      kind: 'HorizontalPodAutoscaler',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'tenant-ingressgateway',
        namespace: 'istio-tenant-gateway',
      },
      spec: {
        maxReplicas: 5,
        metrics: [
          { resource: { name: 'cpu', target: { averageUtilization: 80, type: 'Utilization' } }, type: 'Resource' },
        ],
        minReplicas: 1,
        scaleTargetRef: { apiVersion: 'apps/v1', kind: 'Deployment', name: 'tenant-ingressgateway' },
      },
      status: {
        conditions: [
          {
            status: 'True',
            type: 'AbleToScale',
          },
          {
            status: 'True',
            type: 'ScalingActive',
          },
          {
            status: 'False',
            type: 'ScalingLimited',
          },
        ],
        currentMetrics: [
          { resource: { current: { averageUtilization: 4, averageValue: '4m' }, name: 'cpu' }, type: 'Resource' },
        ],
        currentReplicas: 1,
        desiredReplicas: 1,
      },
    },
  ] as unknown as V2HorizontalPodAutoscaler[]

  const expectedTables = [
    {
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      metrics: '4% / 80%',
      min_pods: 1,
      max_pods: 5,
      replicas: 1,
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      status: 'AbleToScale ScalingActive',
    },
  ]

  testK8sResourceStore(
    'HorizontalPodAutoscaler',
    mockData,
    expectedTables,
    '/api/v1/resources/cluster-ops/hpas?dense=true',
    createStore,
  )
})
