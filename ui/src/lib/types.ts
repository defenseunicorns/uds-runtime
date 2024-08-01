// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { SvelteComponent } from 'svelte'

export type PatchOperation = {
  op: string
  path: string
  value: string
}

export type PeprDetails = {
  component: SvelteComponent
  messages?: string[]
  operations?: { [key: string]: PatchOperation[] }
}

export type PeprEvent = {
  _name: string
  count: number
  event: string
  header: string
  repeated?: number
  ts?: string
  epoch: number
  msg: string
  res?: Record<string, unknown>
  details?: PeprDetails | undefined
}
