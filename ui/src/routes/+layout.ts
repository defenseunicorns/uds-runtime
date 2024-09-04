// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { createStore } from '$features/k8s/namespaces/store'
import { updateApiAuthEnabled } from '$lib/utils/helpers'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()

  updateApiAuthEnabled()

  return {
    namespaces,
  }
}
