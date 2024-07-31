// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import type { NameAndDesc } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import Component from './component.svelte'
import { createStore } from './store'

suite('PersistentVolume Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'PersistentVolumes'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['storage_class'], ['capacity'], ['claim'], ['age'], ['status']],
    resource,
  })

  testK8sTableWithCustomColumns(Component, { createStore, resource })
})
