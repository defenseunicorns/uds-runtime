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
    test.setTimeout(150000)
    await page.goto('/workloads/pods')

    // Step 2: Stop the cluster
    execCommand('k3d cluster stop runtime')
    let errToastFound: boolean | undefined = false
    let toastExists = await page.waitForFunction(
      () => document.querySelector('div')?.innerText.includes('Cluster health check failed: no connection'),
      { timeout: 10000 },
    )

    errToastFound = await toastExists.jsonValue()
    expect(errToastFound).toBe(true)

    // Step 4: Start the cluster again
    if (errToastFound) {
      execCommand('k3d cluster start runtime')
    }

    // Step 5: Wait for the error toast to disappear and the success toast to appear
    let successToastFound: boolean | undefined = false
    toastExists = await page.waitForFunction(
      () => document.querySelector('div')?.innerText.includes('Cluster connection restored'),
      { timeout: 10000 }, // Wait up to 15 seconds for the toast to appear
    )
    successToastFound = await toastExists.jsonValue()
    expect(successToastFound).toBe(true)
    expect(page.getByText('Cluster health check failed: no connection')).not.toBeVisible()
    expect(page.getByText('Cluster connection restored')).toBeVisible()

    // Step 6: Use KFC to create a new pod
    await createPod()

    // Step 7: Assert that the new pod is visible in the view without reloading the page
    let newPodFound: boolean | undefined = false
    const podExists = await page.waitForFunction(() => document.querySelector('div')?.innerText.includes('new-pod'), {
      timeout: 10000,
    })

    newPodFound = await podExists.jsonValue()
    expect(newPodFound).toBe(true)
    expect(page.getByText('new-pod')).toBeVisible()
  })
})
