// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import type { Resource } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import Component from './component.svelte'
import { createStore } from './store'

suite('CronjobTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'CronJobs'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['schedule'], ['suspend'], ['last_scheduled'], ['age']],
    resource,
  })

  testK8sTableWithCustomColumns(Component, { createStore, resource })
})
