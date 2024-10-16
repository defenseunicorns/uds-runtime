import { execSync } from 'child_process'

import { expect, test } from '@playwright/test'
import { K8s, kind } from 'kubernetes-fluent-client'

// Utility function to run shell commands
function execCommand(command: string) {
  execSync(command, { stdio: 'inherit' })
}

async function createPod() {
  await K8s(kind.Pod).Apply({
    metadata: {
      name: 'new-pod',
      namespace: 'default',
    },
    spec: {
      containers: [
        {
          name: 'my-container',
          image: 'nginx',
        },
      ],
    },
  })
}

test.describe('Cluster Reconnection and Pod Creation Test', () => {
  test('should handle cluster disconnection, reconnection, and pod creation', async ({ page }) => {
    test.setTimeout(120000)
    await page.goto('/workloads/pods')

    // Stop the cluster
    execCommand('k3d cluster stop runtime')

    await expect(page.getByText('Cluster health check failed: no connection')).toBeVisible()

    // Start the cluster again
    execCommand('k3d cluster start runtime')

    await expect(page.getByText('Cluster connection restored')).toBeVisible()

    // Use KFC to create a new pod
    await createPod()

    await expect(page.getByText('new-pod')).toBeVisible({ timeout: 10000 })
  })
})
