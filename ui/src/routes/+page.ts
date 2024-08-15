// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// Temporary redirect to /overview until home page is implemented
import { redirect } from '@sveltejs/kit'

export function load() {
  throw redirect(307, '/overview')
}
