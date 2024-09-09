// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'
import * as fs from 'node:fs'

test.describe('Monitor', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/monitor/pepr')

    // wait for data to load
    await page.waitForSelector('.pepr-event.ALLOWED')
  })

  test('searching with "All Pepr Events selected', async ({ page }) => {
    await page.getByTestId('datatable-search').fill('podinfo/podinfo')

    const results = await page.getByTestId('table-header-results').textContent()
    expect(results).toBeTruthy()

    const numbers = results!.match(/\d+/g)!.map(Number)
    expect(numbers![0]).toBeLessThan(numbers![1])
  })

  test('searching while using filter dropdown', async ({ page }) => {
    await page.getByTestId('datatable-search').fill('podinfo/podinfo')
    await page.getByTestId('datatable-filter-dropdown').click()
    await page.getByText('UDS Policies: Allowed').click()

    // wait for data to load
    await page.waitForSelector('.pepr-event.ALLOWED')

    const results = await page.getByTestId('table-header-results').textContent()
    expect(results).toBeTruthy()

    const numbers = results!.match(/\d+/g)!.map(Number)
    expect(numbers![0]).toBeLessThan(numbers![1])
  })

  test('sorting by count', async ({ page }) => {
    // get first row
    let rows = await page.$$('table tr:first-child')
    let firstRow = await rows[1].textContent() // index 0 is the table header
    expect(firstRow).toBeTruthy()
    let match = firstRow!.match(/\s(\d+)\s/)
    expect(match).toBeTruthy()
    const originalCount = parseInt(match![1])

    // click "count" to sort
    await page.getByText('count').click()

    // get first row after sorting
    rows = await page.$$('table tr:first-child')
    firstRow = await rows[1].textContent() // index 0 is the table header
    expect(firstRow).toBeTruthy()
    match = firstRow!.match(/\s(\d+)\s/)
    expect(match).toBeTruthy()
    const newCount = parseInt(match![1])

    // sorted 'count' value should be greater than the original
    expect(newCount).toBeGreaterThan(originalCount)
  })

  test('Exports logs', async ({ page }) => {
    // wait for pepr data to load
    await page.waitForSelector('.pepr-event.ALLOWED')

    // download logs
    const [download] = await Promise.all([
      page.waitForEvent('download'),
      await page.getByRole('button', { name: 'Export' }).click(),
    ])
    const path = await download.path()
    await download.saveAs(path)

    // Read and inspect the contents of the downloaded file
    fs.readFile(path, 'utf8', (err, data) => {
      if (err) {
        console.error('Error reading the file:', err)
        return
      }
      const fileContents = JSON.parse(data)
      expect(fileContents.length).toBeGreaterThan(0)
    })
  })
})
