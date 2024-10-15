// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import type { CoreV1Event } from '@kubernetes/client-node'
import { resourceDescriptions } from '$lib/utils/descriptions'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '../test-helper'
import type { ResourceWithTable } from '../types'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Events'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['namespace', 'w-2/12'],
      ['age', 'w-1/12'],
      ['type', 'w-2/12'],
      ['reason', 'w-3/12'],
      ['object_kind', 'w-2/12'],
      ['object_name', 'w-3/12'],
      ['count', 'w-1/12'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../store.ts', async (importOriginal) => {
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
        },
        reason: 'Pulling',
        reportingComponent: 'kubelet',
        reportingInstance: 'k3d-uds-server-0',
        source: { component: 'kubelet', host: 'k3d-uds-server-0' },
        type: 'Normal',
      },
    ] as unknown as CoreV1Event[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'pepr-uds-core-watcher-8495d97876-xvbml.17e6d9becb8b1d47',
      namespace: 'pepr-system',
      object_kind: 'Pod',
      object_name: 'pepr-uds-core-watcher-8495d97876-xvbml',
      reason: 'Pulling',
      type: { component: {}, props: { type: 'Logs', status: 'Normal' } },
      count: 1,
      message: 'Pulling image "127.0.0.1:31999/defenseunicorns/pepr/controller:v0.32.7-zarf-804409620"',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<CoreV1Event, any>[]
  expect(store.url).toEqual(`/api/v1/resources/events?dense=true`)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp', 'type.component'])
})
