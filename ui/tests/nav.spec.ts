import { expect, test } from '@playwright/test'

// experimenting with different approach to testing routes
// could collapse all these into 1 test
test('home page has expected h1', async ({ page }) => {
  await page.goto('/')

  await expect(page.locator('h1')).toBeVisible()
})

test('navigation to monitor pepr', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Monitor')
  await page.click('text=Pepr')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Monitor  Pepr')
})

test('navigation to monitor events', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Monitor')
  await page.click('text=Events')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Monitor  Events')
})

test('navigation to resources pods', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Resources')
  await page.click('text=Pods')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Resources  Pods')
})

test('navigation to resources namespaces', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Resources')
  await page.click('text=Namespaces')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Resources  Namespaces')
})

// Add more tests for other routes here
test('navigation to resources deployments', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Resources')
  await page.click('text=Deployments')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Resources  Deployments')
})

test('navigation to resources services', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Resources')
  await page.click('text=Services')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Resources  Services')
})

test('navigation to resources packages', async ({ page }) => {
  await page.goto('/')
  await page.click('text=Resources')
  await page.click('text=Packages')
  const breadcrumb = page.getByLabel('Breadcrumb')
  await expect(breadcrumb).toHaveText('Resources  Packages')
})

test('navigation to settings', async ({ page }) => {
  await page.goto('/')
  await page.getByLabel('cog outline').click()
  await expect(page.locator('h1')).toHaveText('Settings')
})

test('navigation to preferences', async ({ page }) => {
  await page.goto('/')
  await page.getByLabel('adjustments horizontal outline').click() 
  await expect(page.locator('h1')).toHaveText('Preferences')
})


test('navigation to help', async ({ page }) => {
  await page.goto('/')
  await page.getByLabel('question circle outline').click() 
  await expect(page.locator('h1')).toHaveText('Help')
})