// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
const BASE_URL = '/api/v1'

const headers = new Headers({
  'Content-Type': 'application/json',
})

export class TokenAuth {
  // wrapper for handling the request/response cycle.
  async request<T>(token: string): Promise<T> {
    const hasToken = token != ''
    const url = hasToken ? `${BASE_URL}/auth?token=${token}` : `${BASE_URL}/auth`

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

const tokenAuth = new TokenAuth()
const Auth = {
  connect: async (token: string) => {
    return await tokenAuth.request(token)
  },
}

export { Auth }