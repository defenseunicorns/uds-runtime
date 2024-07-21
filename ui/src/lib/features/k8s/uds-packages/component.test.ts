// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import { testCustomColumns, testDefaultColumns } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testDefaultColumns(Component, createStore, [
    ['name', 'emphasize'],
    ['namespace'],
    ['status'],
    ['endpoints'],
    ['monitors'],
    ['ssoClients'],
    ['networkPolicies'],
    ['age'],
  ])

  testCustomColumns(Component, createStore)
})
