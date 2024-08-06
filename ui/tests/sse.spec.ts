// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import k8s from '@kubernetes/client-node'
import { expect, test } from '@playwright/test'

async function deletePod(namespace: string, podName: string) {
  try {
    const kc = new k8s.KubeConfig()
    kc.loadFromDefault() // Load the kubeconfig file from default location

    const k8sApi = kc.makeApiClient(k8s.CoreV1Api)
    await k8sApi.deleteNamespacedPod({ name: podName, namespace: namespace })
    console.log(`Pod ${podName} deleted successfully`)
  } catch (err) {
    console.error(`Failed to delete pod ${podName}:`, err)
  }
}

test.describe('SSE and reactivity', async () => {
  test('Pods are updated', async ({ page }) => {
    await page.goto('/workloads/pods')
    const originalPodName = await page.getByRole('cell', { name: 'podinfo' }).first().textContent()

    // get pod name
    expect(originalPodName).not.toBeNull()

    // delete pod and wait for it to disappear
    await deletePod('podinfo', originalPodName ?? '')
    await expect(page.getByRole('cell', { name: originalPodName ?? '' })).toBeHidden()

    // get new pod
    const newPodName = await page.getByRole('cell', { name: 'podinfo' }).first().textContent()

    expect(newPodName).not.toBeNull()
    expect(newPodName).not.toEqual(originalPodName)
  })
})
