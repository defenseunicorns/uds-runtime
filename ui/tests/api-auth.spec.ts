// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'
import { ChildProcess, exec } from 'child_process'

let serverProcess: ChildProcess
const serverLogs: string[] = []
let extractedToken: string | null = null

test.beforeAll(async () => {
  // Start the server
  return new Promise<void>((resolve, reject) => {
    serverProcess = exec('../build/uds-runtime', (error) => {
      if (error) {
        console.error(`Error starting server: ${error}`)
      }
    })

    if (serverProcess && serverProcess.stderr) {
      serverProcess.stderr.on('data', (data) => {
        const log = data.toString()
        console.error(`stderr: ${log}`)
        serverLogs.push(`stderr: ${log}`)
        extractToken(log)
        resolve()
      })
    }

    // Handle process exit
    serverProcess.on('exit', (code) => {
      console.log(`Server process exited with code ${code}`)
      if (code !== 0) {
        reject(new Error(`Server process exited with code ${code}`))
      }
    })
  })
})

test.afterAll(async () => {
  // Stop the server
  if (serverProcess) {
    serverProcess.kill()
  }
})

function extractToken(log: string) {
  const match = log.match(/auth\?token=([^&\s]+)/)
  if (match) {
    extractedToken = match[1]
    // ANSI escape codes are being appended to the token
    extractedToken = stripAnsiCodes(extractedToken)
    console.log(`Extracted token: ${extractedToken}`)
  }
}

function stripAnsiCodes(str: string): string {
  // This regex matches all ANSI escape codes
  // eslint-disable-next-line no-control-regex
  const ansiRegex = /[\x1B\x9B][[()#;?]*(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[0-9A-ORZcf-nqry=><]/g
  return str.replace(ansiRegex, '')
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
