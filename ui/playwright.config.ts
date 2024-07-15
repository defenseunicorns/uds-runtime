import { defineConfig } from '@playwright/test'
import { loadEnv } from 'vite'

const { VITE_PORT_ENV } = loadEnv('dev', process.cwd())
const port = VITE_PORT_ENV ?? '8080'

export default defineConfig({
  testDir: 'tests',
  /* Run tests in files in parallel */
  fullyParallel: true,
  testMatch: /(.+\.)?(test|spec)\.[jt]s/,
  use: {
    baseURL: `http://localhost:${port}/`,
  },
})
