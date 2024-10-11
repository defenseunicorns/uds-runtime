// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { authenticated } from '$features/auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import { Auth } from '$lib/utils/token-auth'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()

  const url = new URL(window.location.href)
  const token = url.searchParams.get('token') || ''

  // validate token
  if (await Auth.connect(token)) {
    namespaces.start()
    authenticated.set(true)
  } else {
    authenticated.set(false)
  }
  return {
    namespaces,
  }
}
