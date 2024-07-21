// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testCustomColumns, testDefaultColumns } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('DeploymentTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testDefaultColumns(Component, createStore, [
    ['name', 'emphasize'],
    ['namespace'],
    ['ready'],
    ['up_to_date'],
    ['available'],
    ['age'],
  ])

  testCustomColumns(Component, createStore)
})
