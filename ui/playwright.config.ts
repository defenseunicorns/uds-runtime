import type { PlaywrightTestConfig } from '@playwright/test'

const config: PlaywrightTestConfig = {
  webServer: {
    command: 'npm run build && npm run preview',
    port: 5173,
  },
  testDir: 'tests',
  testMatch: /(.+\.)?(spec)\.[jt]s/,
  use: {
    baseURL: 'http://localhost:5173/',
  },
}

export default config
