// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { addToast } from '$features/toast'
import { toast } from '$features/toast/store'

export class ClusterCheck {
  #clusterCheck: EventSource
  #disconnectedMsg = 'Cluster health check failed: no connection'
  #disconnected = false

  constructor() {
    this.#clusterCheck = new EventSource('/health')

    this.#clusterCheck.addEventListener('close', function () {
      this.close()
    })

    this.#clusterCheck.onmessage = (msg) => {
      if (msg.data === 'error') {
        this.#handleDisconnectedEvt()
      } else if (msg.data === 'success' && this.#disconnected) {
        this.#handleReconnectionEvt()
      }
    }
  }

  #handleDisconnectedEvt() {
    addToast({
      type: 'error',
      message: this.#disconnectedMsg,
      noClose: true,
    })
    this.#disconnected = true
  }

  #handleReconnectionEvt() {
    // a disconnection occured but has now been resolved
    // clear the disconnection toast message
    toast.update(() => [])
    this.#disconnected = false

    addToast({
      type: 'success',
      message: 'Cluster connection restored',
      timeoutSecs: 10,
    })

    // Dispatch custom event for reconnection
    // use window instead of svelte createEventDispatcher to trigger event globally
    const event = new CustomEvent('cluster-reconnected', {
      detail: { message: 'Cluster connection restored' },
    })
    window.dispatchEvent(event)
  }

  close() {
    this.#clusterCheck.close()
  }
}
