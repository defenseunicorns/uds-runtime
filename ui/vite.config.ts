// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vitest/config'

export default defineConfig(({ mode }) => ({
  plugins: [sveltekit()],
  server: {
    proxy: {
      // Proxy all requests starting with /api to the go server
      // noting that we ues https and 8443 because by default we use TLS when running locally
      '/api': {
        target: 'https://runtime-local.uds.dev:8443',
        changeOrigin: true,
      },
      '/health': {
        target: 'https://runtime-local.uds.dev:8443',
        changeOrigin: true,
      },
    },
  },
  test: {
    include: ['src/**/*.test.{js,ts}'],
    environment: 'jsdom',
    globals: true,
    setupFiles: ['src/setupTests.ts'],
  },
  resolve: {
    conditions: mode === 'test' ? ['browser'] : [],
  },
}))
