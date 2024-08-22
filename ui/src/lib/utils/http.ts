// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors
import { apiAuthEnabled } from '$lib/features/api-auth/store'

const BASE_URL = '/api/v1'

const headers = new Headers({
  'Content-Type': 'application/json',
})

interface APIRequest<T> {
  path: string
  method: string
  body?: T
}

type ResponseType = 'json' | 'boolean' | 'text'

export class HTTP {
  constructor() {
    const token = sessionStorage.getItem('token') || ''
    if (!token) {
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

  head(path: string) {
    return this.request<boolean>({ path, method: 'HEAD' }, 'boolean')
  }

  // Private wrapper for handling the request/response cycle.
  private async request<T>(req: APIRequest<T>, responseType: ResponseType = 'json'): Promise<T> {
    let token = sessionStorage.getItem('token')
    const url = BASE_URL + req.path + (token ? `?token=${token}` : '')
    const payload: RequestInit = { method: req.method, headers }

    try {
      // Actually make the request
      const response = await fetch(url, payload)

      // Head just returns response.ok
      if (req.method === 'HEAD') {
        return response.ok as T
      }

      // If the response is not OK, throw an error.
      if (!response.ok) {
        // all API errors should be 500s w/ a text body
        const errMessage = await response.text()
        throw new Error(errMessage)
      }

      switch (responseType) {
        case 'boolean':
          return response.ok as T
        case 'text':
          return (await response.text()) as T
        default:
          return (await response.json()) as T
      }

      // Return the response as the expected type
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
    return await http.head('/')
  },
}

export async function updateApiAuthEnabled() {
  try {
    const response = await fetch('/config')
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const envVars = await response.json()
    apiAuthEnabled.set(envVars.VITE_API_AUTH?.toLowerCase() === 'true')
  } catch (e) {
    console.error('Failed to fetch config:', e)
  }
}

export { Auth }
