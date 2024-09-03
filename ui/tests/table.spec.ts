// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('DataTable', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')
  })

  test('filters rows when we click the namespace link in a row', async ({ page }) => {
    await page.getByRole('button', { name: 'podinfo' }).click()

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')

    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'All Namespaces' })

    await page.getByRole('button', { name: 'kube-system' }).first().click()

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
  })

  test('filters rows when we select the namespace from the drop down option', async ({ page }) => {
    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'podinfo' })

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')

    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'kube-system' })

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
  })

  test('filters rows when entering search values with "Anywhere" selected', async ({ page }) => {
    await page.getByTestId('datatable-search').fill('pepr')

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 4 of 8)')

    await page.getByTestId('datatable-search').fill('podinfo')

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')
  })

  test('filters rows when enterting search values with "Metadata" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Metadata').click()

    await page.getByTestId('datatable-search').fill('pepr')

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 4 of 8)')
  })

  test('filters rows when enterting search values with "Name" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Name').click()

    await page.getByTestId('datatable-search').fill('pepr')

    expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
  })
})
