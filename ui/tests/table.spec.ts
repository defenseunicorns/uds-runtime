// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, test } from '@playwright/test'

const isInCluster = process.env.RUNTIME_ENVIRONMENT === 'in-cluster'

test.describe('DataTable', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/workloads/pods')
  })

  test('filters rows when we click the namespace link in a row', async ({ page }) => {
    await page.getByRole('button', { name: 'podinfo' }).last().click()

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')
    }

    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'All Namespaces' })

    await page.getByRole('button', { name: 'kube-system' }).first().click()

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
    }
  })

  test('filters rows when we select the namespace from the drop down option', async ({ page }) => {
    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'podinfo' })

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')
    }

    await page.getByTestId('table-filter-namespace-select').selectOption({ label: 'kube-system' })

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
    }
  })

  test('filters rows when entering search values with "Anywhere" selected', async ({ page }) => {
    await page.getByTestId('datatable-search').fill('pepr')

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 6 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 4 of 8)')
    }

    await page.getByTestId('datatable-search').fill('podinfo')

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 8 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 1 of 8)')
    }
  })

  test('filters rows when entering search values with "Metadata" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Metadata').click()

    await page.getByTestId('datatable-search').fill('pepr')

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 6 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 4 of 8)')
    }
  })

  test('filters rows when entering search values with "Name" selected', async ({ page }) => {
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByLabel('Name').click()

    await page.getByTestId('datatable-search').fill('pepr')

    if (isInCluster) {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 19)')
    } else {
      expect(await page.getByTestId('table-header-results').textContent()).toBe('(showing 3 of 8)')
    }
  })
})
