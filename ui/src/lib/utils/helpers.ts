// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { apiAuthEnabled } from '$features/api-auth/store'
import { get } from 'svelte/store'

export const stringToSnakeCase = (name: string) => name.split(' ').join('-').toLocaleLowerCase()

export async function fetechAPIAuthStatus(): Promise<Record<string, string>> {
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
