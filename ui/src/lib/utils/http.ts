// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors
import { apiAuthEnabled } from '$lib/features/api-auth/store'
import { fetchConfig } from '$lib/utils/helpers'
import { get } from 'svelte/store'

const BASE_URL = '/api/v1'

const headers = new Headers({
  'Content-Type': 'application/json',
})

export class HTTP {
  constructor() {
    const token = sessionStorage.getItem('token') || ''
    const isApiAuthEnabled = get(apiAuthEnabled)
    if (!token && isApiAuthEnabled) {
      this.invalidateAuth()
    }
  }

  // Updates the internal token used for authentication.
  updateToken(token: string) {
    sessionStorage.setItem('token', token)
  }

  private invalidateAuth() {
    sessionStorage.removeItem('token')
    if (location.pathname !== '/auth') {
      location.pathname = '/auth'
    }
  }

  // wrapper for handling the request/response cycle.
  async request<T>(): Promise<T> {
    const token = sessionStorage.getItem('token')
    const url = BASE_URL + '/' + (token ? `?token=${token}` : '')
    const payload: RequestInit = { method: 'HEAD', headers }

    try {
      // Actually make the request
      const response = await fetch(url, payload)
      return response.ok as T
    } catch (e) {
      // Something went really wrong--abort the request.
      console.error(e)
      return Promise.reject(e)
    }
  }
}

const http = new HTTP()
const Auth = {
  connect: async (token: string) => {
    if (!token) {
      return false
    }
    http.updateToken(token)
    return await http.request()
  },
}

export async function updateApiAuthEnabled() {
  const envVars = await fetchConfig()
  apiAuthEnabled.set(envVars.API_AUTH_ENABLED?.toLowerCase() === 'true')
}

export { Auth }
