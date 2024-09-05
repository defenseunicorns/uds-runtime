// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors
import { writable } from 'svelte/store'

export const authenticated = writable(false)
export const apiAuthEnabled = writable<null | boolean>(null)
