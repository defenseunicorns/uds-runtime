<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { KubernetesObject } from '@kubernetes/client-node'
  import { onMount } from 'svelte'

  import { page } from '$app/stores'
  import type { Row as NamespaceRow } from '$features/k8s/namespaces/store'
  import type { ResourceStoreInterface } from '$features/k8s/store'
  import { ChevronDown, ChevronUp, Filter, Search } from 'carbon-icons-svelte'

  // We have to be a bit generic here to handle the various Column/Row types coming from the various stores
  export let columns: [name: string, styles?: string][]
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  export let createStore: () => ResourceStoreInterface<KubernetesObject, any>

  // Load the namespaces from the page store
  const namespaces = $page.data.namespaces as ResourceStoreInterface<KubernetesObject, NamespaceRow>

  const rows = createStore()
  const { namespace, search, searchBy, searchTypes, sortAsc, sortBy } = rows

  // If the route is a namespace route, we don't want to show the namespace dropdown
  let isNamespace = $page.url.pathname.includes('/namespaces')

  onMount(() => {
    return rows.start()
  })
</script>

<section class="table-section">
  <div class="table-container">
    <div class="table-content">
      <div class="table-filter-section">
        <div class="relative lg:w-96">
          <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <Search class="h-5 w-5 text-gray-400" />
          </div>
          <input
            type="text"
            name="email"
            class="focus:ring-primary-500 focus:border-primary-500 dark:focus:ring-primary-500 dark:focus:border-primary-500 block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 pl-9 text-gray-900 sm:text-sm dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400"
            placeholder="Search"
            bind:value={$search}
          />
        </div>
        <button
          id="filterDropdownButton"
          data-dropdown-toggle="filterDropdown"
          class="hover:text-primary-700 flex items-center justify-center rounded-lg border border-gray-200 bg-white px-4 py-2 text-sm font-medium text-gray-900 hover:bg-gray-100 focus:z-10 focus:outline-none focus:ring-4 focus:ring-gray-200 md:w-auto dark:border-gray-600 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:ring-gray-700"
          type="button"
        >
          <Filter class="mr-2 h-4 w-4 text-gray-400" />
          {$searchBy}
          <ChevronDown class="ml-2 h-4 w-4 text-gray-400" />
        </button>
        <div id="filterDropdown" class="z-10 hidden w-48 rounded-lg bg-white p-3 shadow dark:bg-gray-700">
          <h6 class="mb-3 text-sm font-medium text-gray-900 dark:text-white">Search By</h6>
          <ul class="space-y-2 text-sm" aria-labelledby="filterDropdownButton">
            {#each searchTypes as searchType}
              <li class="flex items-center">
                <input
                  id={searchType}
                  type="radio"
                  name="searchType"
                  value={searchType}
                  class="h-4 w-4 border-gray-300 focus:ring-2 focus:ring-blue-300 dark:border-gray-600 dark:bg-gray-700 dark:focus:bg-blue-600 dark:focus:ring-blue-600"
                  bind:group={$searchBy}
                />
                <label for={searchType} class="ms-2 block text-sm font-medium text-gray-900 dark:text-gray-300">
                  {searchType}
                </label>
              </li>
            {/each}
          </ul>
        </div>
        <div class="flex-grow"></div>
        <div>
          {#if !isNamespace}
            <select id="stream" bind:value={$namespace}>
              <option value="">All Namespaces</option>
              <hr />
              {#each $namespaces as ns}
                <option value={ns.table.name}>{ns.table.name}</option>
              {/each}
            </select>
          {/if}
        </div>
      </div>
      <div class="table-scroll-container">
        <table>
          <thead>
            <tr>
              {#each columns as [header, style]}
                <th>
                  <button on:click={() => rows.sortByKey(header)}>
                    {header.replaceAll('_', ' ')}
                    <ChevronUp
                      class="sort
                                  {style || ''}
                                  {$sortAsc ? 'rotate-180' : ''}
                                  {$sortBy === header ? 'opacity-100' : 'opacity-0'}"
                    />
                  </button>
                </th>
              {/each}
            </tr>
          </thead>
          <tbody>
            {#each $rows as row}
              <tr>
                {#each columns as [key, style]}
                  {#if row.table[key].component}
                    <td class={style || ''}>
                      <svelte:component this={row.table[key].component} {...row.table[key].props} />
                    </td>
                  {:else if row.table[key].text}
                    <td class={style || ''}>{row.table[key].text}</td>
                  {:else}
                    <td class={style || ''}>{row.table[key]}</td>
                  {/if}
                {/each}
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</section>
