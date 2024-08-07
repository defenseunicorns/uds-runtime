// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('Drawer', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')
  })

  test.describe('is opened when clicking on a table row and', async () => {
    test('will display metadata details', async ({ page }) => {
      await expect(page.getByTestId('drawer')).not.toBeVisible()

      await page.getByRole('row').nth(1).click()
      await expect(page.getByTestId('drawer')).toBeVisible()

      const drawerEl = page.getByTestId('drawer')

      await expect(drawerEl.getByText('Created')).toBeVisible()
      await expect(drawerEl.$(/^Name$/)).toBeVisible()
      await expect(drawerEl.getByText('Namespace')).toBeVisible()
    })
  })
})
