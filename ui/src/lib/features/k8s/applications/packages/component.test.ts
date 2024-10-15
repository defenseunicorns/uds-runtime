// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

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
      ['name', 'w-1/12'],
      ['version', ' w-1/12'],
      ['description', 'italic  w-5/12'],
      ['arch', ' w-1/12'],
      ['flavor', ' w-1/12'],
      ['components', ' w-2/12'],
      ['age', ' w-1/12'],
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
