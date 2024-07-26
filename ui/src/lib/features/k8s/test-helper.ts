import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { KubernetesObject } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'
import type { Mock } from 'vitest'
import type { CommonRow, ResourceStoreInterface } from './types'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type Procedure = (...args: any[]) => any

export function testK8sTableWithDefaults(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

export function testK8sTableWithCustomColumns(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with custom columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    props.columns = [['blah'], ['blah2']]

    render(Component, { columns: props.columns })

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

type Resource = KubernetesObject & {
  status?: Record<string, unknown>
  spec?: Record<string, unknown>
}

export function testK8sResourceStore<R extends Resource, U extends CommonRow>(
  resource: string,
  mockData: R[],
  expectedTable: Record<string, unknown>,
  expectedUrl: string,
  createStore: () => ResourceStoreInterface<R, U>,
) {
  test(`createStore for ${resource}`, async () => {
    vi.useFakeTimers()
    vi.setSystemTime(new Date('2024-07-25T16:10:22.000Z'))

    // set creationTimestamps after mocking the date
    mockData[0].metadata!.creationTimestamp = new Date()
    expectedTable.creationTimestamp = new Date()

    const urlAssertionMock = vi.fn().mockImplementation((url: string) => {
      console.log(url)
    })

    const closeMock = vi.fn()

    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementation((url: string) => new MockEventSource(url, mockData, urlAssertionMock, closeMock)),
    )

    // initialize store
    const store = createStore()
    const cleanup = store.start()

    expect(urlAssertionMock).toHaveBeenCalledTimes(1)
    expect(urlAssertionMock).toHaveBeenCalledWith(expectedUrl)

    vi.advanceTimersByTime(500)
    store.subscribe((data) => {
      const { resource, table } = data[0]
      // todo: fix can't compare metadata as a whole because mockData[0].metadata.creationTimestamp is not wrapped in ""
      expect(resource.metadata?.name).toEqual(mockData[0].metadata?.name)
      expect(resource.status).toEqual(mockData[0].status)
      expect(table).toEqual(expectedTable)
    })

    // call the cleanup function
    cleanup()
    expect(closeMock).toHaveBeenCalledTimes(1)

    vi.unstubAllGlobals()
    vi.useRealTimers()
  })
}

// Mocking EventSource globally
export class MockEventSource {
  onmessage: (event: MessageEvent) => void | null = () => {}
  close: Mock<Procedure>
  constructor(url: string, data: Resource[], urlAssertionMock: Mock<Procedure>, closeMock: Mock<Procedure>) {
    urlAssertionMock(url)
    this.close = closeMock

    setTimeout(() => {
      const messageEvent = {
        data: JSON.stringify(data),
        origin: url,
        lastEventId: '',
      }

      if (typeof this.onmessage === 'function') {
        this.onmessage(messageEvent as MessageEvent)
      }
    }, 500)
  }
}
