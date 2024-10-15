// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

module.exports = {
  useTabs: false,
  tabWidth: 2,
  singleQuote: true,
  trailingComma: 'all',
  printWidth: 120,
  semi: false,
  htmlWhitespaceSensitivity: 'ignore',
  importOrder: [
    '^(svelte/(.*)$)|^(svelte$)',
    '',
    '<THIRD_PARTY_MODULES>',
    '',
    '$app/(.*)$',
    '$features/(.*)$',
    '$components/(.*)$',
    '$lib/(.*)$',
    '',
    '^[./]',
  ],
  plugins: ['prettier-plugin-svelte', '@ianvs/prettier-plugin-sort-imports'],
  overrides: [
    {
      files: '*.svelte',
      options: {
        parser: 'svelte',
      },
    },
  ],
}
