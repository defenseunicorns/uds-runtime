// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { get } from 'svelte/store'

import { apiAuthEnabled, authenticated } from '$features/api-auth/store'

export const stringToSnakeCase = (name: string) => name.split(' ').join('-').toLocaleLowerCase()

export async function fetchAPIAuthStatus(): Promise<Record<string, string>> {
  const response = await fetch('/auth-status')
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }
  const envVars = await response.json()
  return envVars
}

export function createEventSource(path: string): EventSource {
  if (get(apiAuthEnabled)) {
    const apiToken: string = sessionStorage.getItem('token') ?? ''
    // Check if the URL already contains a '?' for urls with multiple search params
    const separator = path.includes('?') ? '&' : '?'
    const url = `${path}${separator}token=${apiToken}`
    return new EventSource(url)
  }
  return new EventSource(path)
}

// Can't live in api-auth.ts because of sessionStorage usage in file, preventing this function
// from being used in load functions
export async function updateApiAuthEnabled() {
  // check session storage for auth status
  const isAuthEnabled = JSON.parse(sessionStorage.getItem('apiAuthEnabled')!)
  const isAuthenticated = JSON.parse(sessionStorage.getItem('authenticated')!)

  const isReload: boolean = isAuthEnabled !== null && isAuthenticated !== null

  // if this is not a page reload and we haven't checked the API auth status yet, fetch it
  if (!isReload && get(apiAuthEnabled) === null) {
    const envVars = await fetchAPIAuthStatus()
    // API Auth is only disabled when API_AUTH_DISABLED is set to 'true'
    const auth: boolean = envVars.API_AUTH_DISABLED?.toLowerCase() !== 'true'
    // set store and session storage with value representing if auth is enabled or disabled
    apiAuthEnabled.set(auth)
    sessionStorage.setItem('apiAuthEnabled', JSON.stringify(auth))
  } else {
    // set the store values with the values grabbed from the session storage
    apiAuthEnabled.set(isAuthEnabled)
    authenticated.set(isAuthenticated)
  }
}
