// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import * as fs from 'node:fs'

import { expect, test } from '@playwright/test'

// Run monitor tests in serial to ensure we capture the un-cached response time in the first test
test.describe.serial('Monitor', async () => {
  // used to store the initial load time of the page before the cache is used
  let initialLoadTime: number

  test.beforeEach(async ({ page }) => {
    const startTime = new Date().getTime()
    await page.goto('/monitor/pepr')

    // wait for data to load
    await page.waitForSelector('.pepr-event.ALLOWED')

    const endTime = new Date().getTime()

    // initial load time is the time to load the data when the cache isn't used (ie. the first visit to the page)
    initialLoadTime = endTime - startTime
  })

  test('Pepr cache', async ({ page }) => {
    // reload the page
    await page.reload()

    // wait for data to load
    const startTime = new Date().getTime()
    await page.waitForSelector('.pepr-event.ALLOWED')
    const endTime = new Date().getTime()
    const cachedLoadTime = endTime - startTime

    // adding debug logs to help diagnose test failures
    console.debug('Initial load time:', initialLoadTime)
    console.debug('Cached load time:', cachedLoadTime)

    expect(cachedLoadTime).toBeLessThan(initialLoadTime)
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
    await page.getByTestId('table-filter-stream-select').click()
    await page.selectOption('select#stream', 'UDS Policies: Allowed')

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

    // sorted 'count' value should be greater than or equal to the original
    expect(newCount).toBeGreaterThanOrEqual(originalCount)
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
