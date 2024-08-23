import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())
const port = VITE_PORT_ENV ?? '8080'

export default defineConfig({
  // webServer: {
  //   command: 'VITE_API_AUTH=true ../build/uds-runtime',
  //   url: `http://localhost:${port}`,
  //   reuseExistingServer: !process.env.CI,
  // },
  timeout: 60 * 1000,
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: false,
  retries: process.env.CI ? 2 : 1,
  testMatch: /(.+\.)?(test|spec)\.[jt]s/,
  use: {
    baseURL: `http://localhost:${port}/`,
  },
})

export { port }
