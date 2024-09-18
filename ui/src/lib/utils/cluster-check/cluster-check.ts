import { addToast } from '$features/toast'
import { getIdByMessage, removeToast } from '$features/toast/store'

// checkClusterConnection checks the connection to the cluster and
// shows a toast message if the connection is lost or restored.
export function checkClusterConnection() {
  const clusterCheck = new EventSource('/health')
  const disconnectedMsg = 'Cluster health check failed: no connection'

  // handle initial connection error
  clusterCheck.onerror = () => {
    addToast({
      type: 'error',
      message: disconnectedMsg,
      timeoutSecs: 1000,
    })
  }

  // handle cluster disconnection and reconnection events
  clusterCheck.onmessage = (msg) => {
    const data = JSON.parse(msg.data) as Record<'success' | 'error' | 'reconnected', string>
    const errToast = getIdByMessage(disconnectedMsg)

    // remove error toast if cluster is reconnected
    if (errToast && data['reconnected']) {
      removeToast(errToast)
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

    // only show error toast once and make timeout really long
    if (!errToast && data['error']) {
      addToast({
        type: 'error',
        message: disconnectedMsg,
        timeoutSecs: 1000,
      })
    }
  }

  return clusterCheck
}
