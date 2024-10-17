// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())

// use port 8443 because by default we use TLS when running locally
const port = VITE_PORT_ENV ?? '8443'
const protocol = 'https'
const host = 'runtime-local.uds.dev'

export default defineConfig({
  webServer: {
    command: '../build/uds-runtime',
    url: `${protocol}://${host}:${port}`,
    reuseExistingServer: !process.env.CI,
    env: { LOCAL_AUTH_ENABLED: 'false' },
  },
  timeout: 10 * 1000,
  testDir: 'tests',
  fullyParallel: false,
  retries: process.env.CI ? 2 : 1,
  testMatch: 'reconnect.spec.ts',
  use: {
    baseURL: `https://runtime-local.uds.dev:${port}/`,
  },
})

export { port }
