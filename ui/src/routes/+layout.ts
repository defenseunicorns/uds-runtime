// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { authenticated } from '$features/auth/store'
import { createStore } from '$features/k8s/namespaces/store'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()

  const url = new URL(window.location.href)
  const token = url.searchParams.get('token') || ''

  // validate token
  if (await tokenAuth(token)) {
    namespaces.start()
    authenticated.set(true)
  } else {
    authenticated.set(false)
  }
  return {
    namespaces,
  }
}

// tokenAuth is a helper function that checks if a token is valid for local auth
async function tokenAuth(token: string): Promise<boolean> {
  const hasToken = token != ''
  const baseURL = '/api/v1'
  const headers = new Headers({
    'Content-Type': 'application/json',
  })
  const url = hasToken ? `${baseURL}/auth?token=${token}` : `${baseURL}/auth`
  const payload: RequestInit = { method: 'HEAD', headers }

  try {
    // Actually make the request
    const response = await fetch(url, payload)
    return response.ok
  } catch (e) {
    // Something went wrong--abort the request.
    console.error(e)
    return Promise.reject(e)
  }
}
