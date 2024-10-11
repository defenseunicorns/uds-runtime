// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())
const port = VITE_PORT_ENV ?? '8443'

export default defineConfig({
  timeout: 60 * 1000,
  testDir: 'tests',
  fullyParallel: false,
  retries: 0,
  testMatch: /(.+\.)?(test|spec)\.[jt]s/,
  use: {
    baseURL: `https://runtime-local.uds.dev:${port}/`,
  },
})

export { port }
