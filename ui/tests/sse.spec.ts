// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import k8s from '@kubernetes/client-node'
import { expect, test } from '@playwright/test'

// Annotate entire file as serial.
test.describe.configure({ mode: 'serial' })

async function deletePod(namespace: string, podName: string, force: boolean = true) {
  try {
    const kc = new k8s.KubeConfig()
    kc.loadFromDefault() // Load the kubeconfig file from default location

    const k8sApi = kc.makeApiClient(k8s.CoreV1Api)
    await k8sApi.deleteNamespacedPod({ name: podName, namespace: namespace, gracePeriodSeconds: force ? 0 : undefined })
    console.log(`Pod ${podName} deleted successfully`)
  } catch (err) {
    console.error(`Failed to delete pod ${podName}:`, err)
  }
}

test.describe('SSE and reactivity', async () => {
  test('Pods are updated', async ({ page }) => {
    await page.goto('/workloads/pods')
    const allPodRows = page.locator('.table .tr').filter({ hasText: /^podinfo-/ })
    let originalPodName = await allPodRows.first().textContent()

    originalPodName = originalPodName ? originalPodName.trim() : ''

    // get pod name
    expect(originalPodName).not.toBeNull()

    // delete pod and wait for it to disappear
    await deletePod('podinfo', originalPodName)
    await expect(page.locator('.table .tr').filter({ hasText: /originalPodName/ })).toBeHidden()

    // get new pod
    const newPodName = await allPodRows.first().textContent()

    expect(newPodName).not.toBeNull()
    expect(newPodName).not.toEqual(originalPodName)
  })
})
