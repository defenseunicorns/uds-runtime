import { createStore } from '$lib/stores/resources/namespaces'

// The /resources routes share a common namespace store
export const load = async () => {
  const namespaces = createStore()
  namespaces.start()

  return {
    namespaces,
  }
}
