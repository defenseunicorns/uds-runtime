import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())

// use port 8443 because by default we use TLS when running locally
const port = VITE_PORT_ENV ?? '8080'
const protocol = 'http'
const host = 'localhost'

export default defineConfig({
  timeout: 10 * 1000,
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  retries: process.env.CI ? 2 : 1,
  testMatch: /^(?!.*local-auth|.*reconnect)(.+\.)?(test|spec)\.[jt]s$/,
  use: {
    baseURL: `${protocol}://${host}:${port}/`,
  },
})

export { port }
