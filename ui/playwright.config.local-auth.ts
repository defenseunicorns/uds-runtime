// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())
const port = VITE_PORT_ENV ?? '8443'

export default defineConfig({
  timeout: 60 * 1000,
  testDir: 'tests',
  fullyParallel: false,
  retries: process.env.CI ? 2 : 1,
  testMatch: /^(?!.*in-cluster|.*reconnect)(.+\.)?(test|spec)\.[jt]s$/,
  use: {
    baseURL: `https://runtime-local.uds.dev:${port}/`,
  },
})

export { port }
