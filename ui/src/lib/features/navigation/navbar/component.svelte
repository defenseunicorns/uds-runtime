<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { authenticated } from '$features/auth/store'
  import { UserMenu } from '$features/navigation'
  import type { UserData } from '$features/navigation/types'
  import { NotificationFilled } from 'carbon-icons-svelte'

  import { isSidebarExpanded } from '../store'

  export let userData: UserData | null

  const inClusterAuth = (userData && userData.inClusterAuth) ?? false

  // Don't expand sidebar if api auth is enabled and user is unauthenticated
  $: {
    if ($authenticated) {
      isSidebarExpanded.set(true)
    } else {
      isSidebarExpanded.set(false)
    }
  }
</script>

<div class="bg-gray-50 antialiased dark:bg-gray-900">
  <nav
    class="fixed left-0 right-0 top-0 z-50 border-b border-gray-200 bg-white px-4 py-2.5 dark:border-gray-700 dark:bg-gray-800"
  >
    <div class="flex flex-wrap items-center justify-between">
      <div class="flex items-center justify-start">
        <!-- Hide Sidebar if api auth is enabled and user is not authenticated-->
        {#if $authenticated}
          <button
            data-testid="toggle-sidebar"
            aria-expanded="true"
            aria-controls="sidebar"
            on:click={() => isSidebarExpanded.update((v) => !v)}
            class="mr-3 hidden cursor-pointer rounded p-2 text-gray-600 hover:bg-gray-100 hover:text-gray-900 lg:inline dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
          >
            <svg class="h-5 w-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M1 1h14M1 6h14M1 11h7"
              />
            </svg>
          </button>
        {/if}

        <a href="/" class="mr-4 flex">
          <img src="/doug.svg" class="mr-3 h-8" alt="FlowBite Logo" />
          <span class="self-center whitespace-nowrap text-2xl font-semibold dark:text-white">UDS</span>
        </a>
      </div>
      <div class="flex items-center lg:order-2">
        <!-- Notifications -->
        <button
          type="button"
          data-dropdown-toggle="notification-dropdown"
          class="mr-1 rounded-lg p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-900 focus:ring-4 focus:ring-gray-300 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:ring-gray-600"
        >
          <NotificationFilled class="h-6 w-6" />
        </button>
        <!-- Dropdown menu -->
        <div
          class="z-50 my-4 hidden max-w-sm list-none divide-y divide-gray-100 overflow-hidden rounded bg-white text-base shadow-lg dark:divide-gray-600 dark:bg-gray-700"
          id="notification-dropdown"
        >
          <div
            class="block bg-gray-50 px-4 py-2 text-center text-base font-medium text-gray-700 dark:bg-gray-700 dark:text-gray-400"
          >
            Notifications
          </div>
        </div>
        {#if inClusterAuth}
          <UserMenu {userData} />
        {/if}
      </div>
    </div>
  </nav>
</div>
