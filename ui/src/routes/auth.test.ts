// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { authenticated } from '$features/auth/store'
import { beforeEach, describe, expect, vi } from 'vitest'

import { load } from './+layout'

// Mock stores
vi.mock('$features/k8s/namespaces/store', () => ({
  createStore: vi.fn(() => ({
    start: vi.fn(),
  })),
}))
vi.mock('$features/auth/store', () => ({
  authenticated: {
    set: vi.fn(),
    subscribe: vi.fn(() => () => {}),
    update: vi.fn(),
  },
}))

describe('load function', () => {
  // Mock fetch
  const fetchMock = vi.fn()
  global.fetch = fetchMock

  let mockUrl: URL

  beforeEach(() => {
    vi.clearAllMocks()

    // Reset URL mock
    mockUrl = new URL('https://example.com')
    Object.defineProperty(window, 'location', {
      value: { href: mockUrl.href },
      writable: true,
    })
  })

  test('successful local authentication with token in URL', async () => {
    // Set up URL with token
    mockUrl.searchParams.set('token', 'valid-token')
    Object.defineProperty(window, 'location', {
      value: { href: mockUrl.href },
      writable: true,
    })

    const mockUserData = {
      name: '',
      'preferred-username': '',
      group: '',
      'in-cluster-auth': false,
    }

    // Mock successful fetch response
    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(mockUserData),
    })

    const result = await load()

    // Verify fetch was called correctly
    expect(fetchMock).toHaveBeenCalledWith('/api/v1/auth?token=valid-token', {
      method: 'GET',
      headers: new Headers({
        'Content-Type': 'application/json',
      }),
    })

    // Verify store operations
    expect(result.namespaces.start).toHaveBeenCalled()
    expect(authenticated.set).toHaveBeenCalledWith(true)

    // Verify return value
    expect(result).toEqual({
      namespaces: expect.any(Object),
      userData: {
        name: '',
        preferredUsername: '',
        group: '',
        inClusterAuth: false,
      },
    })
  })

  test('successful in-cluster authentication (without token in URL)', async () => {
    const mockUserData = {
      name: 'Doug Unicorn',
      'preferred-username': 'doug@defenseunicorns.com',
      group: 'admin',
      'in-cluster-auth': true,
    }

    // Mock successful fetch response
    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve(mockUserData),
    })

    const result = await load()

    // Verify fetch was called with empty token
    expect(fetchMock).toHaveBeenCalledWith('/api/v1/auth', {
      method: 'GET',
      headers: new Headers({
        'Content-Type': 'application/json',
      }),
    })

    // Verify namespaces was called
    expect(result.namespaces.start).toHaveBeenCalled()
    expect(authenticated.set).toHaveBeenCalledWith(true)

    // Verify return value
    expect(result).toEqual({
      namespaces: expect.any(Object),
      userData: {
        name: 'Doug Unicorn',
        preferredUsername: 'doug@defenseunicorns.com',
        group: 'admin',
        inClusterAuth: true,
      },
    })
  })

  test('authentication failure with invalid token', async () => {
    mockUrl.searchParams.set('token', 'invalid-token')
    Object.defineProperty(window, 'location', {
      value: { href: mockUrl.href },
      writable: true,
    })

    fetchMock.mockResolvedValueOnce({
      ok: false,
      status: 401,
      statusText: 'Unauthorized',
    })

    const result = await load()

    // Verify namespaces wasn't started and authenticated state was set to false
    expect(result.namespaces.start).not.toHaveBeenCalled()
    expect(authenticated.set).toHaveBeenCalledWith(false)

    // Verify return value
    expect(result).toEqual({
      namespaces: expect.any(Object),
      userData: {
        name: '',
        preferredUsername: '',
        group: '',
        inClusterAuth: false,
      },
    })
  })

  test('network errors during authentication', async () => {
    const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})
    const networkError = new Error('Network error')

    fetchMock.mockRejectedValueOnce(networkError)

    const result = await load()

    // Verify error was logged
    expect(consoleSpy).toHaveBeenCalledWith('Load error:', expect.any(Error))

    // Verify namespaces wasn't started authenticated state was set to false
    expect(result.namespaces.start).not.toHaveBeenCalled()
    expect(authenticated.set).toHaveBeenCalledWith(false)

    // Verify return value
    expect(result).toEqual({
      namespaces: expect.any(Object),
      userData: {
        name: '',
        preferredUsername: '',
        group: '',
        inClusterAuth: false,
      },
    })

    consoleSpy.mockRestore()
  })

  test('malformed JSON in successful response', async () => {
    mockUrl.searchParams.set('token', 'valid-token')
    Object.defineProperty(window, 'location', {
      value: { href: mockUrl.href },
      writable: true,
    })

    fetchMock.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.reject(new Error('Invalid JSON')),
    })

    const result = await load()

    // Verify authenticated state was set to false
    expect(authenticated.set).toHaveBeenCalledWith(false)

    // Verify return value
    expect(result).toEqual({
      namespaces: expect.any(Object),
      userData: {
        name: '',
        preferredUsername: '',
        group: '',
        inClusterAuth: false,
      },
    })
  })
})
