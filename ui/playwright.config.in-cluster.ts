// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { defineConfig } from '@playwright/test'

const protocol = 'https'
const host = 'runtime.admin.uds.dev'

export default defineConfig({
  globalSetup: './tests/global-setup',
  timeout: 10 * 1000,
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  retries: process.env.CI ? 2 : 1,
  testMatch: /^(?!.*local-auth|.*reconnect)(.+\.)?(test|spec)\.[jt]s$/,
  use: {
    baseURL: `${protocol}://${host}/`,
    storageState: './tests/state.json',
  },
})
