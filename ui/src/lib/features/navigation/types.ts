// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { CarbonIcon } from 'carbon-icons-svelte'

export interface BaseRoute {
  name: string
  icon?: typeof CarbonIcon
  class?: string
  children?: string[]
  path?: string
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

// UserData is the shape of the user data returned from /user
export interface UserData {
  name: string
  preferredUsername: string
  group: string
  inClusterAuth: boolean
}
