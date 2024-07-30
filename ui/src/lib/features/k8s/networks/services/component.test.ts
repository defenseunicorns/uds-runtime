// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1Service } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('ServiceTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['type'],
      ['cluster_ip'],
      ['external_ip'],
      ['ports'],
      ['age'],
      ['status'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'Service',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'kube-prometheus-stack-kube-state-metrics',
        namespace: 'monitoring',
      },
      spec: {
        clusterIP: '10.43.187.242',
        clusterIPs: ['10.43.187.242'],
        internalTrafficPolicy: 'Cluster',
        ipFamilies: ['IPv4'],
        ipFamilyPolicy: 'SingleStack',
        ports: [{ name: 'http', port: 8080, protocol: 'TCP', targetPort: 8080 }],
        selector: {
          'app.kubernetes.io/instance': 'kube-prometheus-stack',
          'app.kubernetes.io/name': 'kube-state-metrics',
        },
        sessionAffinity: 'None',
        type: 'ClusterIP',
      },
      status: { loadBalancer: {} },
    },
    {
      apiVersion: 'v1',
      kind: 'Service',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'passthrough-ingressgateway',
        namespace: 'istio-passthrough-gateway',
      },
      spec: {
        allocateLoadBalancerNodePorts: true,
        clusterIP: '10.43.46.245',
        clusterIPs: ['10.43.46.245'],
        externalTrafficPolicy: 'Cluster',
        internalTrafficPolicy: 'Cluster',
        ipFamilies: ['IPv4'],
        ipFamilyPolicy: 'SingleStack',
        ports: [
          { name: 'status-port', nodePort: 31801, port: 15021, protocol: 'TCP', targetPort: 15021 },
          { name: 'http2', nodePort: 31907, port: 80, protocol: 'TCP', targetPort: 80 },
          { name: 'https', nodePort: 31576, port: 443, protocol: 'TCP', targetPort: 443 },
        ],
        selector: { app: 'passthrough-ingressgateway', istio: 'passthrough-ingressgateway' },
        sessionAffinity: 'None',
        type: 'LoadBalancer',
      },
      status: { loadBalancer: { ingress: [{ ip: '172.25.0.202' }] } },
    },
    {
      apiVersion: 'v1',
      kind: 'Service',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'zarf-docker-registry',
        namespace: 'zarf',
      },
      spec: {
        clusterIP: '10.43.98.34',
        clusterIPs: ['10.43.98.34'],
        externalTrafficPolicy: 'Cluster',
        internalTrafficPolicy: 'Cluster',
        ipFamilies: ['IPv4'],
        ipFamilyPolicy: 'SingleStack',
        ports: [{ name: 'http-5000', nodePort: 31999, port: 5000, protocol: 'TCP', targetPort: 5000 }],
        selector: { app: 'docker-registry', release: 'zarf-docker-registry' },
        sessionAffinity: 'None',
        type: 'NodePort',
      },
      status: { loadBalancer: {} },
    },
  ] as unknown as V1Service[]

  const expectedTables = [
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      cluster_ip: '10.43.187.242',
      external_ip: '-',
      name: mockData[0].metadata?.name,
      namespace: mockData[0].metadata?.namespace,
      ports: '8080/TCP',
      status: 'Succeeded',
      type: mockData[0].spec?.type,
    },
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      cluster_ip: '10.43.46.245',
      external_ip: '172.25.0.202',
      name: mockData[1].metadata?.name,
      namespace: mockData[1].metadata?.namespace,
      ports: '15021:31801/TCP, 80:31907/TCP, 443:31576/TCP',
      status: 'Succeeded',
      type: mockData[1].spec?.type,
    },
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      cluster_ip: '10.43.98.34',
      external_ip: '-',
      name: mockData[2].metadata?.name,
      namespace: mockData[2].metadata?.namespace,
      ports: '5000:31999/TCP',
      status: 'Succeeded',
      type: mockData[2].spec?.type,
    },
  ]

  testK8sResourceStore(
    'services',
    mockData,
    expectedTables,
    `/api/v1/resources/networks/services?dense=true`,
    createStore,
  )
})
