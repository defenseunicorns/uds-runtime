import { chromium, expect, type FullConfig } from '@playwright/test'

// Before running the in-cluster e2e tests, log into the app using
// the doug keycloak user and save the storage state to a file.
async function globalSetup(config: FullConfig) {
  const { baseURL, storageState } = config.projects[0].use
  const browser = await chromium.launch()
  const page = await browser.newPage()
  await page.goto(baseURL!)
  await page.getByLabel('username').fill('doug')
  await page.getByLabel('password').fill('unicorn123!@#UN')
  await page.getByText('Log In').click()

  // ensure we are logged in and the home page's user menu is visible
  await expect(page.getByText('Doug Unicorn')).toBeVisible()
  console.log('Logged in successfully') // useful to know when globalSetup has finished in CI
  await page.context().storageState({ path: storageState as string })
  await browser.close()
}

export default globalSetup
