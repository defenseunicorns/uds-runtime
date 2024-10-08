import { addToast } from '$features/toast'
import { toast } from '$features/toast/store'

// checkClusterConnection checks the connection to the cluster and
// shows a toast message if the connection is lost or restored.

export class ClusterCheck {
  #clusterCheck: EventSource
  #disconnectedMsg = 'Cluster health check failed: no connection'
  #disconnected = false

  constructor() {
    this.#clusterCheck = new EventSource('/health')
    this.#clusterCheck.onmessage = (msg) => {
      this.#handleCloseEvt(msg.data)
      this.#handleDisconnectedEvt(msg.data)
      this.#handleReconnectionEvt(msg.data)
    }
  }

  #handleDisconnectedEvt(data: string) {
    console.log(data, 'error')
    if (data === 'error') {
      addToast({
        type: 'error',
        message: this.#disconnectedMsg,
        noClose: true,
      })
      this.#disconnected = true
    }
  }

  #handleReconnectionEvt(data: string) {
    // a disconnection occured but has now been resolved
    if (data === 'success' && this.#disconnected) {
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
  }

  #handleCloseEvt(data: string) {
    if (data.includes('close')) {
      this.#clusterCheck.close()
    }
  }

  close() {
    this.#clusterCheck.close()
  }
}
