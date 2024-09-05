// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { apiAuthEnabled } from '$features/api-auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import { updateApiAuthEnabled } from '$lib/utils/helpers'
import { get } from 'svelte/store'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async ({ url }) => {
  const namespaces = createStore()

  const isAuthRoute = url.pathname.includes('/auth')
  updateApiAuthEnabled()

  if (!get(apiAuthEnabled) || !isAuthRoute) {
    namespaces.start()
  }
  return {
    namespaces,
  }
}
