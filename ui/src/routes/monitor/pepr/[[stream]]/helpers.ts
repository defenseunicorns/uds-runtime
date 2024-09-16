import type { SvelteComponent } from 'svelte'

import type { PatchOperation, PeprDetails, PeprEvent } from '$lib/types'

import DeniedDetails from './(details)/denied-details/DeniedDetails.svelte'
import MutatedDetails from './(details)/mutated-details/MutatedDetails.svelte'

// Utility function to decode base64
function decodeBase64(base64String: string) {
  try {
    return atob(base64String)
  } catch (e) {
    console.error('Failed to decode base64 string:', e)
    return ''
  }
}

export function getDetails(payload: PeprEvent): PeprDetails | undefined {
  if (!payload.res) {
    return undefined
  }

  if (payload.event === 'DENIED') {
    if (payload.res) {
      const status = payload.res.status as Record<string, string>
      const split = status.message.split(' Authorized: ')

      // No "Authorized" or "Found" in the message
      if (split.length !== 2) {
        return { component: DeniedDetails as unknown as SvelteComponent, messages: split }
      }

      const authorized = `Authorized: ${split[1].split(' Found: ')[0]}`
      const found = `Found: ${split[1].split(' Found:')[1]}`
      return { component: DeniedDetails as unknown as SvelteComponent, messages: [authorized, found] }
    }
  }

  if (payload.event === 'MUTATED') {
    if (payload.res.patch) {
      const decodedPatch = decodeBase64(payload.res.patch as string)
      const parsedPatch = JSON.parse(decodedPatch)

      const opMap: { [key: string]: string } = {
        add: 'ADDED',
        remove: 'REMOVED',
        replace: 'REPLACED',
      }

      // Group by operation type
      const groups: { [key: string]: PatchOperation[] } = {}
      for (const op of parsedPatch) {
        if (!groups[opMap[op.op]]) {
          groups[opMap[op.op]] = []
        }
        groups[opMap[op.op]].push(op)
      }

      return { component: MutatedDetails as unknown as SvelteComponent, operations: groups }
    }
  }

  return undefined
}
