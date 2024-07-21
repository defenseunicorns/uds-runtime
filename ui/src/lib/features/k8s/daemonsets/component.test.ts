// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testCustomColumns, testDefaultColumns } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('DaemonsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testDefaultColumns(Component, createStore, [
    ['name', 'emphasize'],
    ['namespace'],
    ['desired'],
    ['current'],
    ['ready'],
    ['up_to_date'],
    ['available'],
    ['node_selector'],
    ['age'],
  ])

  testCustomColumns(Component, createStore)
})
