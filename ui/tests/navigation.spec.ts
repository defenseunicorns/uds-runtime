// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, test } from '@playwright/test'

test.describe('Navigation', async () => {
  test.beforeEach(async ({ page }) => {
    await page.setViewportSize({
      width: 3024,
      height: 1964,
    })

    await page.goto('/')
  })

  test('Overview page', async ({ page }) => {
    await page.getByRole('link', { name: 'Overview' }).click()

    const nodeCountEl = page.getByTestId('resource-count-nodes')
    await expect(nodeCountEl).toHaveText('1')

    const card = page.getByTestId('card-container')

    await expect(card.getByText('Pods running in cluster')).toBeVisible()
    await expect(card.getByText('Nodes running in cluster')).toBeVisible()
    await expect(card.getByText('CPU Usage')).toBeVisible()
    await expect(card.getByText('Memory Usage')).toBeVisible()

    // Check for Events Widget
    await expect(page.getByText('Event Logs')).toBeVisible()
    await expect(page.getByText('VIEW EVENTS')).toBeVisible()

    // Check for no unavailable tags when metrics server is available
    const count = await page.getByTestId('unavailable-tag').count()
    const overviewPodCount = await page.getByTestId('resource-count-pods').textContent()
    // indicates running in cluster e2e tests and metrics server in not available
    if (overviewPodCount === '20') {
      expect(count === 3).toBe(true)
      // indicates running other e2e tests where metrics server is available
    } else {
      expect(count === 0).toBe(true)
    }
  })

  test('Ensure Overview page and pod page show same number of pods', async ({ page }) => {
    // get pod count from overview page
    await page.getByRole('link', { name: 'Overview' }).click()
    const overviewPodCount = await page.getByTestId('resource-count-pods').textContent()

    // navigate to pods page and get pod count
    await page.goto('/workloads/pods')
    await page.waitForSelector('.emphasize:has-text("podinfo")') // wait for pods to render
    let podCount = await page.getByTestId('table-header-results').textContent()
    expect(podCount).not.toBeNull()

    // remove parentheses
    podCount = podCount!.replace(/\(|\)/g, '')

    await expect(overviewPodCount).toEqual(podCount)
  })

  test('Navigate to page when click on Pod and Node Cards', async ({ page }) => {
    const podsCard = page.getByTestId('card-container').filter({ hasText: 'Pods running in Cluster' })
    await podsCard.click()
    expect(await page.getByTestId('table-header').textContent()).toEqual('Pods')

    await page.goto('/')

    const nodesCard = page.getByTestId('card-container').filter({ hasText: 'Nodes running in Cluster' })
    await nodesCard.click()
    expect(await page.getByTestId('table-header').textContent()).toEqual('Nodes')
  })

  test.describe('navigates to Applications', async () => {
    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Applications' }).click()
      await page.getByRole('link', { name: 'Packages' }).click()

      await expect(page.getByTestId('podinfo-test-testid-1')).toHaveText('podinfo-test')
    })
  })

  test.describe('navigates to Monitor', async () => {
    test('Pepr page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Pepr' }).click()

      const query = 'uds-policy-exemptions/podinfo2' // package name
      await expect(page.getByTestId(`pepr-resource-${query}`)).toHaveText(query)
    })

    test('Events page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Events' }).click()
    })
  })

  test.describe('navigates to Workloads', async () => {
    test('Pods page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: /^Pods$/ }).click()

      const element = page.locator(`.emphasize:has-text("podinfo")`).first()
      await expect(element).toBeVisible()
    })

    test('Deployments page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Deployments' }).click()

      await expect(page.getByTestId('podinfo-testid-1')).toHaveText('podinfo')
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

      await expect(page.getByTestId('podinfo-testid-1')).toHaveText('podinfo')
    })

    test('UDS Exemptions page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Exemptions' }).click()

      await expect(page.getByTestId('podinfo2-testid-1')).toHaveText('podinfo2')

      const policy = 'RequireNonRootUser'
      await expect(page.getByTestId(`${policy}-list-item-test-id`)).toHaveText(`- ${policy}`)
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

      await page.waitForSelector('[data-testid="allow-podinfo-egress-dns-lookup-via-coredns-testid-1"]')

      await expect(page.getByTestId('allow-podinfo-egress-dns-lookup-via-coredns-testid-1')).toHaveText(
        'allow-podinfo-egress-dns-lookup-via-coredns',
      )
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
      await expect(page.getByText('minio-').first()).toBeVisible() // ensure pods have rendered
    })

    test('Storage Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Storage Classes' }).click()
    })
  })

  test('navigates to Namespaces page', async ({ page }) => {
    await page.getByRole('link', { name: 'Namespaces' }).click()

    await expect(page.getByTestId('podinfo-testid-1')).toHaveText('podinfo')
  })

  test('navigates to Nodes page', async ({ page }) => {
    await page.getByRole('link', { name: /^Nodes$/ }).click()

    await expect(page.getByTestId('control-plane, master-testid-3')).toHaveText('control-plane, master')
  })

  test('navigates to Preferences page', async ({ page }) => {
    await page.getByTestId('global-sidenav-preferences').click()

    await expect(page.getByText('Preferences')).toBeVisible()
  })

  test('navigates to Settings page', async ({ page }) => {
    await page.getByTestId('global-sidenav-settings').click()

    await expect(page.getByText('Settings')).toBeVisible()
  })

  test('navigates to Help page', async ({ page }) => {
    await page.getByTestId('global-sidenav-help').click()

    await expect(page.getByText('Help', { exact: true })).toBeVisible()
  })
})
