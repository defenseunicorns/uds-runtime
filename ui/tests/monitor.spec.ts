import { expect, test } from '@playwright/test'
import * as fs from 'node:fs'

test.describe('Monitor', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/monitor/pepr')
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
