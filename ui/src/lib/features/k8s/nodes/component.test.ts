// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import type { V1Node } from '@kubernetes/client-node'
import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('NodeTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['status'], ['roles'], ['taints'], ['version'], ['age']],
    isNamespaced: false,
  })

  testK8sTableWithCustomColumns(Component, { createStore, isNamespaced: false })

  const mockData = [
    {
      apiVersion: 'v1',
      kind: 'Node',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        finalizers: ['wrangler.cattle.io/node'],
        labels: {
          'beta.kubernetes.io/arch': 'amd64',
          'beta.kubernetes.io/instance-type': 'k3s',
          'beta.kubernetes.io/os': 'linux',
          'kubernetes.io/arch': 'amd64',
          'kubernetes.io/hostname': 'k3d-uds-server-0',
          'kubernetes.io/os': 'linux',
          'node-role.kubernetes.io/control-plane': 'true',
          'node-role.kubernetes.io/master': 'true',
          'node.kubernetes.io/instance-type': 'k3s',
        },
        name: 'k3d-uds-server-0',
      },
      status: {
        capacity: {
          cpu: '16',
          'ephemeral-storage': '959786032Ki',
          'hugepages-1Gi': '0',
          'hugepages-2Mi': '0',
          memory: '65377872Ki',
          pods: '110',
        },
        conditions: [
          {
            lastHeartbeatTime: '2024-07-30T03:21:47Z',
            lastTransitionTime: '2024-07-30T01:29:40Z',
            message: 'kubelet has sufficient memory available',
            reason: 'KubeletHasSufficientMemory',
            status: 'False',
            type: 'MemoryPressure',
          },
          {
            lastHeartbeatTime: '2024-07-30T03:21:47Z',
            lastTransitionTime: '2024-07-30T01:29:40Z',
            message: 'kubelet has no disk pressure',
            reason: 'KubeletHasNoDiskPressure',
            status: 'False',
            type: 'DiskPressure',
          },
          {
            lastHeartbeatTime: '2024-07-30T03:21:47Z',
            lastTransitionTime: '2024-07-30T01:29:40Z',
            message: 'kubelet has sufficient PID available',
            reason: 'KubeletHasSufficientPID',
            status: 'False',
            type: 'PIDPressure',
          },
          {
            lastHeartbeatTime: '2024-07-30T03:21:47Z',
            lastTransitionTime: '2024-07-30T01:29:40Z',
            message: 'kubelet is posting ready status',
            reason: 'KubeletReady',
            status: 'True',
            type: 'Ready',
          },
        ],
        nodeInfo: {
          architecture: 'amd64',
          bootID: 'f29e328c-e483-422c-a196-0ec68ebafcc6',
          containerRuntimeVersion: 'containerd://1.7.17-k3s1',
          kernelVersion: '6.5.0-45-generic',
          kubeProxyVersion: 'v1.29.6+k3s2',
          kubeletVersion: 'v1.29.6+k3s2',
          machineID: '',
          operatingSystem: 'linux',
          osImage: 'K3s v1.29.6+k3s2',
          systemUUID: 'b840303a-f79b-1b48-4c8d-48210b5a656b',
        },
      },
    },
  ] as unknown as V1Node[]

  const execptedTables = [
    {
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
      name: mockData[0].metadata?.name,
      status: 'True',
      roles: 'control-plane, master',
      taints: 0,
      version: 'v1.29.6+k3s2',
      pods: 110,
      namespace: '',
    },
  ]

  testK8sResourceStore('nodes', mockData, execptedTables, `/api/v1/resources/nodes`, createStore)
})
