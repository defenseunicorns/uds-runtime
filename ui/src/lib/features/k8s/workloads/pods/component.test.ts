// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi, type Mock } from 'vitest'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { V1Pod } from '@kubernetes/client-node'
import { SvelteComponent } from 'svelte'
import Component from './component.svelte'
import { createStore } from './store'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type Procedure = (...args: any[]) => any

class MockEventSource {
  constructor(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    urlAssertionMock: Mock<Procedure>,
  ) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)
  }
  close() {}
}

suite('PodTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Pods'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['controlled_by'],
      ['containers'],
      ['status'],
      ['restarts'],
      ['metrics'],
      ['node'],
      ['age'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  // TODO: look into adding test capability for store.filterCallback
  // const podMetrics = [
  //   {
  //     containers: [
  //       { name: 'speaker', usage: { cpu: '2336411n', memory: '20236Ki' } },
  //       { name: 'frr', usage: { cpu: '753805n', memory: '21640Ki' } },
  //       { name: 'frr-metrics', usage: { cpu: '9399935n', memory: '7264Ki' } },
  //       { name: 'reloader', usage: { cpu: '0', memory: '792Ki' } },
  //     ],
  //     metadata: {
  //       creationTimestamp: new Date(),
  //       name: 'metallb-speaker-6nl62',
  //       namespace: 'uds-dev-stack',
  //     },
  //   },
  // ] as unknown as PodMetric[]

  const urlAssertionMock = vi.fn().mockImplementation((url: string) => {
    console.log(url)
  })

  vi.stubGlobal(
    'EventSource',
    vi
      .fn()
      // metrics EventSource is created first in createStore()
      .mockImplementationOnce((url: string) => new MockEventSource(url, urlAssertionMock)),
  )

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockContainers = [
      {
        containerID: 'containerd://1cd2b918e208d181dd3be8a6d0b651b962b1ae24946bfc8c27476f33b9e5b80b',
        image: 'quay.io/frrouting/frr:9.0.2',
        imageID: 'quay.io/frrouting/frr@sha256:086acb1278fe86118345f456a1fbfafb80c34d03f7bca9137da0729a1aee5e9c',
        lastState: {
          terminated: {
            containerID: 'containerd://cc2e1f4154a9fcd1c28e55e36635c0255338f64bb2f1d6defb9bf12dd725170f',
            exitCode: 255,
            finishedAt: '2024-07-25T13:03:14Z',
            reason: 'Unknown',
            startedAt: '2024-07-23T14:16:27Z',
          },
        },
        name: 'frr',
        ready: true,
        restartCount: 1,
        started: true,
        state: { running: { startedAt: '2024-07-25T13:03:26Z' } },
      },
      {
        containerID: 'containerd://7b0e9ea8d2c615883f4e808d9dec53294ca2f5a2e0590a2c7586c681ddb207d9',
        image: 'quay.io/frrouting/frr:9.0.2',
        imageID: 'quay.io/frrouting/frr@sha256:086acb1278fe86118345f456a1fbfafb80c34d03f7bca9137da0729a1aee5e9c',
        name: 'cp-frr-files',
        ready: true,
        restartCount: 1,
        started: false,
        state: {
          terminated: {
            containerID: 'containerd://7b0e9ea8d2c615883f4e808d9dec53294ca2f5a2e0590a2c7586c681ddb207d9',
            exitCode: 0,
            finishedAt: '2024-07-25T13:03:21Z',
            reason: 'Completed',
            startedAt: '2024-07-25T13:03:21Z',
          },
        },
      },
    ]

    const mockData = [
      {
        apiVersion: 'v1',
        kind: 'Pod',
        metadata: {
          creationTimestamp: '2024-07-25T16:10:22Z',
          name: 'metallb-speaker-6nl62',
          namespace: 'uds-dev-stack',
          ownerReferences: [
            {
              apiVersion: 'apps/v1',
              blockOwnerDeletion: true,
              controller: true,
              kind: 'DaemonSet',
              name: 'metallb-speaker',
              uid: 'f189a116-2613-4391-9066-8d9c569107c9',
            },
          ],
        },
        status: {
          conditions: [
            { lastProbeTime: null, lastTransitionTime: '2024-07-23T14:16:27Z', status: 'True', type: 'Initialized' },
          ],
          containerStatuses: [mockContainers[0]],
          initContainerStatuses: [mockContainers[1]],
          phase: 'Running',
        },
      },
    ] as unknown as V1Pod[]

    const original: any = await importOriginal()
    return {
      ...original,
      ResourceStore: vi
        .fn()
        .mockImplementation((url, transform, ...args) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTable = {
    name: 'metallb-speaker-6nl62',
    namespace: 'uds-dev-stack',
    containers: {
      component: SvelteComponent,
      props: {
        // Combine all containers (init, regular, ephemeral)
        containers: [
          {
            containerID: 'containerd://1cd2b918e208d181dd3be8a6d0b651b962b1ae24946bfc8c27476f33b9e5b80b',
            image: 'quay.io/frrouting/frr:9.0.2',
            imageID: 'quay.io/frrouting/frr@sha256:086acb1278fe86118345f456a1fbfafb80c34d03f7bca9137da0729a1aee5e9c',
            lastState: {
              terminated: {
                containerID: 'containerd://cc2e1f4154a9fcd1c28e55e36635c0255338f64bb2f1d6defb9bf12dd725170f',
                exitCode: 255,
                finishedAt: '2024-07-25T13:03:14Z',
                reason: 'Unknown',
                startedAt: '2024-07-23T14:16:27Z',
              },
            },
            name: 'frr',
            ready: true,
            restartCount: 1,
            started: true,
            state: {
              running: {
                startedAt: '2024-07-25T13:03:26Z',
              },
            },
          },
          {
            containerID: 'containerd://7b0e9ea8d2c615883f4e808d9dec53294ca2f5a2e0590a2c7586c681ddb207d9',
            image: 'quay.io/frrouting/frr:9.0.2',
            imageID: 'quay.io/frrouting/frr@sha256:086acb1278fe86118345f456a1fbfafb80c34d03f7bca9137da0729a1aee5e9c',
            name: 'cp-frr-files',
            ready: true,
            restartCount: 1,
            started: false,
            state: {
              terminated: {
                containerID: 'containerd://7b0e9ea8d2c615883f4e808d9dec53294ca2f5a2e0590a2c7586c681ddb207d9',
                exitCode: 0,
                finishedAt: '2024-07-25T13:03:21Z',
                reason: 'Completed',
                startedAt: '2024-07-25T13:03:21Z',
              },
            },
          },
        ],
      },
      sort: 2,
    },
    metrics: {
      component: SvelteComponent,
      sort: 0,
      // metrics added by store.filterCallback (not currently called in this test)
      props: { containers: [] },
    },
    restarts: 1,
    controller: 'DaemonSet',
    status: 'Running',
    node: '',
  }

  const store = createStore()

  // Assert podmetrics EventSource was created
  expect(urlAssertionMock).toHaveBeenCalledWith(`/api/v1/resources/workloads/podmetrics`)

  const start = store.start as unknown as () => ResourceWithTable<V1Pod, any>[]
  expect(store.url).toEqual('/api/v1/resources/workloads/pods')
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTable as unknown, [
    'creationTimestamp',
    'containers.component',
    'metrics.component',
  ])
  vi.unstubAllGlobals()
})
