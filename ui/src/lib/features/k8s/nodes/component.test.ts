// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import type { NameAndDesc } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '../test-helper'
import Component from './component.svelte'
import { createStore } from './store'

suite('NodeTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'Nodes'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['status'], ['roles'], ['taints'], ['version'], ['age']],
    isNamespaced: false,
    resource,
  })

  testK8sTableWithCustomColumns(Component, { createStore, isNamespaced: false, resource })
})
