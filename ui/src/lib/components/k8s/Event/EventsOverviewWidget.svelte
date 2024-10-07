<script lang="ts">
  import { onMount } from 'svelte'

  import type { KubernetesObject } from '@kubernetes/client-node'
  import { goto } from '$app/navigation'
  import { getColorAndStatus } from '$features/k8s/helpers'
  import { type ResourceStoreInterface } from '$features/k8s/types'
  import { type Columns } from '$lib/features/k8s/events/store'
  import { ChevronRight, Information, Search } from 'carbon-icons-svelte'

  let columns: Columns = [
    ['namespace', 'w-2/12'],
    ['age', 'w-1/12'],
    ['type', 'w-2/12'],
    ['reason', 'w-2/12'],
    ['object_kind', 'w-1/12'],
    ['object_name', 'w-3/12'],
    ['count', 'w-1/12'],
  ]

  export let title: string
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  export let createStore: () => ResourceStoreInterface<KubernetesObject, any>
  export let description: string

  let rows = createStore()
  $: ({ namespace, search } = rows)

  // check for filtering
  let isFiltering = false
  $: {
    isFiltering = !!($search || $namespace)
  }

  onMount(() => {
    let stop = rows.start()

    return () => stop()
  })

  const calculateTypeClass = (key: string, rowData: string): string => {
    let color: string = ''

    if (key === 'type' && rowData === 'Normal') {
      color = getColorAndStatus('Logs', 'Normal')
    }

    if (key === 'type' && rowData === 'Warning') {
      color = getColorAndStatus('Logs', 'Warning')
    }

    return color
  }
</script>

<div class="events">
  <!-- Header which has Title, Dropdown and Search-->
  <div class="events__header">
    <div class="w-3/12 min-[1400px]:w-6/12 flex">
      <div class="flex h-6 items-center space-x-1">
        <h2 class="events__header-title">{title}</h2>

        <div class="relative group">
          <Information class="ml-1 w-4 h-4 dark:text-gray-400 text-blue-500" />

          <div class="tooltip tooltip-right min-w-72">
            <div class="whitespace-normal">{description}</div>
          </div>
        </div>
      </div>
    </div>

    <div class="w-9/12 min-[1400px]:w-6/12 flex justify-end">
      <div class="w-2/3">
        <form class="flex items-center">
          <label for="simple-search" class="sr-only">Search</label>
          <div class="relative w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <Search size={16} class="text-gray-500 dark:text-gray-400" />
            </div>

            <input
              type="text"
              id="simple-search"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full pl-8 p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
              placeholder="Search"
              required
              bind:value={$search}
            />
          </div>
        </form>
      </div>
    </div>
  </div>

  <!-- Rows -->
  <div class="events__rows {$rows.length === 0 && !isFiltering ? '!h-30' : 'h-96'}">
    <div class="events__rows-header">
      {#each columns as [header, style]}
        <span class={style}>
          {header.replaceAll('_', ' ').toUpperCase()}
        </span>
      {/each}
    </div>

    {#if $rows.length === 0 && isFiltering}
      <div class="!pointer-events-none !border-b-0 flex h-12 justify-center items-center">
        <span>No matching entries found</span>
      </div>
    {:else if $rows.length === 0}
      <div class="!pointer-events-none !border-b-0 flex h-12 justify-center items-center">
        <span>No resources found</span>
      </div>
    {:else}
      {#each $rows as row}
        <div id={row.resource.metadata?.uid} class="events__rows-item">
          {#each columns as [key, style]}
            <!-- Check object to avoid issues with `false` values -->
            {@const value = Object.hasOwn(row.table, key) ? row.table[key] : ''}
            <span class={`${style} ${calculateTypeClass(key, row.table[key])}`}>
              {value.text || (value === 0 ? '0' : value) || '-'}
            </span>
          {/each}
        </div>
      {/each}
    {/if}
  </div>

  <!-- Footer with link-->
  <div class="events__footer">
    <button
      class="text-sm text-blue-500 dark:text-blue-300 flex items-center space-x-1"
      on:click={() => goto('/monitor/events')}
    >
      <span>VIEW EVENTS</span>
      <ChevronRight />
    </button>
  </div>
</div>

<style lang="postcss">
  .events {
    @apply bg-white dark:bg-gray-800 w-full relative shadow-md rounded-lg overflow-hidden mt-10;
  }

  .events__header {
    @apply p-6 dark:border-gray-700 flex items-start;
  }

  .events__header-title {
    @apply text-lg dark:text-white font-semibold justify-items-start;
  }

  .events__rows {
    @apply flex flex-col text-xs overflow-hidden overflow-y-scroll;
  }

  .events__rows-header {
    @apply flex justify-start text-gray-700 uppercase bg-gray-100 dark:bg-gray-700 dark:text-gray-400 px-6 py-5;
  }

  .events__rows-item {
    @apply flex p-6 border-b border-b-gray-600;
  }

  .events__footer {
    @apply bg-white dark:bg-gray-800 rounded-b-lg px-10 h-20 flex items-center justify-end;
  }
</style>
