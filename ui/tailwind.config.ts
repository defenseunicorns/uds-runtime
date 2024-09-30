// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import flowbitePlugin from 'flowbite/plugin'
import type { Config } from 'tailwindcss'

const TAILWIND_SIZE_OPTIONS = [16, 20, 24, 28, 32, 36, 40, 44, 48] as const

export default {
  content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite/**/*.js'],

  theme: {
    fontFamily: {
      body: [
        'Montserrat',
        'ui-sans-serif',
        'system-ui',
        '-apple-system',
        'system-ui',
        'Segoe UI',
        'Roboto',
        'Helvetica Neue',
        'Arial',
        'Noto Sans',
        'sans-serif',
        'Apple Color Emoji',
        'Segoe UI Emoji',
        'Segoe UI Symbol',
        'Noto Color Emoji',
      ],
      sans: [
        'Montserrat',
        'ui-sans-serif',
        'system-ui',
        '-apple-system',
        'system-ui',
        'Segoe UI',
        'Roboto',
        'Helvetica Neue',
        'Arial',
        'Noto Sans',
        'sans-serif',
        'Apple Color Emoji',
        'Segoe UI Emoji',
        'Segoe UI Symbol',
        'Noto Color Emoji',
      ],
    },
  },

  variants: {},

  plugins: [flowbitePlugin],
  safelist: TAILWIND_SIZE_OPTIONS.map((item: number) => `h-${item}`),
} as Config
