import type { KubernetesObject } from '@kubernetes/client-node'

export async function loadResourceDetail(url: string, uid: string) {
  const [path] = url.split('?')
  const response = await fetch(`${path}/${uid}`)
  const data = await response.json()
  return data.Object as KubernetesObject
}
