// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

export const stringToSnakeCase = (name: string) => name.split(' ').join('-').toLocaleLowerCase()

export async function fetchConfig(): Promise<any> {
  const response = await fetch('/config')
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }
  const envVars = await response.json()
  return envVars
}
