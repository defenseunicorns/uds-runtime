// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { beforeEach, vi } from 'vitest'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Application Endpoints'
  const description = 'Application Endpoints are exposed by the UDS Operator via a UDS Package CR.'

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['url', 'emphasize', 'link-external'], ['namespace'], ['status'], ['age']],
    name,
    description,
    disableRowClick: true,
  })

  testK8sTableWithCustomColumns(Component, {
    createStore,
    name,
    description,
    disableRowClick: true,
  })
})
