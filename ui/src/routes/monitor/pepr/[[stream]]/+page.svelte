<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onDestroy } from 'svelte'
  import { derived, writable, type Unsubscriber } from 'svelte/store'

  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import { type PeprEvent } from '$lib/types'
  import { ChevronUp, Export, Information, Search } from 'carbon-icons-svelte'

  import './page.postcss'

  import { getDetails } from './helpers'

  let loaded = false
  let streamFilter = ''
  let eventSource: EventSource | null = null
  let unsubscribePage: Unsubscriber

  const streamSelectOptions = [
    { value: '', label: 'All Pepr Events' },
    { value: 'policies', label: 'All UDS Policies' },
    { value: 'allowed', label: 'UDS Policies: Allowed' },
    { value: 'denied', label: 'UDS Policies: Denied' },
    { value: 'mutated', label: 'UDS Policies: Mutated' },
    { value: 'operator', label: 'UDS Operator' },
    { value: 'failed', label: 'Errors and Denials' },
  ]

  const peprStream = writable<PeprEvent[]>([])
  export let columns = [
    { name: 'event', style: 'w-2/12' },
    { name: 'resource', style: 'w-3/12' },
    { name: 'details', style: 'w-1/12' },
    { name: 'count', style: 'w-1/12' },
    { name: 'timestamp', style: 'w-5/12' },
  ]

  // Initialize the stores
  let search = writable<string>('')
  let sortBy = writable<string>('timestamp')
  let sortAsc = writable<boolean>(true)

  // check for filtering
  let isFiltering = false
  $: {
    isFiltering = !!$search
  }

  function filterEvents(events: PeprEvent[], searchTerm: string): PeprEvent[] {
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

  function sortEvents(events: PeprEvent[], sortKey: string, isAscending: boolean): PeprEvent[] {
    const sortDirection = isAscending ? 1 : -1 // sort events in ascending order by default
    // sort events based on the sort key
    return events.sort((a, b) => {
      if (sortKey === 'timestamp') {
        const aTime = a.ts ? new Date(a.ts).getTime() : a.epoch
        const bTime = b.ts ? new Date(b.ts).getTime() : b.epoch
        return (aTime - bTime) * sortDirection
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

  export const rows = derived([peprStream, search, sortBy, sortAsc], () => {
    const filteredEvents = filterEvents($peprStream, $search)
    return sortEvents(filteredEvents, $sortBy, $sortAsc)
  })

  onDestroy(() => {
    eventSource?.close()
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

    const path: string = `/api/v1/monitor/pepr/${streamFilter}`
    eventSource = new EventSource(path)

    // Set the loaded flag when the connection is established
    eventSource.onopen = () => {
      loaded = true
    }

    eventSource.onmessage = (e) => {
      try {
        const payload: PeprEvent = JSON.parse(e.data)
        // The event type is the first word in the header
        payload.event = payload.header.split(' ')[0]
        payload.details = getDetails(payload)

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
    const data = $rows.map((item) => ({
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

  function handleStreamChange(event: Event) {
    const target = event.target as HTMLSelectElement
    goto(`/monitor/pepr/${target.value}`)
  }
</script>

<section class="table-section">
  <div class="table-container">
    <div class="table-content">
      <div class="table-header">
        <span class="dark:text-white" data-testid="table-header">{'Pepr Events'}</span>
        {#if isFiltering}
          <span class="dark:text-gray-500 pl-2" data-testid="table-header-results">
            (showing {$rows.length} of {$peprStream.length})
          </span>
        {:else}
          <span class="dark:text-gray-500 pl-2" data-testid="table-header-results">({$peprStream.length})</span>
        {/if}
        <div class="relative group">
          <Information class="ml-2 w-4 h-4 text-gray-400" />
          <div class="tooltip tooltip-right min-w-72">
            <div class="whitespace-normal">
              {'These are UDS Operator logs scraped from Pepr running in the cluster'}
            </div>
          </div>
        </div>
      </div>
      <div class="table-filter-section">
        <div class="relative lg:w-96">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <Search class="h-5 w-5 text-gray-400" />
          </div>
          <input
            type="text"
            name="input-search"
            autocomplete="off"
            class="focus:ring-primary-500 focus:border-primary-500 dark:focus:ring-primary-500 dark:focus:border-primary-500 block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 pl-9 text-gray-900 sm:text-sm dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400"
            placeholder="Search"
            data-testid="datatable-search"
            bind:value={$search}
          />
        </div>
        <div>
          <select
            id="stream"
            bind:value={streamFilter}
            on:change={handleStreamChange}
            data-testid="table-filter-stream-select"
          >
            {#each streamSelectOptions as s}
              <option value={s.value} data-testid={`stream-select-${s.label}`}>
                {s.label}
              </option>
            {/each}
          </select>
        </div>
        <div class="flex-grow"></div>
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
              {#each columns as col}
                <th class={col.style}>
                  <button
                    on:click={() => {
                      $sortBy = col.name === 'resource' ? '_name' : col.name
                      $sortAsc = !$sortAsc
                    }}
                  >
                    {col.name.replaceAll('_', ' ')}
                    <ChevronUp
                      class="sort
                      {$sortAsc ? 'rotate-180' : ''}
                      {$sortBy === col.name ? 'opacity-100' : 'opacity-0'}"
                    />
                  </button>
                </th>
              {/each}
            </tr>
          </thead>
          {#if loaded}
            <tbody>
              {#if $rows.length === 0}
                <tr>
                  <td colspan="4" class="text-center">No matching entries found</td>
                </tr>
              {:else}
                {#each $rows as item}
                  <tr>
                    <td>
                      <span class="pepr-event {item.event}">{item.event}</span>
                    </td>
                    <td data-testid={`pepr-resource-${item._name}`}>{item._name}</td>
                    <td class="flex flex-row items-center">
                      {#if item.details}
                        <svelte:component this={item.details.component} details={item.details} />
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
                  <td class={columns[0].style}>
                    <div class="h-6 rounded bg-gray-600 w-20"></div>
                  </td>
                  <td class={columns[1].style}>
                    <div class="h-6 bg-gray-500 rounded {w}"></div>
                  </td>
                  <td class={columns[2].style}>
                    <div class="h-6 bg-gray-600 rounded w-20"></div>
                  </td>
                  <td class={columns[3].style}>
                    <div class="h-6 bg-gray-600 rounded w-8"></div>
                  </td>
                  <td class={columns[4].style}>
                    <div class="h-6 bg-gray-600 rounded w-1/2"></div>
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
