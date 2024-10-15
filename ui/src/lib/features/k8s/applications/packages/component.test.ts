// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import '@testing-library/jest-dom'

import { testK8sTableWithCustomColumns, testK8sTableWithDefaults } from '$features/k8s/test-helper'

import Component from './component.svelte'
import { createStore } from './store'

suite('LimitRangesTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Application Packages'
  const description = 'Packages are either UDS or Zarf packages that are deployed in the cluster.'

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize'],
      ['version'],
      ['description', 'italic'],
      ['arch'],
      ['flavor'],
      ['components'],
      ['age'],
    ],
    name,
    description,

    isNamespaced: false,
    disableRowClick: true,
  })

  testK8sTableWithCustomColumns(Component, {
    createStore,
    name,
    description,
    isNamespaced: false,
    disableRowClick: true,
  })
})
