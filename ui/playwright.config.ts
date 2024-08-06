import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())
const port = VITE_PORT_ENV ?? '8080'

export default defineConfig({
  webServer: {
    command: '../build/uds-runtime',
    url: `http://localhost:${port}`,
    reuseExistingServer: !process.env.CI,
  },
  timeout: 10 * 1000,
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  testMatch: /(.+\.)?(test|spec)\.[jt]s/,
  use: {
    baseURL: `http://localhost:${port}/`,
  },
})
