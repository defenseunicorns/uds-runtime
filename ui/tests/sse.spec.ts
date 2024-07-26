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
    const originalPodName = await page.getByRole('cell', { name: 'minio' }).first().textContent()

    // get pod name
    expect(originalPodName).not.toBeNull()

    // delete pod and wait for it to disappear
    await deletePod('uds-dev-stack', originalPodName ?? '')
    await expect(page.getByRole('cell', { name: originalPodName ?? '' })).toBeHidden()

    // get new pod
    const newPodName = await page.getByRole('cell', { name: 'minio' }).first().textContent()

    expect(newPodName).not.toBeNull()
    expect(newPodName).not.toEqual(originalPodName)
  })

  test('PVCs with pod are updated', async ({ page }) => {
    await page.goto('/storage/persistent-volume-claims')
    const originalPVCPodName = await page.getByText('minio-').first().textContent()

    // get pod attached to pvc's name
    expect(originalPVCPodName).not.toBeNull()

    // delete pod attached to PVC and wait for it to disappear
    await deletePod('uds-dev-stack', originalPVCPodName ?? '')
    await expect(page.getByRole('cell', { name: originalPVCPodName ?? '' })).toBeHidden()

    // get new pod attached to PVC
    const newPVCPodName = await page.getByRole('cell', { name: 'minio-' }).first().textContent()

    expect(newPVCPodName).not.toBeNull()
    expect(newPVCPodName).not.toEqual(originalPVCPodName)
  })
})
