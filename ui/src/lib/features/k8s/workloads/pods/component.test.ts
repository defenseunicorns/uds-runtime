// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import {
  expectEqualIgnoringFields,
  MockEventSource,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { PodMetric, V1Pod } from '@kubernetes/client-node'
import { SvelteComponent } from 'svelte'
import Component from './component.svelte'
import { createStore } from './store'

suite('PodTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['controller'],
      ['containers'],
      ['status'],
      ['restarts'],
      ['metrics'],
      ['node'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  test(`createStore for pods`, async () => {
    vi.useFakeTimers()
    vi.setSystemTime(new Date())

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

    const mockPods = [
      {
        apiVersion: 'v1',
        kind: 'Pod',
        metadata: {
          creationTimestamp: new Date().toString(),
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

    const podMetrics = [
      {
        containers: [
          { name: 'speaker', usage: { cpu: '2336411n', memory: '20236Ki' } },
          { name: 'frr', usage: { cpu: '753805n', memory: '21640Ki' } },
          { name: 'frr-metrics', usage: { cpu: '9399935n', memory: '7264Ki' } },
          { name: 'reloader', usage: { cpu: '0', memory: '792Ki' } },
        ],
        metadata: {
          creationTimestamp: new Date(),
          name: 'metallb-speaker-6nl62',
          namespace: 'uds-dev-stack',
        },
      },
    ] as unknown as PodMetric[]

    const tableCols = [
      'name',
      'namespace',
      'creationTimestamp',
      'containers',
      'metrics',
      'restarts',
      'controller',
      'status',
      'node',
      'age',
    ]

    const expectedTables = {
      name: mockPods[0].metadata!.name,
      namespace: mockPods[0].metadata!.namespace,
      creationTimestamp: new Date(),
      containers: {
        component: SvelteComponent,
        props: {
          // Combine all containers (init, regular, ephemeral)
          containers: mockContainers,
        },
        sort: 2,
      },
      metrics: {
        component: SvelteComponent,
        sort: 12.490151,
        // metrics added by store.filterCallback
        props: { containers: podMetrics[0].containers },
      },
      restarts: 1,
      controller: 'DaemonSet',
      status: 'Running',
      node: '',
      age: { text: 'less than a minute', sort: 1721994792000 },
    } as unknown

    const urlAssertionMock = vi.fn().mockImplementation((url: string) => {
      console.log(url)
    })

    const closeMock = vi.fn()

    vi.stubGlobal(
      'EventSource',
      vi
        .fn()
        // metrics EventSource is created first in createStore()
        .mockImplementationOnce((url: string) => new MockEventSource(url, podMetrics, urlAssertionMock, closeMock))
        // next is the pods EventSource in store.start()
        .mockImplementation((url: string) => new MockEventSource(url, mockPods, urlAssertionMock, closeMock)),
    )

    // initialize store
    const store = createStore()
    const cleanup = store.start()

    // Assert the correct URLs were called by EventSources
    expect(urlAssertionMock).toHaveBeenCalledWith(`/api/v1/resources/workloads/podmetrics`)
    expect(urlAssertionMock).toHaveBeenCalledWith(`/api/v1/resources/workloads/pods`)

    // advance timers triggers the EventSource.onmessage callback in MockEventSource
    vi.advanceTimersByTime(500)
    store.subscribe((data) => {
      const { resource, table } = data[0]

      // Assert the data was passed from eventSource to transformer
      expect(resource).toEqual(mockPods[0])

      // Assert the data was transformed correctly
      expect(Object.keys(table)).toEqual(tableCols)
      expectEqualIgnoringFields(table, expectedTables, [
        'containers.component',
        'metrics.component',
        'creationTimestamp',
        'age.sort',
      ])

      expect(table.age?.text).toEqual('less than a minute')
    })

    // call store.stop()
    cleanup()
    // expect 2 because this.stopCallback was set to metricsEvents.close
    expect(closeMock).toHaveBeenCalledTimes(2)
  })

  vi.unstubAllGlobals()
  vi.useRealTimers()
})
