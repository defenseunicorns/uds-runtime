// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

export const stringToSnakeCase = (name: string) => name.split(' ').join('-').toLocaleLowerCase()
export const routeToTitle = (route: string): string =>
  (route.split('/').pop() || '').replace(/-/g, ' ').replace(/\b\w/g, (char) => char.toUpperCase())
