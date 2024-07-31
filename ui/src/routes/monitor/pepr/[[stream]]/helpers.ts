import type { PeprEvent } from '$lib/types'

export type PatchOperation = {
  op: string
  path: string
  value: string
}

// Utility function to decode base64
function decodeBase64(base64String: string) {
  try {
    return atob(base64String)
  } catch (e) {
    console.error('Failed to decode base64 string:', e)
    return ''
  }
}

export function getDetails(payload: PeprEvent): Record<string, any[]> | undefined {
  if (!payload.res) {
    return undefined
  }

  if (payload.event === 'DENIED') {
    if (payload.res) {
      const msgs = []
      const status = payload.res.status as Record<string, string>

      const split = status.message.split(' Authorized: ')

      if (split.length !== 2) {
        return { messages: split }
      }
      const authorized = `Authorized: ${split[1].split(' Found: ')[0]}`
      const found = `Found: ${split[1].split(' Found:')[1]}`
      msgs.push(authorized)
      msgs.push(found)
      return { messages: msgs }
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
      const operations: { [key: string]: PatchOperation[] } = {}
      for (const op of parsedPatch) {
        if (!operations[opMap[op.op]]) {
          operations[opMap[op.op]] = []
        }
        operations[opMap[op.op]].push(op)
      }

      return operations
    }
  }

  return undefined
}
