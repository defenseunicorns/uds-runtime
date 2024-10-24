// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, test } from '@playwright/test'

// NOTE: these tests should only be run against the in-cluster environment
test.describe('in-cluster', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/')
  })

  test.describe('home page', async () => {
    test('user menu expands and collapses when clicked on', async ({ page }) => {
      const userMenuButton = page.getByText('Doug Unicorn')

      await userMenuButton.click()
      const signOutButton = page.getByText('Sign Out')

      await expect(signOutButton).toBeVisible()

      await userMenuButton.click()
      await expect(signOutButton).not.toBeVisible()
    })

    test('user menu expands and collapses when clicked away from', async ({ page }) => {
      const userMenuButton = page.getByText('Doug Unicorn')

      await userMenuButton.click()
      const signOutButton = page.getByText('Sign Out')

      await expect(signOutButton).toBeVisible()

      await page.getByRole('link', { name: 'Overview' }).click()
      await expect(signOutButton).not.toBeVisible()
    })
  })
})
