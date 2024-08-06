/* eslint-disable @typescript-eslint/no-explicit-any */

import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { KubernetesObject } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'
import type { Mock } from 'vitest'
import type { CommonRow, ResourceWithTable } from './types'

// Vitest type redeclared cause it's not exported from vitest
type Procedure = (...args: any[]) => any

export function testK8sTableWithDefaults(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Ensure name and ooltip desc aren't empty
    const name: string = props.name as string
    expect(name).toBeTruthy()
    const description: string = props.description as string
    expect(description).toBeTruthy()

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

// TODO: look into deep copies since nested objects are still references and are getting mutated
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

// MockResourceStore is a lightweight mock of ResourceStore for testing URLs and data transformation callbacks
export class MockResourceStore {
  url: string
  data: KubernetesObject[]
  #tableCallback: (data: KubernetesObject[]) => ResourceWithTable<KubernetesObject, CommonRow>[]

  constructor(
    url: string,
    transform: <R extends KubernetesObject, U extends CommonRow>(resources: R[]) => ResourceWithTable<R, U>[],
    data: KubernetesObject[],
  ) {
    this.url = url
    this.data = data
    this.#tableCallback = transform
  }

  // call the given transform function, imitating the store.start()
  start() {
    return this.#tableCallback(this.data)
  }

  // added to satisfy store.sortByKey.bind(store),
  sortByKey() {}
}

export class MockEventSource {
  constructor(url: string, urlAssertionMock: Mock<Procedure>) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)
  }
  // satisfies store.stopCallback = metricsEvents.close.bind(metricsEvents)
  close() {}
}
