// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('Drawer', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')

    await page.getByRole('row').nth(1).click()
  })

  test.describe('is opened when clicking on a table row and', async () => {
    test('will display Metadata details', async ({ page }) => {
      const drawerEl = page.getByTestId('drawer')

      await expect(drawerEl).toBeVisible()
      await expect(drawerEl.getByText('Created')).toBeVisible()
      await expect(drawerEl.getByText('Name', { exact: true })).toBeVisible()
      await expect(drawerEl.getByText('Namespace')).toBeVisible()
      await expect(drawerEl.getByText('istio-admin-gateway')).toBeVisible()
    })

    test('will display YAML details', async ({ page }) => {
      const drawerEl = page.getByTestId('drawer')

      await drawerEl.getByRole('button', { name: 'YAML' }).click()
      await expect(drawerEl.getByText('namespace:')).toBeVisible()
      await expect(drawerEl.getByText('istio-admin-gateway', { exact: true })).toBeVisible()
    })
  })
})
