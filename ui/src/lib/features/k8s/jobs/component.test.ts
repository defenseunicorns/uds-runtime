// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testCustomColumns, testDefaultColumns } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testDefaultColumns(Component, createStore, [
    ['name', 'emphasize'],
    ['namespace'],
    ['completions'],
    ['durations'],
    ['age'],
  ])

  testCustomColumns(Component, createStore)
})
