<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { Export } from 'carbon-icons-svelte'
  import { onDestroy } from 'svelte'
  import { writable, type Unsubscriber } from 'svelte/store'
  import Detail from './Detail.svelte'

  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import { type PeprEvent } from '$lib/types'
  import './page.postcss'
  import { extractOps } from './helpers'

  let loaded = false
  let streamFilter = ''
  let eventSource: EventSource | null = null
  let unsubscribePage: Unsubscriber

  const peprStream = writable<PeprEvent[]>([])

  onDestroy(() => {
    unsubscribePage()
  })

  unsubscribePage = page.subscribe(({ route, params }) => {
    // Reset the page when the route changes
    eventSource?.close()
    loaded = false

    // This will trigger when leaving the page too, so skip if not the right route
    if (route.id !== '/monitor/pepr/[[stream]]') {
      return
    }

    peprStream.set([])
    streamFilter = params.stream || ''

    eventSource = new EventSource(`/api/v1/monitor/pepr/${streamFilter}`)

    // Set the loaded flag when the connection is established
    eventSource.onopen = () => {
      loaded = true
    }

    eventSource.onmessage = (e) => {
      try {
        const payload: PeprEvent = JSON.parse(e.data)

        // The event type is the first word in the header
        payload.event = payload.header.split(' ')[0]

        if (payload.event === 'MUTATED') {
          payload.details = extractOps(payload.res)
        }

        // If this is a repeated event, update the count
        if (payload.repeated) {
          // Find the first item in the peprStream that matches the header
          peprStream.update((collection) => {
            const idx = collection.findIndex((item) => item.header === payload.header)
            if (idx !== -1) {
              collection[idx].count = payload.repeated!
              collection[idx].ts = payload.ts
            }
            return collection
          })
        } else {
          // Otherwise, add the new event to the peprStream
          peprStream.update((collection) => [payload, ...collection])
        }
      } catch (error) {
        console.error('Error updating peprStream:', error)
      }
    }

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error)
    }
  })

  const exportPeprStream = () => {
    const data = $peprStream.map((item) => ({
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

  const widths = ['w-1/6', 'w-1/3', 'w-1/4', 'w-2/5', 'w-1/2', 'w-1/5', 'w-1/3', 'w-1/4']
  const skeletonRows = widths.sort(() => Math.random() - 0.5)
</script>

<section class="table-section">
  <div class="table-container">
    <div class="table-content">
      <div class="table-filter-section">
        <div class="grid w-full grid-cols-1 md:grid-cols-4 md:gap-4 lg:w-2/3">
          <div class="w-full">
            <select
              id="stream"
              bind:value={streamFilter}
              on:change={(val) => {
                goto(`/monitor/pepr/${val.target.value}`)
              }}
            >
              <option value="">All Data</option>
              <hr />
              <option value="policies">UDS Policies</option>
              <option value="allowed">UDS Policies: Allowed</option>
              <option value="denied">UDS Policies: Denied</option>
              <option value="mutated">UDS Policies: Mutated</option>
              <hr />
              <option value="operator">UDS Operator</option>
              <option value="failed">Errors and Denials</option>
            </select>
          </div>
        </div>
        <div
          class="flex flex-shrink-0 flex-col space-y-3 md:flex-row md:items-center md:space-x-3 md:space-y-0 lg:justify-end"
        >
          <button name="Export" type="button" on:click={exportPeprStream}>
            <Export class="mr-2" />
            Export
          </button>
        </div>
      </div>
      <div class="table-scroll-container">
        <table>
          <thead>
            <tr>
              <th>Event</th>
              <th>Resource</th>
              <th>Details</th>
              <th>Count</th>
              <th>Timestamp</th>
            </tr>
          </thead>
          {#if loaded}
            <tbody>
              {#if $peprStream.length === 0}
                <tr>
                  <td colspan="4" class="text-center">No matching entries found</td>
                </tr>
              {:else}
                {#each $peprStream as item}
                  <tr>
                    <td>
                      <span class="pepr-event {item.event}">{item.event}</span>
                    </td>
                    <td>{item._name}</td>
                    <td class="flex flex-row items-center">
                      {#if item.details}
                        <Detail details={item.details} />
                      {:else}
                        -
                      {/if}
                    </td>
                    <td>{item.count || 1}</td>
                    <td>{item.ts}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          {:else}
            <tbody class="animate-pulse">
              {#each skeletonRows as w}
                <tr class="border-b border-gray-700">
                  <td class="py-2 px-4 w-36">
                    <div class="h-6 rounded px-2 py-0.5 bg-gray-600"></div>
                  </td>
                  <td class="py-2 px-4">
                    <div class="h-6 bg-gray-500 rounded {w}"></div>
                  </td>
                  <td class="py-2 px-4 w-24">
                    <div class="h-6 bg-gray-600 rounded w-8"></div>
                  </td>
                  <td class="py-2 px-4 w-64">
                    <div class="h-6 bg-gray-600 rounded w-full"></div>
                  </td>
                </tr>
              {/each}
            </tbody>
          {/if}
        </table>
      </div>
    </div>
  </div>
</section>
