import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { NameAndDesc } from '$features/k8s/types'
import type { ComponentType } from 'svelte'

export function testK8sTableWithDefaults(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Ensure tooltip desc isn't empty
    const nameAndDesc: NameAndDesc = props.nameAndDesc as NameAndDesc
    expect(nameAndDesc.desc).toBeTruthy()

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
