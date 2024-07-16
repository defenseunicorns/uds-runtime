import { expect, test, type Locator } from '@playwright/test'

test.describe('Navigation', async () => {
  let breadcrumb: Locator

  test.beforeEach(async ({ page }) => {
    breadcrumb = page.getByLabel('Breadcrumb')
    await page.goto('/')
  })

  test('Overview page', async ({ page }) => {
    await page.getByRole('link', { name: 'Overview' }).click()
    await expect(breadcrumb.locator('li', { hasText: 'Overview' })).toBeVisible()
  })

  test.describe('navigates to Monitor', async () => {
    test('Pepr page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Pepr' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Monitor' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Pepr' })).toBeVisible()
    })

    test('Events page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Events' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Monitor' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Events' })).toBeVisible()
    })
  })

  test.describe('navigates to Workloads', async () => {
    test('Pods page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Pods' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Pods' })).toBeVisible()
    })

    test('Deployments page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Deployments' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Deployments' })).toBeVisible()
    })

    test('DaemonSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'DaemonSets' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'DaemonSets' })).toBeVisible()
    })

    test('StatefulSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'StatefulSets' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'StatefulSets' })).toBeVisible()
    })

    test('Jobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: /^Jobs$/ }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Jobs' })).toBeVisible()
    })

    test('CronJobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'CronJobs' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Workloads' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'CronJobs' })).toBeVisible()
    })
  })

  test.describe('navigates to Config', async () => {
    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Packages' }).click()
      await expect(breadcrumb.locator('li', { hasText: 'Config' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'UDS Packages' })).toBeVisible()
    })

    test('UDS Exemptions page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Exemptions' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Config' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'UDS Exemptions' })).toBeVisible()
    })

    test('ConfigMaps page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'ConfigMaps' }).click()

      await expect(breadcrumb.locator('li', { hasText: /^Config$/ })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'ConfigMaps' })).toBeVisible()
    })

    test('Secrets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'Secrets' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Config' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Secrets' })).toBeVisible()
    })
  })

  test.describe('navigates to Cluster Ops', async () => {
    test('Mutating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Mutating Webhooks' }).click()
      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Mutating Webhooks' })).toBeVisible()
    })

    test('Validating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Validating Webhooks' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Validating Webhooks' })).toBeVisible()
    })

    test('HPA page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'HPA' }).click()

      await expect(breadcrumb.locator('li', { hasText: /^Cluster Ops$/ })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'HPA' })).toBeVisible()
    })

    test('Pod Disruption Budgets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Pod Disruption Budgets' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Pod Disruption Budgets' })).toBeVisible()
    })

    test('Resource Quotas page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Resource Quotas' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Resource Quotas' })).toBeVisible()
    })

    test('Limit Ranges page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Limit Ranges' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Limit Ranges' })).toBeVisible()
    })

    test('Priority Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Priority Classes' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Priority Classes' })).toBeVisible()
    })

    test('Runtime Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Runtime Classes' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Cluster Ops' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Runtime Classes' })).toBeVisible()
    })
  })

  test.describe('navigates to Network', async () => {
    test('Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: /^Services$/ }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Network' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Services' })).toBeVisible()
    })

    test('Virtual Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Virtual Services' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Network' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Virtual Services' })).toBeVisible()
    })

    test('Network Policies page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Network Policies' }).click()

      await expect(breadcrumb.locator('li', { hasText: /^Network$/ })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Network Policies' })).toBeVisible()
    })

    test('Endpoints page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Endpoints' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Network' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Endpoints' })).toBeVisible()
    })
  })

  test.describe('navigates to Storage', async () => {
    test('Persistent Volumes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volumes' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Storage' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Persistent Volumes' })).toBeVisible()
    })

    test('Persistent Volume Claims page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volume Claims' }).click()

      await expect(breadcrumb.locator('li', { hasText: 'Storage' })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Persistent Volume Claims' })).toBeVisible()
    })

    test('Storage Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Storage Classes' }).click()

      await expect(breadcrumb.locator('li', { hasText: /^Storage$/ })).toBeVisible()
      await expect(breadcrumb.locator('li', { hasText: 'Storage Classes' })).toBeVisible()
    })
  })

  test('Namespaces page', async ({ page }) => {
    await page.getByRole('link', { name: 'Namespaces' }).click()
    await expect(breadcrumb.locator('li', { hasText: 'Namespaces' })).toBeVisible()
  })

  test('navigates to Docs page', async ({ page }) => {
    await page.getByRole('link', { name: 'Docs' }).click()

    await expect(page.locator('h1')).toHaveText('Docs')
  })

  test('navigates to Preferences page', async ({ page }) => {
    await page.getByTestId('global-sidenav-preferences').click()

    await expect(page.locator('h1')).toHaveText('Preferences')
  })

  test('navigates to Settings page', async ({ page }) => {
    await page.getByTestId('global-sidenav-settings').click()

    await expect(page.locator('h1')).toHaveText('Settings')
  })

  test('navigates to Help page', async ({ page }) => {
    await page.getByTestId('global-sidenav-help').click()

    await expect(page.locator('h1')).toHaveText('Help')
  })
})
