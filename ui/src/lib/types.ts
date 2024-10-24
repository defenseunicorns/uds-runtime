// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { SvelteComponent } from 'svelte'

import type { UDSPackageStatus } from '$features/k8s/types'

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

export type CoreServiceType = {
  metadata: {
    name: string
    namespace: string
  }
  status: {
    phase: UDSPackageStatus
  }
}
