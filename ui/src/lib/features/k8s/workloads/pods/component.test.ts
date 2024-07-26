// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import { MockEventSource, testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
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
          containerStatuses: [
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
          ],
          initContainerStatuses: [
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
          phase: 'Running',
        },
      },
    ]
    const podMetrics = [
      {
        containers: [
          { name: 'speaker', usage: { cpu: '2336411n', memory: '20236Ki' } },
          { name: 'frr', usage: { cpu: '753805n', memory: '21640Ki' } },
          { name: 'frr-metrics', usage: { cpu: '9399935n', memory: '7264Ki' } },
          { name: 'reloader', usage: { cpu: '0', memory: '792Ki' } },
        ],
        metadata: {
          creationTimestamp: '2024-07-25T19:33:42Z',
          name: 'metallb-speaker-6nl62',
          namespace: 'uds-dev-stack',
        },
      },
    ]

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

    vi.useFakeTimers()
    vi.setSystemTime(new Date())

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
    expect(urlAssertionMock).toHaveBeenCalledWith(`/api/v1/resources/workloads/pods`)
    expect(urlAssertionMock).toHaveBeenCalledWith(`/api/v1/resources/workloads/podmetrics`)

    vi.advanceTimersByTime(500)
    store.subscribe((data) => {
      const { resource, table } = data[0]
      // Assert the data was passed from eventSource to transformer
      expect(resource).toEqual(mockPods[0])

      // Assert the data was transformed correctly
      expect(Object.keys(table)).toEqual(tableCols)
      expect(table.name).toEqual(mockPods[0].metadata.name)
      expect(table.containers.props.containers).toEqual([
        ...mockPods[0].status.containerStatuses,
        ...mockPods[0].status.initContainerStatuses,
      ])
      expect(table.restarts).toEqual(1)
      expect(table.namespace).toEqual(mockPods[0].metadata.namespace)
      // expect(table.creationTimestamp).toEqual(mockPods[0].metadata.creationTimestamp)
      expect(table.controller).toEqual(mockPods[0].metadata.ownerReferences[0].kind)
      expect(table.status).toEqual(mockPods[0].status.phase)
      expect(table.node).toEqual('')
      expect(table.age?.text).toEqual('less than a minute')
      // Assert filterCallback was called and metrics were added to the table
      expect(table.metrics.props.containers).toEqual(podMetrics[0].containers)
    })

    // call store.stop()
    cleanup()
    // expect 2 because this.stopCallback is set to metricsEvents.close
    expect(closeMock).toHaveBeenCalledTimes(2)
  })

  vi.unstubAllGlobals()
  vi.useRealTimers()
})
