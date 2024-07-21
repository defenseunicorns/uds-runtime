import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { CommonRow, ResourceStoreInterface } from '$features/k8s/types'
import type { KubernetesObject } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'

export function testDefaultColumns<T extends CommonRow>(
  Component: ComponentType,
  createStore: () => ResourceStoreInterface<KubernetesObject, T>,
  defaultColumns: string[][],
) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props: {
        columns: defaultColumns,
        createStore: createStore,
      },
    })
  })
}

export function testCustomColumns<T extends CommonRow>(
  Component: ComponentType,
  createStore: () => ResourceStoreInterface<KubernetesObject, T>,
) {
  test('creates component with custom columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    const customColumns = [['blah'], ['blah2']]

    render(Component, {
      columns: customColumns,
    })

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props: {
        columns: customColumns,
        createStore: createStore,
      },
    })
  })
}
