<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { onDestroy } from 'svelte'
  import { derived, writable, type Unsubscriber } from 'svelte/store'

  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import { type PeprEvent } from '$lib/types'
  import { ChevronUp, Export, Information, Search } from 'carbon-icons-svelte'

  import './page.postcss'

  import { exportPeprStream, filterEvents, handlePeprMessage, sortEvents } from './helpers'

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

  let columns = [
    { name: 'event', style: 'w-2/12' },
    { name: 'resource', style: 'w-3/12' },
    { name: 'details', style: 'w-1/12' },
    { name: 'count', style: 'w-1/12' },
    { name: 'timestamp', style: 'w-5/12' },
  ]

  // Initialize the stores
  const peprStream = writable<PeprEvent[]>([])
  let search = writable<string>('')
  let sortBy = writable<string>('timestamp')
  let sortAsc = writable<boolean>(true)

  // check for filtering
  let isFiltering = false
  $: {
    isFiltering = !!$search
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
      handlePeprMessage(e, peprStream, $peprStream)
    }

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error)
    }
  })

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
          <button name="Export" type="button" on:click={() => exportPeprStream($rows)}>
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
                  <td colspan="5" class="text-center">No matching entries found</td>
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
