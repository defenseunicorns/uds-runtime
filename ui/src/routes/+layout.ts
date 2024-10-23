// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { authenticated } from '$features/auth/store'
import { createStore } from '$features/k8s/namespaces/store'
import type { UserData } from '$features/navigation/types'

export const ssr = false

interface AuthResponse {
  authenticated: boolean
  userData: UserData | null
}

// auth function that returns both auth status and user data
async function auth(token: string): Promise<AuthResponse> {
  const baseURL = '/api/v1'
  const headers = new Headers({
    'Content-Type': 'application/json',
  })

  try {
    const response = await fetch(`${baseURL}/auth?token=${token}`, {
      method: 'GET',
      headers,
    })
    if (response.ok) {
      const json = await response.json()
      return {
        authenticated: response.ok,
        userData: {
          name: json['name'],
          preferredUsername: json['preferred-username'],
          group: json['group'],
          inClusterAuth: json['in-cluster-auth'],
        },
      }
    } else {
      return {
        authenticated: false,
        userData: null,
      }
    }
  } catch (error) {
    console.error('Authentication error:', error)
    throw error // Let the caller handle the error
  }
}

// load namespace and auth data before rendering the app
export const load = async () => {
  const namespaces = createStore()
  const url = new URL(window.location.href)
  const localAuthToken = url.searchParams.get('token') || ''
  let userData: UserData | null = null

  try {
    const authResult = await auth(localAuthToken)

    if (authResult.authenticated) {
      namespaces.start()
      authenticated.set(true)
      userData = authResult.userData
    } else {
      authenticated.set(false)
    }
  } catch (error) {
    console.error('Load error:', error)
    authenticated.set(false)
  }

  return {
    namespaces,
    userData,
  }
}
