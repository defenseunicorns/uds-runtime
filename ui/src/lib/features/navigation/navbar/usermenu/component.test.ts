// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { cleanup, fireEvent, render, screen } from '@testing-library/svelte'
import { goto } from '$app/navigation'
import { UserMenu } from '$features/navigation'
import type { UserData } from '$features/navigation/types'
import { afterEach, beforeEach, describe, expect, vi } from 'vitest'

// Mock the navigation module
vi.mock('$app/navigation', () => ({
  goto: vi.fn(),
}))

describe('UserMenu', () => {
  const mockUserData: UserData = {
    name: 'Doug Unicorn',
    preferredUsername: 'doug@example.com',
    group: 'Product',
    inClusterAuth: true,
  }

  beforeEach(() => {
    // Reset all mocks before each test
    vi.clearAllMocks()
  })

  afterEach(() => {
    cleanup()
  })

  test('renders with initial closed state', () => {
    render(UserMenu, { props: { userData: mockUserData } })

    // Check if the main button is rendered with user name
    expect(screen.getByText('Doug Unicorn')).toBeInTheDocument()

    // Verify dropdown is not visible initially
    expect(screen.queryByText('Sign Out')).not.toBeInTheDocument()
  })

  test('opens dropdown when clicked', async () => {
    render(UserMenu, { props: { userData: mockUserData } })

    const button = screen.getByText('Doug Unicorn')
    await fireEvent.click(button)

    // Check if dropdown content is visible
    expect(screen.getByText('Sign Out')).toBeInTheDocument()
    expect(screen.getByText('doug@example.com')).toBeInTheDocument()
    expect(screen.getByText('Product')).toBeInTheDocument()
  })

  test('navigates to logout page when signing out', async () => {
    render(UserMenu, { props: { userData: mockUserData } })

    // Open dropdown
    const button = screen.getByText('Doug Unicorn')
    await fireEvent.click(button)

    // Click sign out
    const signOutButton = screen.getByText('Sign Out')
    await fireEvent.click(signOutButton)

    // Verify navigation
    expect(goto).toHaveBeenCalledWith('/logout')
  })

  it('keeps dropdown open when clicking inside', async () => {
    render(UserMenu, { props: { userData: mockUserData } })

    // Open dropdown
    const button = screen.getByText('Doug Unicorn')
    await fireEvent.click(button)

    // Click on an element inside the dropdown
    const signOutButton = screen.getByText('Sign Out')
    await fireEvent.click(signOutButton)

    // Verify dropdown remains visible
    expect(screen.getByText('Sign Out')).toBeInTheDocument()
  })
})
