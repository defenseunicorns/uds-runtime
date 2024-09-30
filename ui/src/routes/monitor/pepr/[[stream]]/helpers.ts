import type { SvelteComponent } from 'svelte'
import type { Writable } from 'svelte/store'

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

function getDetails(payload: PeprEvent): PeprDetails | undefined {
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

export function filterEvents(events: PeprEvent[], searchTerm: string): PeprEvent[] {
  // filter events by the search term if one exists
  if (!searchTerm) return events
  const searchValue = searchTerm.toLowerCase()
  return events.filter(
    (item) =>
      item._name.toLowerCase().includes(searchValue) ||
      item.event.toLowerCase().includes(searchValue) ||
      item.header.toLowerCase().includes(searchValue) ||
      item.msg.toLowerCase().includes(searchValue),
  )
}

export function sortEvents(events: PeprEvent[], sortKey: string, isAscending: boolean): PeprEvent[] {
  const sortDirection = isAscending ? 1 : -1 // sort events in ascending order by default
  // sort events based on the sort key
  return events.sort((a, b) => {
    if (sortKey === 'timestamp') {
      const aTime = a.ts ? new Date(a.ts).getTime() : a.epoch
      const bTime = b.ts ? new Date(b.ts).getTime() : b.epoch
      return (aTime - bTime) * sortDirection * -1 // latest events on top?
    } else if (sortKey === 'count') {
      const aValue = Number(a[sortKey as keyof typeof a]) || 0
      const bValue = Number(b[sortKey as keyof typeof b]) || 0
      return (aValue - bValue) * sortDirection
    } else {
      const aValue = String(a[sortKey as keyof typeof a] || '').toLowerCase()
      const bValue = String(b[sortKey as keyof typeof b] || '').toLowerCase()
      return aValue.localeCompare(bValue) * sortDirection
    }
  })
}

export const exportPeprStream = (rows: PeprEvent[]) => {
  const data = rows.map((item) => ({
    event: item.event,
    resource: item._name,
    count: item.count,
    timestamp: item.ts,
  }))

  const blob = new Blob([JSON.stringify(data, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `pepr-stream-${new Date().toISOString()}.json`

  try {
    a.click()
  } finally {
    setTimeout(() => {
      URL.revokeObjectURL(url)
    }, 100) // debounce to ensure download has started
  }
}

export function handlePeprMessage(e: MessageEvent, peprStreamStore: Writable<PeprEvent[]>, peprStream: PeprEvent[]) {
  try {
    const payload: PeprEvent = JSON.parse(e.data)
    // The event type is the first word in the header
    payload.event = payload.header.split(' ')[0]
    payload.details = getDetails(payload)

    // handle "repeated"-type payloads
    if (payload.repeated) {
      // these events don't have a _name attribute so we just update the pepr stream row directly
      const idx = peprStream.findIndex((item) => item.header === payload.header)
      if (idx !== -1) {
        peprStreamStore.update((collection) => {
          collection[idx].count = payload.repeated!
          return collection
        })
      }
      return // "repeated"-type payload handled, no need to continue
    }

    // check existing rows for duplicates
    const dupIdx = peprStream.findIndex((item) => item.header === payload.header && item.ts === payload.ts)
    if (dupIdx !== -1) {
      // remove duplicate from the stream and update with the latest payload
      peprStreamStore.update((collection) => {
        collection.splice(dupIdx, 1)
        return [payload, ...collection]
      })
      return
    }
    // payload isn't a dup, add to stream
    peprStreamStore.update((collection) => [payload, ...collection])
  } catch (error) {
    console.error('Error updating peprStream:', error)
  }
}
