import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { V1Deployment as Resource } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'
import type { Mock } from 'vitest'
import type { ResourceStoreInterface } from './types'

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

export function testK8sResourceStore(
  resource: string,
  mockData: Resource[],
  expectedTable: Record<string, unknown>,
  expectedUrl: string,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  createStore: () => ResourceStoreInterface<Resource, any>,
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

    vi.stubGlobal(
      'EventSource',
      vi
        .fn()
        .mockImplementation(
          (url: string) => new MockEventSource(url, mockData as unknown as Resource[], urlAssertionMock),
        ),
    )

    // initialize store
    const store = createStore()
    store.start()

    expect(urlAssertionMock).toHaveBeenCalledWith(expectedUrl)

    vi.advanceTimersByTime(500)
    store.subscribe((data) => {
      // todo: fix can't compare metadata as a whole because mockData[0].metadata.creationTimestamp is not wrapped in ""
      expect(data[0].resource.metadata?.name).toEqual(mockData[0].metadata?.name)
      expect(data[0].resource['status']).toEqual(mockData[0]['status'])
      expect(data[0].table).toEqual(expectedTable)
    })
  })

  vi.unstubAllGlobals()
  vi.useRealTimers()
}

// Mocking EventSource globally
class MockEventSource {
  onmessage: (event: MessageEvent) => void | null = () => {}
  constructor(url: string, data: Resource[], urlAssertionMock: Mock<Procedure>) {
    urlAssertionMock(url)

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
