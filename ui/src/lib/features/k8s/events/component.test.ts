// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import type { Resource } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('EventTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'Events'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['namespace'], ['age'], ['type'], ['reason'], ['object_kind'], ['object_name'], ['count']],
    resource,
  })

  testK8sTableWithCustomColumns(Component, { createStore, resource })
})
