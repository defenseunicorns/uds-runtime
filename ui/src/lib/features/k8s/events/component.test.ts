// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import type { CoreV1Event } from '@kubernetes/client-node'
import { testK8sResourceStore, testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['namespace'], ['age'], ['type'], ['reason'], ['object_kind'], ['object_name'], ['count']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'v1',
      count: 1,
      eventTime: null,
      firstTimestamp: '2024-07-30T01:35:20Z',
      involvedObject: {
        apiVersion: 'v1',
        fieldPath: 'spec.containers{watcher}',
        kind: 'Pod',
        name: 'pepr-uds-core-watcher-8495d97876-xvbml',
        namespace: 'pepr-system',
        resourceVersion: '1451',
        uid: '898ee594-8c5e-48bb-b86b-ad604dae2b86',
      },
      kind: 'Event',
      lastTimestamp: '2024-07-30T01:35:20Z',
      message: 'Pulling image "127.0.0.1:31999/defenseunicorns/pepr/controller:v0.32.7-zarf-804409620"',
      metadata: {
        creationTimestamp: '2024-07-30T01:35:20Z',
        name: 'pepr-uds-core-watcher-8495d97876-xvbml.17e6d9becb8b1d47',
        namespace: 'pepr-system',
        resourceVersion: '1499',
        uid: 'eacea403-6806-4074-afa6-f7362c9542dc',
      },
      reason: 'Pulling',
      reportingComponent: 'kubelet',
      reportingInstance: 'k3d-uds-server-0',
      source: { component: 'kubelet', host: 'k3d-uds-server-0' },
      type: 'Normal',
    },
  ] as unknown as CoreV1Event[]

  const expectedTable = {
    namespace: mockData[0].metadata?.namespace,
    age: {
      sort: 1721923822000,
      text: 'less than a minute',
    },
    type: mockData[0].type,
    reason: mockData[0].reason,
    object_kind: mockData[0].involvedObject?.kind,
    object_name: mockData[0].involvedObject?.name,
    count: mockData[0].count,
    message: 'Pulling image "127.0.0.1:31999/defenseunicorns/pepr/controller:v0.32.7-zarf-804409620"',
    name: 'pepr-uds-core-watcher-8495d97876-xvbml.17e6d9becb8b1d47',
  }

  testK8sResourceStore('events', mockData, expectedTable, `/api/v1/resources/events?dense=true`, createStore)
})
