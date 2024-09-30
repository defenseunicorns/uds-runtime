// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('Drawer', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')

    await page.getByRole('cell', { name: 'podinfo-' }).click()
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
