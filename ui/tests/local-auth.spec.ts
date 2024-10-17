// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { ChildProcess, exec } from 'child_process'

import { expect, test } from '@playwright/test'

let serverProcess: ChildProcess
const serverLogs: string[] = []
let extractedToken: string | null = null

test.beforeAll(async () => {
  // Start the server here (not in playwright config) so we can grab the token from the logs
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
  const match = log.match(/\?token=([^&\s]+)/)
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

    const nodeCountEl = page.getByTestId('resource-count-nodes')
    expect(nodeCountEl).toBe('1')
  })

  test('data is visible on load, refresh, and new tab', async ({ page, context }) => {
    await page.goto(`/auth?token=${extractedToken}`)
    await page.getByRole('button', { name: 'Workloads' }).click()
    await page.getByRole('link', { name: 'Pods' }).click()
    const element = page.locator(`.emphasize:has-text("podinfo")`).first()

    expect(element).toBeVisible()

    // Check details view
    await page
      .locator('.table .tr')
      .filter({ hasText: /^podinfo-/ })
      .click()

    let drawerEl = page.getByTestId('drawer')

    expect(drawerEl).toBeVisible()
    expect(drawerEl.getByText('Created')).toBeVisible()
    expect(drawerEl.getByText('Name', { exact: true })).toBeVisible()
    expect(drawerEl.getByText('Annotations')).toBeVisible()
    expect(drawerEl.getByText('podinfo', { exact: true })).toBeVisible()

    // test data still visible after reload (drawer should still be open)
    await page.reload()
    const reloadedElement = page.locator(`.emphasize:has-text("podinfo")`).first()
    expect(reloadedElement).toBeVisible()

    drawerEl = page.getByTestId('drawer')
    expect(drawerEl).toBeVisible()
    expect(drawerEl.getByText('Created')).toBeVisible()

    // Test opening in a new tab
    const deploymentsLink = page.getByText('Deployments')
    const [newPage] = await Promise.all([
      context.waitForEvent('page'),
      deploymentsLink.click({ button: 'middle' }), // Middle-click to open in new tab
    ])
    await newPage.waitForLoadState()
    const newPageElement = newPage.locator(`.emphasize:has-text("podinfo")`).first()
    expect(newPageElement).toBeVisible()

    await newPage.close()
  })

  test('unauthenticated access', async ({ page }) => {
    await page.goto(`/auth?token=insecure`)
    const unauthenticated = page.getByText('Could not authenticate')

    expect(unauthenticated).toBeVisible()
  })
})
