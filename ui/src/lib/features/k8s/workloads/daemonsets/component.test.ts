// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('DaemonsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['desired'],
      ['current'],
      ['ready'],
      ['up_to_date'],
      ['available'],
      ['node_selector'],
      ['age'],
    ],
  })

  testK8sTableWithCustomColumns(Component, { createStore })
})