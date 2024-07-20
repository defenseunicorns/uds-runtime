<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { page } from '$app/stores'

  import { ChevronRight, DocumentMultiple_01, Help, SettingsAdjust, SettingsEdit } from 'carbon-icons-svelte'
  import { routes } from '../routes'
  import { isSidebarExpanded } from '../store'
  import './styles.postcss'

  const toggleSubmenus: Record<string, boolean> = {}

  routes.forEach((route) => {
    toggleSubmenus[route.path] = $page.url.pathname.includes(route.path)
  })

  let filtered = routes

  // Filter routes, if matching parent, show all children
  function filterRoutes(event: KeyboardEvent) {
    const filter = (event.target as HTMLInputElement).value.toLowerCase()
    filtered = routes
      // Deep-cloning routes to avoid modifying the original array
      .map((route) => ({ ...route }))
      // Filter routes based on the search query
      .filter((route) => {
        // If the parent route matches the search query, show all children
        if (route.name.toLowerCase().includes(filter)) {
          return true
        }

        // If the parent route doesn't match the search query, filter children
        if (route.children) {
          route.children = route.children.filter((child) => child.name.toLowerCase().includes(filter))
          return route.children.length > 0
        }
      })
      // Show all children of the matching parent
      .map((route) => {
        toggleSubmenus[route.path] = true
        return route
      })
  }
</script>

<aside
  id="main-sidebar"
  class="fixed left-0 top-14 z-40 h-screen -translate-x-full transition-all duration-300 ease-in-out sm:translate-x-0 hover:w-64 {$isSidebarExpanded
    ? 'w-64'
    : 'w-16'}"
>
  <div
    class="h-full overflow-y-auto border-r border-gray-200 bg-white px-3 py-5 dark:border-gray-700 dark:bg-gray-800 flex flex-col"
  >
    <div class="flex items-center mb-4">
      <div class="relative w-full">
        <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
          <svg
            class="w-5 h-5 text-gray-500 dark:text-gray-400"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
            ><path
              d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z"
            ></path></svg
          >
        </div>
        <input
          type="search"
          id="sidebar-filter"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Filter Pages"
          on:keyup={filterRoutes}
        />
      </div>
    </div>
    <ul class="space-y-2">
      {#each filtered as route}
        <li class={route.class}>
          {#if route.children}
            <button
              type="button"
              class="group flex w-full items-center rounded-lg p-2 text-base font-normal text-gray-900 transition duration-300 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-600"
              on:click={() => (toggleSubmenus[route.path] = !toggleSubmenus[route.path])}
            >
              <svelte:component this={route.icon} class="icon" />
              <span class="expanded-only ml-3 flex-1 whitespace-nowrap text-left">{route.name}</span>
              <ChevronRight
                class="expanded-only h-6 w-6 transition duration-300 {toggleSubmenus[route.path]
                  ? 'rotate-90 transform'
                  : ''}"
              />
            </button>
            <ul class="expanded-only space-y-2 py-2 {toggleSubmenus[route.path] ? '' : 'hidden'}">
              {#each route.children as child}
                <li>
                  <a
                    href={route.path + child.path}
                    class="group flex w-full items-center rounded-lg p-2 pl-11 text-base font-light text-gray-900 transition duration-300 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-600"
                    class:active={$page.url.pathname.includes(route.path + child.path)}>{child.name}</a
                  >
                </li>
              {/each}
            </ul>
          {:else}
            <a
              href={route.path}
              class="group flex items-center rounded-lg p-2 text-base font-normal text-gray-900 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-600"
              class:active={$page.url.pathname === route.path}
            >
              <svelte:component this={route.icon} class="icon" />
              <span class="expanded-only ml-3">{route.name}</span>
            </a>
          {/if}
        </li>
      {/each}
    </ul>
    <ul class="mt-5 space-y-2 border-t border-gray-200 pt-5 dark:border-gray-700">
      <li>
        <a
          href="/docs"
          class="group flex items-center rounded-lg p-2 text-base font-normal text-gray-900 transition duration-300 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-600"
        >
          <DocumentMultiple_01 class="icon" />
          <span class="expanded-only ml-3">Docs</span>
        </a>
      </li>
    </ul>
    <div class="grow"></div>
    <div id="sidebar-footer" class="flex hidden justify-center bg-white mb-16 mt-8 lg:flex dark:bg-gray-800">
      <a
        data-testid="global-sidenav-preferences"
        href="/preferences"
        class="icon inline-flex cursor-pointer justify-center rounded p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-900 dark:text-gray-400 dark:hover:bg-gray-600 dark:hover:text-white"
      >
        <SettingsAdjust class="h-6 w-6" />
      </a>
      <a
        data-testid="global-sidenav-settings"
        href="/settings"
        class="icon inline-flex cursor-pointer justify-center rounded p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-900 dark:text-gray-400 dark:hover:bg-gray-600 dark:hover:text-white"
      >
        <SettingsEdit class="h-6 w-6" />
      </a>
      <a
        data-testid="global-sidenav-help"
        href="/help"
        class="icon inline-flex cursor-pointer justify-center rounded p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-900 dark:text-gray-400 dark:hover:bg-gray-600 dark:hover:text-white"
      >
        <Help class="h-6 w-6" />
      </a>
    </div>
  </div>
</aside>
