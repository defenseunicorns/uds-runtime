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

// Helper function to compare two objects while ignoring certain fields; can ignore nested fields (eg 'metadata.creationTimestamp')
export function expectEqualIgnoringFields<T>(actual: T, expected: T, fieldsToIgnore: string[]) {
  const expectedWithoutFields = { ...expected }
  const actualWithoutFields = { ...actual }

  fieldsToIgnore.forEach((field) => {
    if (field.includes('.')) {
      const paths = field.split('.')

      // Create temporary objects to traverse the object and delete the last field
      let tmpExpect = expectedWithoutFields
      let tmpActual = actualWithoutFields
      // Traverse the object to the second to last field (e.g. [field1, field2, field3] -> field2)
      for (let i = 0; i <= paths.length - 2; i++) {
        tmpExpect = tmpExpect[paths[i] as keyof typeof tmpExpect] as T
        tmpActual = tmpActual[paths[i] as keyof typeof tmpActual] as T

        // when second to last field reached (e.g. field2 of 3), delete the last field (e.g. delete {field3: value})
        if (i === paths.length - 2) {
          delete tmpExpect[paths[i + 1] as keyof typeof tmpExpect]
          delete tmpActual[paths[i + 1] as keyof typeof tmpActual]
        }
      }
    } else {
      delete expectedWithoutFields[field as keyof typeof expectedWithoutFields]
      delete actualWithoutFields[field as keyof typeof actualWithoutFields]
    }
  })

  expect(actualWithoutFields).toEqual(expectedWithoutFields)
}

export function testK8sResourceStore<R extends KubernetesObject, U extends CommonRow>(
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
    const storeStop = store.start()

    // Assert the correct URL was given to the EventSource
    expect(urlAssertionMock).toHaveBeenCalledTimes(1)
    expect(urlAssertionMock).toHaveBeenCalledWith(expectedUrl)

    // advance timers triggers the EventSource.onmessage callback in MockEventSource
    vi.advanceTimersByTime(500)
    store.subscribe((data) => {
      const { resource, table } = data[0]
      // Assert the data was passed from eventSource to transformer (avoid date time inconsistencies by ignoring creationTimestamp)
      expectEqualIgnoringFields(resource, mockData[0], ['metadata.creationTimestamp'])
      // Assert the data was transformed correctly to create the desired table rows
      expect(table).toEqual(expectedTable)
    })

    // call store.stop()
    storeStop()
    expect(closeMock).toHaveBeenCalledTimes(1)

    vi.unstubAllGlobals()
    vi.useRealTimers()
  })
}

export class MockEventSource {
  onmessage: (event: MessageEvent) => void | null = () => {}
  close: Mock<Procedure>

  constructor(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: KubernetesObject | any[],
    urlAssertionMock: Mock<Procedure>,
    closeMock: Mock<Procedure>,
  ) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)
    // Used for testing the EventSource was closed
    this.close = closeMock

    // After 500ms simulate message event and pass the data to the onmessage callback
    setTimeout(() => {
      const messageEvent = {
        data: JSON.stringify(data),
        origin: url,
        lastEventId: '',
      }

      // Check that onmessage has been set to handler by MockEventSource caller (eg. this.onmessage = ({data}) => {do stuff...})
      // then fire the onmessage callback with the messageEvent
      if (typeof this.onmessage === 'function') {
        this.onmessage(messageEvent as MessageEvent)
      }
    }, 500)
  }
}
