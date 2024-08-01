// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'
import type { NameAndDesc } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import Component from './component.svelte'
import { createStore } from './store'

suite('StorageClass Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const resourceName = 'StorageClasses'
  const nameAndDesc: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['provisioner'], ['reclaim_policy'], ['default'], ['age']],
    nameAndDesc,
  })

  testK8sTableWithCustomColumns(Component, { createStore, nameAndDesc })
})
