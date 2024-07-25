// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { MockEventSource, testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import '@testing-library/jest-dom'
import Component from './component.svelte'
import { createStore } from './store'

suite('DeploymentTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['ready'], ['up_to_date'], ['available'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  test.only('createStore', async () => {
    vi.useFakeTimers()
    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementation((url) => new MockEventSource(url, mockData, urlAssertionMock)),
    )
    vi.setSystemTime(new Date('2024-07-25T16:10:22.000Z'))

    const mockData = [
      {
        metadata: { name: 'test', namespace: 'default', creationTimestamp: new Date().toString() },
        status: { readyReplicas: 1, replicas: 2, updatedReplicas: 1, conditions: [{ type: 'Available' }] },
      },
    ]

    const expectedTable = {
      name: 'test',
      namespace: 'default',
      creationTimestamp: new Date(),
      ready: '1 / 2',
      up_to_date: 1,
      available: 1,
      age: { text: 'less than a minute', sort: 1721923822000 },
    }

    const urlAssertionMock = vi.fn().mockImplementation((url) => {
      console.log(url)
    })

    const store = createStore()
    store.start()

    vi.advanceTimersByTime(500)
    expect(urlAssertionMock).toHaveBeenCalledWith('/api/v1/resources/workloads/deployments')
    store.subscribe((data: any) => {
      expect(data[0].resource).toEqual(mockData[0])
      expect(data[0].table).toEqual(expectedTable)
    })
  })
})
