// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors
const BASE_URL = '/api/v1'

const headers = new Headers({
  'Content-Type': 'application/json',
})

export class APIAuth {
  // wrapper for handling the request/response cycle.
  async request<T>(token: string): Promise<T> {
    const hasToken = token != '' ? true : false
    const url = hasToken ? `${BASE_URL}/api-auth?token=${token}` : `${BASE_URL}/api-auth`

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

const apiAuth = new APIAuth()
const Auth = {
  connect: async (token: string) => {
    return await apiAuth.request(token)
  },
}

export { Auth }
