import { createStore } from '$stores/resources/namespaces'

export const ssr = false

// Provide shared access to the cluster namespace store
export const load = async () => {
  const namespaces = createStore()
  namespaces.start()

  return {
    namespaces,
  }
}
