// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

export async function tokenAuth(token: string): Promise<boolean> {
  const hasToken = token != ''
  const BASE_URL = '/api/v1'
  const headers = new Headers({
    'Content-Type': 'application/json',
  })
  const url = hasToken ? `${BASE_URL}/auth?token=${token}` : `${BASE_URL}/auth`
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
