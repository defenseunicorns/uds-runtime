// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { apiAuthEnabled } from '$features/api-auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import { updateApiAuthEnabled } from '$lib/utils/helpers'
import { get } from 'svelte/store'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async ({ url }) => {
  await updateApiAuthEnabled()

  const namespaces = createStore()
  const isInitialAPIAuthentication = url.pathname.includes('/auth')

  // start namespaces store if API auth is disabled or if doing a reload
  if (!get(apiAuthEnabled) || !isInitialAPIAuthentication) {
    namespaces.start()
  }
  return {
    namespaces,
  }
}
