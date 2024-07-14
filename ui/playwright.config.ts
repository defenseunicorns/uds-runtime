// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { defineConfig } from '@playwright/test'

export default defineConfig({
  webServer: {
    command: 'npm run build && npm run preview',
    url: 'http://localhost:4173/',
  },
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  testMatch: /(.+\.)?(spec|test)\.[jt]s/,
  use: {
    baseURL: 'http://localhost:4173/',
  },
})
