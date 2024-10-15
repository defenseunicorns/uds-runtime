// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { authenticated } from '$features/api-auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import { Auth } from '$lib/utils/api-auth'

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
