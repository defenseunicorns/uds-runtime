import { expect, test } from '@playwright/test'
import * as fs from 'node:fs'

test.describe('Navigation', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/')
  })

  test('Overview page', async ({ page }) => {
    await page.getByRole('link', { name: 'Overview' }).click()
    await expect(page.getByTestId('breadcrumb-item-overview')).toBeVisible()
  })

  test.describe('navigates to Monitor', async () => {
    test('Pepr page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Pepr' }).click()

      await expect(page.getByTestId('breadcrumb-item-monitor')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-pepr')).toBeVisible()
    })

    test('Events page', async ({ page }) => {
      await page.getByRole('button', { name: 'Monitor' }).click()
      await page.getByRole('link', { name: 'Events' }).click()

      await expect(page.getByTestId('breadcrumb-item-monitor')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-events')).toBeVisible()
    })
  })

  test.describe('navigates to Workloads', async () => {
    test('Pods page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Pods' }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-pods')).toBeVisible()
    })

    test('Deployments page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'Deployments' }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-deployments')).toBeVisible()
    })

    test('DaemonSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'DaemonSets' }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-daemonsets')).toBeVisible()
    })

    test('StatefulSets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'StatefulSets' }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-statefulsets')).toBeVisible()
    })

    test('Jobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: /^Jobs$/ }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-jobs')).toBeVisible()
    })

    test('CronJobs page', async ({ page }) => {
      await page.getByRole('button', { name: 'Workloads' }).click()
      await page.getByRole('link', { name: 'CronJobs' }).click()

      await expect(page.getByTestId('breadcrumb-item-workloads')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-cronjobs')).toBeVisible()
    })
  })

  test.describe('navigates to Config', async () => {
    test('Packages page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Packages' }).click()

      await expect(page.getByTestId('breadcrumb-item-config')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-uds-packages')).toBeVisible()
    })

    test('UDS Exemptions page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'UDS Exemptions' }).click()

      await expect(page.getByTestId('breadcrumb-item-config')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-uds-exemptions')).toBeVisible()
    })

    test('ConfigMaps page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'ConfigMaps' }).click()

      await expect(page.getByTestId('breadcrumb-item-config')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-configmaps')).toBeVisible()
    })

    test('Secrets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Config' }).click()
      await page.getByRole('link', { name: 'Secrets' }).click()

      await expect(page.getByTestId('breadcrumb-item-config')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-secrets')).toBeVisible()
    })
  })

  test.describe('navigates to Cluster Ops', async () => {
    test('Mutating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Mutating Webhooks' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-mutating-webhooks')).toBeVisible()
    })

    test('Validating Webhooks page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Validating Webhooks' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-validating-webhooks')).toBeVisible()
    })

    test('HPA page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'HPA' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-hpa')).toBeVisible()
    })

    test('Pod Disruption Budgets page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Pod Disruption Budgets' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-pod-disruption-budgets')).toBeVisible()
    })

    test('Resource Quotas page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Resource Quotas' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-resource-quotas')).toBeVisible()
    })

    test('Limit Ranges page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Limit Ranges' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-limit-ranges')).toBeVisible()
    })

    test('Priority Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Priority Classes' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-priority-classes')).toBeVisible()
    })

    test('Runtime Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Cluster Ops' }).click()
      await page.getByRole('link', { name: 'Runtime Classes' }).click()

      await expect(page.getByTestId('breadcrumb-item-cluster-ops')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-runtime-classes')).toBeVisible()
    })
  })

  test.describe('navigates to Network', async () => {
    test('Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: /^Services$/ }).click()

      await expect(page.getByTestId('breadcrumb-item-network')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-services')).toBeVisible()
    })

    test('Virtual Services page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Virtual Services' }).click()

      await expect(page.getByTestId('breadcrumb-item-network')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-virtual-services')).toBeVisible()
    })

    test('Network Policies page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Network Policies' }).click()

      await expect(page.getByTestId('breadcrumb-item-network')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-network-policies')).toBeVisible()
    })

    test('Endpoints page', async ({ page }) => {
      await page.getByRole('button', { name: 'Network' }).click()
      await page.getByRole('link', { name: 'Endpoints' }).click()

      await expect(page.getByTestId('breadcrumb-item-network')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-endpoints')).toBeVisible()
    })
  })

  test.describe('navigates to Storage', async () => {
    test('Persistent Volumes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volumes' }).click()

      await expect(page.getByTestId('breadcrumb-item-storage')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-persistent-volumes')).toBeVisible()
    })

    test('Persistent Volume Claims page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Persistent Volume Claims' }).click()

      await expect(page.getByTestId('breadcrumb-item-storage')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-persistent-volume-claims')).toBeVisible()
    })

    test('Storage Classes page', async ({ page }) => {
      await page.getByRole('button', { name: 'Storage' }).click()
      await page.getByRole('link', { name: 'Storage Classes' }).click()

      await expect(page.getByTestId('breadcrumb-item-storage')).toBeVisible()
      await expect(page.getByTestId('breadcrumb-item-storage-classes')).toBeVisible()
    })
  })

  test('Namespaces page', async ({ page }) => {
    await page.getByRole('link', { name: 'Namespaces' }).click()

    await expect(page.getByTestId('breadcrumb-item-namespaces')).toBeVisible()
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
