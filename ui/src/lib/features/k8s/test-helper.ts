import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { KubernetesObject } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'
import type { CommonRow, ResourceWithTable } from './types'

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

// export const TestCreationTimestamp = '2024-07-25T16:11:22.000Z'

export class MockResourceStore {
  url: string
  #tableCallback: (data: KubernetesObject[]) => ResourceWithTable<KubernetesObject, CommonRow>[]
  data: KubernetesObject[]
  constructor(
    url: string,
    transform: <R extends KubernetesObject, U extends CommonRow>(resources: R[]) => ResourceWithTable<R, U>[],
    data: KubernetesObject[],
  ) {
    // Used for testing the correct URL was passed to the EventSource
    this.url = url
    this.#tableCallback = transform
    this.data = data
  }

  start() {
    return this.#tableCallback(this.data)
  }

  sortByKey() {}
}
