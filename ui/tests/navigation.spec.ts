// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('Navigation', async () => {
  // podinfo is deployed into the test cluster, we will use this to check if various pages render correctly
  const query = 'podinfo'

  test.beforeEach(async ({ page }) => {
    await page.goto('/')
  })

  test('Overview page', async ({ page }) => {
    await page.getByRole('link', { name: 'Overview' }).click()

    const query = '1' // number of running nodes
    const element = page.locator(`//dd[normalize-space(text())="${query}"]`)
    await expect(element).toHaveText('1') // ensure exact match
  })

  test.describe('navigates to Applications', async () => {
    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Applications' }).click()
      await page.getByRole('link', { name: 'Packages' }).click()

      const query = 'podinfo-test' // package name
      const element = page.locator(`.emphasize:has-text("${query}")`)
      await expect(element).toBeVisible()
    })
  })

  test.describe('navigates to Monitor', async () => {
    test('Pepr page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Pepr' }).click()

      const query = 'uds-policy-exemptions/podinfo2' // package name
      const element = page.locator(`td:has-text("${query}")`).first()
      await expect(element).toBeVisible()
    })

    test('Events page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Events' }).click()
    })
  })

  test.describe('navigates to Workloads', async () => {
    test('Pods page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Pods' }).click()

      const element = page.locator(`.emphasize:has-text("${query}")`)
      await expect(element).toBeVisible()
    })

    test('Deployments page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Deployments' }).click()

      const element = page.locator(`.emphasize:has-text("${query}")`)
      await expect(element).toBeVisible()
    })

    test('DaemonSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'DaemonSets' }).click()
    })

    test('StatefulSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'StatefulSets' }).click()
    })

    test('Jobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: /^Jobs$/ }).click()
    })

    test('CronJobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'CronJobs' }).click()
    })
  })

  test.describe('navigates to Config', async () => {
    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Packages' }).click()

      const element = page.locator(`.emphasize:has-text("${query}")`)
      await expect(element).toBeVisible()
    })

    test('UDS Exemptions page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Exemptions' }).click()

      let query = 'podinfo2' // exemption name
      let element = page.locator(`.emphasize:has-text("${query}")`)
      await expect(element).toBeVisible()

      query = '- RequireNonRootUser' // exemption policy name
      element = page.locator(`td:has-text("${query}")`)
      await expect(element).toBeVisible()
    })

    test('ConfigMaps page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'ConfigMaps' }).click()
    })

    test('Secrets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'Secrets' }).click()
    })
  })

  test.describe('navigates to Cluster Ops', async () => {
    test('Mutating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Mutating Webhooks' }).click()
    })

    test('Validating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Validating Webhooks' }).click()
    })

    test('HPA page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'HPA' }).click()
    })

    test('Pod Disruption Budgets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Pod Disruption Budgets' }).click()
    })

    test('Resource Quotas page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Resource Quotas' }).click()
    })

    test('Limit Ranges page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Limit Ranges' }).click()
    })

    test('Priority Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Priority Classes' }).click()
    })

    test('Runtime Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Runtime Classes' }).click()
    })
  })

  test.describe('navigates to Network', async () => {
    test('Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: /^Services$/ }).click()
    })

    test('Virtual Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Virtual Services' }).click()
    })

    test('Network Policies page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Network Policies' }).click()

      const query = 'allow-podinfo-egress-dns-lookup-via-coredns' // network policy name
      const element = page.locator(`.emphasize:has-text("${query}")`).first()
      await expect(element).toBeVisible()
    })

    test('Endpoints page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Endpoints' }).click()
    })
  })

  test.describe('navigates to Storage', async () => {
    test('Persistent Volumes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volumes' }).click()
    })

    test('Persistent Volume Claims page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volume Claims' }).click()
    })

    test('Storage Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Storage Classes' }).click()
    })
  })

  test('navigates to Namespaces page', async ({ page }) => {
    await page.getByRole('link', { name: 'Namespaces' }).click()

    const element = page.locator(`.emphasize:has-text("${query}")`).first()
    await expect(element).toBeVisible()
  })

  test('navigates to Nodes page', async ({ page }) => {
    await page.getByRole('link', { name: 'Nodes' }).click()

    const query = 'k3d-runtime-server-0'
    const element = page.locator(`.emphasize:has-text("${query}")`).first()
    await expect(element).toBeVisible()
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
