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

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'autoscaling/v2',
        kind: 'HorizontalPodAutoscaler',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
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

    const original: any = await importOriginal()
    return {
      ...original,
      ResourceStore: vi
        .fn()
        .mockImplementation((url, transform, ...args) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'tenant-ingressgateway',
      namespace: 'istio-tenant-gateway',
      metrics: '4% / 80%',
      min_pods: 1,
      max_pods: 5,
      replicas: 1,
      status: 'AbleToScale ScalingActive',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<V2HorizontalPodAutoscaler, any>[]
  expect(store.url).toEqual('/api/v1/resources/cluster-ops/hpas?dense=true')
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
