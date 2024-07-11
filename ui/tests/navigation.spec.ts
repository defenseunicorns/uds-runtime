import { expect, test, type Locator } from '@playwright/test'

test.describe('Navigation', async () => {
  let breadcrumb: Locator

  test.beforeEach(async ({ page }) => {
    breadcrumb = page.getByLabel('Breadcrumb')

    await page.goto('/')
  })

  test.describe('navigates to Monitor', async () => {
    test('Pepr page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Pepr' }).click()

      await expect(page.getByRole('link', { name: 'Monitor' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Pepr' })).toBeVisible()
    })

    test('Events page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Events' }).click()

      await expect(page.getByRole('link', { name: 'Monitor' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Events' })).toBeVisible()
    })
  })

  test.describe('navigates to Resources', async () => {
    test('Pods page', async ({ page }) => {
      await page.getByRole('button', { name: 'Resources' }).click()
      await page.getByRole('link', { name: 'Pods' }).click()

      await expect(page.getByRole('link', { name: 'Resources' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Pods' })).toBeVisible()
    })

    test('Namespaces page', async ({ page }) => {
      await page.getByRole('button', { name: 'Resources' }).click()
      await page.getByRole('link', { name: 'Namespaces' }).click()

      await expect(page.getByRole('link', { name: 'Resources' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Namespaces' })).toBeVisible()
    })

    test('Deployments page', async ({ page }) => {
      await page.getByRole('button', { name: 'Resources' }).click()
      await page.getByRole('link', { name: 'Deployments' }).click()

      await expect(page.getByRole('link', { name: 'Resources' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Deployments' })).toBeVisible()
    })

    test('Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Resources' }).click()
      await page.getByRole('link', { name: 'Services' }).click()

      await expect(page.getByRole('link', { name: 'Resources' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Services' })).toBeVisible()
    })

    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Resources' }).click()
      await page.getByRole('link', { name: 'Packages' }).click()

      await expect(page.getByRole('link', { name: 'Resources' })).toBeVisible()
      await expect(breadcrumb.getByRole('link', { name: 'Packages' })).toBeVisible()
    })
  })

  test('navigates to Docs page', async ({ page }) => {
    await page.getByRole('link', { name: 'Docs' }).click()

    await expect(page.locator('h1')).toHaveText('Docs')
  })

  test('navigates to Settings page', async ({ page }) => {
    await page.getByLabel('cog outline').click()

    await expect(page.locator('h1')).toHaveText('Settings')
  })

  test('navigates to Preferences page', async ({ page }) => {
    await page.getByLabel('adjustments horizontal outline').click()

    await expect(page.locator('h1')).toHaveText('Preferences')
  })

  test('navigates to Help page', async ({ page }) => {
    await page.getByLabel('question circle outline').click()

    await expect(page.locator('h1')).toHaveText('Help')
  })
})
