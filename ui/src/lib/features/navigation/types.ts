// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { CarbonIcon } from 'carbon-icons-svelte'

export interface BaseRoute {
  name: string
  icon?: typeof CarbonIcon
  class?: string
  children?: string[]
}

export interface Route {
  name: string
  path: string
  icon?: typeof CarbonIcon
  class?: string
  children?: RouteChild[]
}

export interface RouteChild {
  name: string
  path: string
}
