// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, test, type Page } from '@playwright/test'

test.describe('DataTable', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')
  })

  test('filters rows when we click the namespace link in a row', async ({ page }) => {
    await page.getByRole('button', { name: 'podinfo' }).last().click()
    await checkTableHeaderResults(page, 1, 8)
    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'kube-system' })
    await checkTableHeaderResults(page, 1, 8)
  })

  test('filters rows when we select the namespace from the drop down option', async ({ page }) => {
    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'podinfo' })
    await checkTableHeaderResults(page, 1, 8)
    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'kube-system' })
    await checkTableHeaderResults(page, 1, 8)
  })

  test('filters rows when entering search values with "Anywhere" selected', async ({ page }) => {
    await page.getByTestId('datatable-search').fill('pepr')
    await checkTableHeaderResults(page, 4, 8)
    await page.getByTestId('datatable-search').fill('podinfo')
    await checkTableHeaderResults(page, 1, 8)
  })

  test('filters rows when entering search values with "Metadata" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Metadata').click()
    await page.getByTestId('datatable-search').fill('pepr')
    await checkTableHeaderResults(page, 4, 8)
  })

  test('filters rows when entering search values with "Name" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Name').click()
    await page.getByTestId('datatable-search').fill('pepr')
    await checkTableHeaderResults(page, 3, 8)
  })
})

// Checks to see if the table header results are as expected. For example: "Showing 1 of 8"
async function checkTableHeaderResults(page: Page, expectedActual: number, expectedTotal: number) {
  const tableHeaderResults = page.getByTestId('table-header-results')
  await tableHeaderResults.waitFor()
  const textContent = await tableHeaderResults.textContent()

  const regex = /showing (\d+) of (\d+)/
  const match = textContent?.match(regex)
  if (match) {
    const actual = parseInt(match[1], 10)
    const total = parseInt(match[2], 10)

    expect(actual).toBeGreaterThanOrEqual(expectedActual)
    expect(total).toBeGreaterThanOrEqual(expectedTotal)
  } else {
    throw new Error('Failed to extract numbers from table header results')
  }
}
