// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testCustomColumns, testDefaultColumns } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

// @todo: had to mock these components because they were causing major
// slow-downs in the transform process for vitest
vi.mock('./exemption-details/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

vi.mock('./exemption-matcher/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

vi.mock('./exemption-policies/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

suite('UDSExemptionTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  testDefaultColumns(Component, createStore, [
    ['name', 'emphasize'],
    ['namespace'],
    ['details'],
    ['matcher'],
    ['policies'],
    ['age'],
  ])

  testCustomColumns(Component, createStore)
})
