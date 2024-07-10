import { expect, test } from '@playwright/test'

test.describe('Navigation', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/')
  })

  test('takes me to Pepr page when I click on Monitor followed by Pepr link', async ({ page }) => {
    await page.getByRole('button', { name: 'Monitor' }).click()
    await page.getByRole('link', { name: 'Pepr' }).click()

    await expect(page.getByRole('link', { name: 'Monitor' })).toBeVisible()
    await expect(page.getByLabel('Breadcrumb').getByRole('link', { name: 'Pepr' })).toBeVisible()
  })
})
