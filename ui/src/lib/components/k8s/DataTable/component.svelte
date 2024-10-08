<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { type Unsubscriber } from 'svelte/store'

  import type { KubernetesObject } from '@kubernetes/client-node'
  import { goto } from '$app/navigation'
  import { page } from '$app/stores'
  import { Drawer, Link, Tooltip } from '$components'
  import type { Row as NamespaceRow } from '$features/k8s/namespaces/store'
  import { type ResourceStoreInterface } from '$features/k8s/types'
  import { addToast } from '$features/toast'
  import { ChevronDown, ChevronUp, Filter, Information, Search } from 'carbon-icons-svelte'

  // Determine if the data is namespaced
  export let isNamespaced = true

  // Disable row click
  export let disableRowClick = false

  // We have to be a bit generic here to handle the various Column/Row types coming from the various stores
  export let columns: [name: string, styles?: string, modifier?: 'link-external' | 'link-internal'][]

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  export let createStore: () => ResourceStoreInterface<KubernetesObject, any>

  // name and description of K8s resource
  export let name = ''
  export let description = 'No description available'

  let rows = createStore()
  $: ({ namespace, search, searchBy, searchTypes, sortAsc, sortBy, numResources } = rows)

  // check for filtering
  let isFiltering = false
  $: {
    isFiltering = !!($search || $namespace)
  }

  let resource: KubernetesObject | null = null
  let baseURL: string
  let pathName: string
  let unsubscribePage: Unsubscriber
  let uid = ''
  let namespaces: ResourceStoreInterface<KubernetesObject, NamespaceRow>

  onDestroy(() => {
    unsubscribePage()
  })

  unsubscribePage = page.subscribe(async ({ params, data, url }) => {
    namespaces = data.namespaces as ResourceStoreInterface<KubernetesObject, NamespaceRow>
    uid = params.uid || ''

    pathName = url.pathname

    // If UID is present, load the data
    if (uid) {
      try {
        // Strip the UID from the URL
        baseURL = pathName.replace(`/${uid}`, '')

        // Split because new URL() doesn't work without a complete URL
        const [apiPath] = rows.url.split('?')

        // Fetch the resource data
        let results
        results = await fetch(`${apiPath}/${uid}`)

        // If the fetch is successful, set the resource data
        if (results.ok) {
          const data = await results.json()
          resource = data.Object as KubernetesObject
          return
        } else {
          // Otherwise, throw an error
          throw new Error(`Failed to fetch resource: ${results.statusText}`)
        }
      } catch (e) {
        // If an error occurs, set the resource to null
        resource = null

        // Display an error toast if the fetch fails
        addToast({
          timeoutSecs: 5,
          message: e.message,
          type: 'error',
        })
      }
    } else {
      // If no UID is present, the path is the base URL and the resource is null
      baseURL = pathName
      resource = null
    }
  })

  onMount(() => {
    let stop = rows.start()

    // Function to navigate using the keyboard
    const keyboardNavigate = (e: KeyboardEvent) => {
      if (uid) {
        let nextID: string | undefined

        switch (e.key) {
          case 'ArrowDown': {
            nextID = document.getElementById(uid)?.nextElementSibling?.id
            break
          }
          case 'ArrowUp': {
            nextID = document.getElementById(uid)?.previousElementSibling?.id
            break
          }
        }

        if (nextID) {
          // Navigate to the next row
          goto(`${baseURL}/${nextID}`)
        }
      }
    }

    // Function to restart the store on cluster reconnection
    const handleClusterReconnected = () => {
      console.log('Cluster reconnected, restarting store')
      // stop current store first
      stop()

      // recreate rows to trigger re-render
      rows = createStore()
      stop = rows.start()
    }

    // Bind the keyboard navigation event
    window.addEventListener('keydown', keyboardNavigate)
    window.addEventListener('cluster-reconnected', handleClusterReconnected)

    // Clean up the event listener when the component is destroyed
    return () => {
      window.removeEventListener('keydown', keyboardNavigate)
      window.removeEventListener('cluster-reconnected', handleClusterReconnected)
      stop()
    }
  })
</script>

{#if resource}
  <Drawer {resource} {baseURL} />
{/if}

<section class="table-section">
  <div class="table-container">
    <div class="table-content">
      <div class="table-header">
        <span class="dark:text-white" data-testid="table-header">{name}</span>
        {#if isFiltering}
          <span class="dark:text-gray-500 pl-2" data-testid="table-header-results">
            (showing {$rows.length} of {$numResources})
          </span>
        {:else}
          <span class="dark:text-gray-500 pl-2" data-testid="table-header-results">({$numResources})</span>
        {/if}
        <div class="relative group">
          <Information class="ml-2 w-4 h-4 text-gray-400" />
          <div class="tooltip tooltip-right min-w-72">
            <div class="whitespace-normal">{description}</div>
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
        <button
          id="filterDropdownButton"
          data-dropdown-toggle="filterDropdown"
          class="hover:text-primary-700 flex items-center justify-center rounded-lg border border-gray-200 bg-white px-4 py-2 text-sm font-medium text-gray-900 hover:bg-gray-100 focus:z-10 focus:outline-none focus:ring-4 focus:ring-gray-200 md:w-auto dark:border-gray-600 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:ring-gray-700"
          type="button"
          data-testid="datatable-filter-dropdown"
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
          {#if isNamespaced}
            <select id="stream" bind:value={$namespace} data-testid="table-filter-namespace-select">
              <option value="" data-testid="namespace-select-all">All Namespaces</option>
              <hr />
              {#each $namespaces as ns}
                <option value={ns.table.name} data-testid={`namespace-select-${ns.table.name}`}>
                  {ns.table.name}
                </option>
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
            {#if $rows.length === 0 && isFiltering}
              <tr class="!pointer-events-none !border-b-0">
                <td class="text-center" colspan="9">No matching entries found</td>
              </tr>
            {:else if $rows.length === 0}
              <tr class="!pointer-events-none !border-b-0">
                <td class="text-center" colspan="9">No resources found</td>
              </tr>
              <!--This case only happens when returning an error when a CRD does not exist in the cluster-->
            {:else if typeof $rows[0].resource === 'string'}
              <tr class="!pointer-events-none !border-b-0">
                <td class="text-center" colspan="9">
                  The CRD for the resources you are trying to view does not exist in the cluster.
                  <br />
                  Install CRD and refresh page to view resources.
                </td>
              </tr>
            {:else}
              {#each $rows as row}
                <tr
                  id={row.resource.metadata?.uid}
                  on:click={() =>
                    !disableRowClick && row.resource.metadata?.uid && goto(`${baseURL}/${row.resource.metadata?.uid}`)}
                  class:active={row.resource.metadata?.uid && pathName.includes(row.resource.metadata?.uid ?? '')}
                  class:cursor-pointer={!disableRowClick}
                >
                  {#each columns as [key, style, modifier], idx}
                    <!-- Check object to avoid issues with `false` values -->
                    {@const value = Object.hasOwn(row.table, key) ? row.table[key] : ''}
                    <td
                      class={style || ''}
                      data-testid={typeof value !== 'object'
                        ? `${value}-testid-${idx + 1}`
                        : `object-test-id-${idx + 1}`}
                    >
                      {#if value.component}
                        <svelte:component this={value.component} {...value.props} />
                      {:else if value.list}
                        <ul class="line-clamp-4 mt-4 text-sm">
                          {#each value.list as item}
                            <li data-testid={`${item}-list-item-test-id`}>- {item}</li>
                          {/each}
                        </ul>
                      {:else if modifier === 'link-external'}
                        <Link href={value.href || ''} text={value.text || ''} target={'_blank'} />
                      {:else if modifier === 'link-internal'}
                        <Link href={value.href || ''} text={value.text || ''} target={''} />
                      {:else if key === 'namespace'}
                        <button
                          on:click|stopPropagation={() => namespace.set(value)}
                          class="text-blue-600 dark:text-blue-500 hover:underline pr-4 text-left"
                        >
                          {value}
                        </button>
                      {:else if style?.includes('truncate')}
                        <Tooltip title={value}>
                          <div class={`w-full ${style}`}>
                            {value}
                          </div>
                        </Tooltip>
                      {:else}
                        {value.text || (value === 0 ? '0' : value) || '-'}
                      {/if}
                    </td>
                  {/each}
                </tr>
              {/each}
            {/if}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</section>
