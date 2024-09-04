// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { createStore } from '$features/k8s/namespaces/store'
import { fetechAPIAuthStatus } from '$lib/utils/helpers'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()

  //Check if apiAuthEnabled
  const envVars = await fetechAPIAuthStatus()
  // API Auth is only disabled when API_AUTH_DISABLED is set to 'true'
  const apiAuthEnabled = envVars.API_AUTH_DISABLED?.toLowerCase() !== 'true'
  // namespaces.start() called in auth page when apiAuthEnabled
  if (!apiAuthEnabled) {
    namespaces.start()
  }

  return {
    namespaces,
  }
}
