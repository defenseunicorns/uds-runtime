// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// Test file prefixed with "z-" to ensure it runs last since it stops the server and restarts it
import { expect, test } from '@playwright/test'
import { ChildProcess, exec } from 'child_process'
import * as net from 'net'

let serverProcess: ChildProcess
const serverLogs: string[] = []
let extractedToken: string | null = null

function isPortInUse(port: number): Promise<boolean> {
  return new Promise((resolve) => {
    const server = net
      .createServer()
      .once('error', () => resolve(true))
      .once('listening', () => {
        server.close()
        resolve(false)
      })
      .listen(port)
  })
}

test.beforeAll(async () => {
  // await killProcessOnPort(port.toString())
  // Check what is running on port 8080
  try {
    const inUse = await isPortInUse(8080)
    if (inUse) {
      console.log('Port 8080 is in use')
    } else {
      console.log('Port 8080 is not in use')
    }
  } catch (error) {
    console.error(`Error checking port 8080: ${error}`)
  }

  // Start the server
  await new Promise<void>((resolve, reject) => {
    console.time('Server Start Time') // Start the timer

    serverProcess = exec('VITE_API_AUTH=true ../build/uds-runtime', (error) => {
      if (error) {
        console.error(`Error starting server: ${error}`)
        console.timeEnd('Server Start Time')
        reject(error)
        return
      }
    })

    // Capture stdout
    if (serverProcess && serverProcess.stdout) {
      serverProcess.stdout.on('data', (data) => {
        const log = data.toString()
        serverLogs.push(`stdout: ${log}`)
        resolve() // Resolve the promise after pushing to serverLogs
        extractToken(log)
      })
    }
  })

  // Wait for the server to be ready
  await new Promise((resolve) => setTimeout(resolve, 10000)) // Adjust the timeout as needed
})

test.afterAll(async () => {
  // Stop the server
  if (serverProcess) {
    serverProcess.kill()
  }
})

// async function killProcessOnPort(port: string) {
//   const list = await find('port', port)
//   list.forEach((proc) => {
//     try {
//       process.kill(proc.pid)
//     } catch (e) {
//       console.error(`Failed to kill process ${proc.pid}: ${e.message}`)
//     }
//   })
// }

function extractToken(log: string) {
  const match = log.match(/auth\?token=([^&\s]+)/)
  if (match) {
    extractedToken = match[1]
  }
}

test.describe.serial('Authentication Tests', () => {
  test('authenticated access', async ({ page }) => {
    await page.goto(`/auth?token=${extractedToken}`)
    await page.waitForSelector('role=link[name="Overview"]', { state: 'visible', timeout: 10000 })
    await page.getByRole('link', { name: 'Overview' }).click()
    const nodeCountEl = page.getByTestId(`node-count`)
    await expect(nodeCountEl).toHaveText('1')
  })

  test('pod view access', async ({ page }) => {
    await page.goto(`/auth?token=${extractedToken}`)
    await page.getByRole('button', { name: 'Workloads' }).click()
    await page.getByRole('link', { name: 'Pods' }).click()
    const element = page.locator(`.emphasize:has-text("podinfo")`)
    await expect(element).toBeVisible()
  })

  test('unauthenticated access', async ({ page }) => {
    await page.goto(`/auth?token=insecure`)
    const unauthenticated = page.getByText('Could not authenticate')
    await expect(unauthenticated).toBeVisible()
  })
})
