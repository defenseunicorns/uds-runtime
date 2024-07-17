// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'
import { render } from '@testing-library/svelte'

import * as components from '$components'
import Component from './component.svelte'
import { createStore, type Columns } from './store'

suite('DaemonsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    const defaultColumns = [
      ['name', 'emphasize'],
      ['namespace'],
      ['desired'],
      ['current'],
      ['ready'],
      ['up_to_date'],
      ['available'],
      ['node_selector'],
      ['age'],
    ]

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

  test('creates component with custom columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    const customColumns = [['blah'], ['blah2']] as unknown as Columns

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
})
