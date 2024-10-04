import { get } from 'svelte/store'

import { addToast } from '$features/toast'
import { toast } from '$features/toast/store'

// checkClusterConnection checks the connection to the cluster and
// shows a toast message if the connection is lost or restored.
export function checkClusterConnection() {
  const clusterCheck = new EventSource('/health')
  const disconnectedMsg = 'Cluster health check failed: no connection'

  // handle cluster disconnection and reconnection events
  clusterCheck.onmessage = (msg) => {
    const data = JSON.parse(msg.data)
    const errToast = get(toast).find((t) => t.message === disconnectedMsg)

    // a disconnection occured but has now been resolved
    if (data === 'success' && errToast) {
      toast.update(() => [])
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

    // show a disconnection toast message
    if (data === 'error') {
      addToast({
        type: 'error',
        message: disconnectedMsg,
        noClose: true,
      })
    }
  }

  return clusterCheck
}
