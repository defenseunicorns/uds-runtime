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

export function extractOps(res: Record<string, unknown> | undefined): Record<string, PatchOperation[]> {
  //   const parsed = JSON.parse(res)
  if (!res) {
    return {}
  }

  if (res.patch) {
    const decodedPatch = decodeBase64(res.patch as string)
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

  return {}
}
