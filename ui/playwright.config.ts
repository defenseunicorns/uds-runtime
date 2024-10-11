// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())

// use port 8443 because by default we use TLS when running locally
const port = VITE_PORT_ENV ?? '8443'
const protocol = 'https'
const host = 'runtime-local.uds.dev'

export default defineConfig({
  webServer: {
    command: 'LOCAL_AUTH_ENABLED=false ../build/uds-runtime',
    url: `${protocol}://${host}:${port}`,
    reuseExistingServer: !process.env.CI,
  },
  timeout: 10 * 1000,
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  retries: process.env.CI ? 2 : 1,
  testMatch: /^(?!.*local-auth)(.+\.)?(test|spec)\.[jt]s$/,
  use: {
    baseURL: `${protocol}://${host}:${port}/`,
  },
})

export { port }
