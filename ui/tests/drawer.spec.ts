// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, test } from '@playwright/test'

test.describe('Drawer', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')

    page
      .locator('.table .tr')
      .filter({ hasText: /^podinfo-/ })
      .click()
  })

  test.describe('is opened when clicking on a table row and', async () => {
    test('will display Metadata details', async ({ page }) => {
      const drawerEl = page.getByTestId('drawer')

      await expect(drawerEl).toBeVisible()
      await expect(drawerEl.getByText('Created')).toBeVisible()
      await expect(drawerEl.getByText('Name', { exact: true })).toBeVisible()
      await expect(drawerEl.getByText('Annotations')).toBeVisible()
      await expect(drawerEl.getByText('podinfo', { exact: true })).toBeVisible()
    })

    test('will display Events details', async ({ page }) => {
      const drawerEl = page.getByTestId('drawer')

      await expect(drawerEl).toBeVisible()
      await drawerEl.getByRole('button', { name: 'Events' }).click()

      await expect(drawerEl.getByText('Created container podinfo')).toBeVisible()
    })

    test('will display YAML details', async ({ page }) => {
      const drawerEl = page.getByTestId('drawer')
      const labelName = await drawerEl.locator(':text("app.kubernetes.io/name:")').textContent()
      const podID = page.url().split('/').pop()

      await drawerEl.getByRole('button', { name: 'YAML' }).click()
      await expect(drawerEl.getByText('namespace:')).toBeVisible()

      // Ensure label:app matches pod info
      expect(labelName).toEqual(`app.kubernetes.io/name: podinfo`)
      // Ensure metadata:uid matches url pod:id
      await expect(drawerEl.locator(`:text("uid: ${podID}")`)).toBeVisible()
    })
  })
})
