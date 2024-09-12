/** @type {import('prettier').Config} */

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
