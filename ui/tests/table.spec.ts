// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('DataTable', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')
  })

  test('filters rows when we click the namespace link in a row', async ({ page }) => {
    await page.getByRole('button', { name: 'podinfo' }).click()

    expect(await page.getByTestId('Pods-table-header-results').textContent()).toBe('(showing 1 of 7)')

    await page.locator('#stream').selectOption({ label: 'All Namespaces' })

    await page.getByRole('button', { name: 'kube-system' }).first().click()

    expect(await page.getByTestId('Pods-table-header-results').textContent()).toBe('(showing 3 of 7)')
  })

  test('filters rows when we select the namespace from the drop down option', async ({ page }) => {
    await page.locator('#stream').selectOption({ label: 'podinfo' })

    expect(await page.getByTestId('Pods-table-header-results').textContent()).toBe('(showing 1 of 7)')

    await page.locator('#stream').selectOption({ label: 'kube-system' })

    expect(await page.getByTestId('Pods-table-header-results').textContent()).toBe('(showing 3 of 7)')
  })
})
