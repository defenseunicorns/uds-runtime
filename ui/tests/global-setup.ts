// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { chromium, expect, type FullConfig } from '@playwright/test'

// Before running the in-cluster e2e tests, log into the app using
// the doug keycloak user and save the storage state to a file.
async function globalSetup(config: FullConfig) {
  const { baseURL, storageState } = config.projects[0].use
  const browser = await chromium.launch()
  const page = await browser.newPage()

  const maxRetries = 3
  let attempt = 0

  // logging in is flaky (especially in CI), so we retry a few times
  while (attempt < maxRetries) {
    try {
      // Navigate to the page
      await page.goto(baseURL!, {
        waitUntil: 'networkidle',
        timeout: 30000, // 30 seconds timeout
      })

      // Wait for a reliable indicator that the page is ready
      await Promise.all([
        page.waitForLoadState('domcontentloaded'),
        page.waitForLoadState('networkidle'),
        page.waitForSelector('#username', {
          state: 'visible',
          timeout: 10000,
        }),
      ])

      // If we reach here, the page loaded successfully
      console.log('Page loaded successfully')
      break
    } catch (error) {
      attempt++
      console.log(`Page load attempt ${attempt} failed: ${error.message}`)
      await page.screenshot({ path: `page-load-failure-${attempt}.png` })

      if (attempt === maxRetries) {
        throw new Error(`Failed to load page after ${maxRetries} attempts`)
      }

      // Wait before retrying
      await page.waitForTimeout(2000)
    }
  }

  // perform login
  await page.getByLabel('username').fill('doug')
  await page.getByLabel('password').fill('unicorn123!@#UN')
  await page.getByText('Log In').click()

  // ensure we are logged in and the home page's user menu is visible
  await expect(page.getByText('Doug Unicorn')).toBeVisible()
  console.log('Logged in successfully') // useful to know when globalSetup has finished in CI
  await page.context().storageState({ path: storageState as string })
  await browser.close()
}

export default globalSetup
