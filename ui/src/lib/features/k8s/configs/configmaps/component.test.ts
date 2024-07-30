// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { V1ConfigMap } from '@kubernetes/client-node'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['keys', 'line-clamp-3 max-w-screen-md'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      data: {
        'add-policy': null,
        'add-svcacct': null,
        'add-user': null,
        'custom-command': null,
        initialize: null,
        'policy_0.json': null,
      },
      kind: 'ConfigMap',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
        name: 'minio',
        namespace: 'uds-dev-stack',
      },
    },
  ] as unknown as V1ConfigMap[]

  const expectedTable = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata!.namespace,
      keys: 'add-policy, add-svcacct, add-user, custom-command, initialize, policy_0.json',
      age: {
        sort: 1721923882000,
        text: '1 minute',
      },
    },
  ]

  testK8sResourceStore('ConfigMaps', mockData, expectedTable, `/api/v1/resources/configs/configmaps`, createStore)
})
