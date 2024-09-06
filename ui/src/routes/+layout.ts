// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { apiAuthEnabled } from '$features/api-auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import { updateApiAuthEnabled } from '$lib/utils/helpers'
import { get } from 'svelte/store'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()

  await updateApiAuthEnabled()

  // check session storage for auth status
  const isAuthEnabled = JSON.parse(sessionStorage.getItem('apiAuthEnabled')!)
  const isAuthenticated = JSON.parse(sessionStorage.getItem('authenticated')!)

  // use auth status to determine if this is a page reload vs initial load
  const isReload: boolean = isAuthEnabled !== null && isAuthenticated !== null

  // start namespaces store if API auth is disabled or if doing a a reload
  if (!get(apiAuthEnabled) || isReload) {
    namespaces.start()
  }
  return {
    namespaces,
  }
}
