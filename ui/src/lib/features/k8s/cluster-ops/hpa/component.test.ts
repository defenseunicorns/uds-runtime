// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import type { NameAndDesc } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import Component from './component.svelte'
import { createStore } from './store'

suite('PriorityClassesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'HorizontalPodAutoscalers'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['namespace'],
      ['metrics'],
      ['min_pods'],
      ['max_pods'],
      ['replicas'],
      ['age'],
      ['status'],
    ],
    resource,
  })

  testK8sTableWithCustomColumns(Component, { createStore, resource })
})
