// Utility function to decode base64
function decodeBase64(base64String: string) {
  try {
    return atob(base64String)
  } catch (e) {
    console.error('Failed to decode base64 string:', e)
    return ''
  }
}

export function formatPatch(data: string): string {
  const parsed = JSON.parse(data)
  if (parsed.res) {
    const decodedPatch = decodeBase64(parsed.res.patch)
    const parsedPatch = JSON.parse(decodedPatch)

    // Format the JSON patch
    const formattedPatch: string[] = []

    // Group by operation type
    const groups: { [key: string]: any[] } = {}
    for (const op of parsedPatch) {
      if (!groups[op.op]) {
        groups[op.op] = []
      }
      groups[op.op].push(op)
    }

    const opMap: { [key: string]: string } = {
      add: 'ADDED',
      remove: 'REMOVED',
      replace: 'REPLACED',
    }

    // Write the patch for each operation type
    for (const name in groups) {
      const ops = groups[name]
      let format = '\n%s             %v=%v'

      if (name === 'remove') {
        format = '\n%s              %v'
      }

      // Write the subheader for the operation type
      formattedPatch.push(`\n${opMap[name]}:`)

      // Write the patch for each operation group
      for (const op of ops) {
        const key = op.path
        const val = JSON.stringify(op.value)
        formattedPatch.push(format.replace('%s', ' ').replace('%v', key).replace('%v', val))
      }
    }
    console.log(formattedPatch.join(''))

    return formattedPatch.join('')
  }

  return ''
}
