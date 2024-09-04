<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import 'flowbite'
  import { initFlowbite } from 'flowbite'
  import { onMount } from 'svelte'

  import { afterNavigate, goto } from '$app/navigation'
  import { isSidebarExpanded, Navbar, Sidebar } from '$features/navigation'
  import { ToastPanel } from '$features/toast'
  import '../app.postcss'
  import { authenticated } from '$lib/features/api-auth/store'
  import { apiAuthEnabled } from '$lib/features/api-auth/store'

  // These initiFlowbite calls help load the js necessary to target components which use flowbite js
  // i.e. data-dropdown-toggle
  onMount(() => {
    initFlowbite()
    const unsubscribe = authenticated.subscribe((value) => {
      if (!value && window.location.pathname !== '/auth') {
        goto('/auth')
      }
    })

    return unsubscribe
  })

  afterNavigate(initFlowbite)
</script>

<Navbar />

<!-- Hide Sidebar if api auth is enabled and user is not authenticated-->
{#if !$apiAuthEnabled || ($apiAuthEnabled && $authenticated)}
  <Sidebar />
{/if}

<main
  class="flex h-screen flex-col pt-16 transition-all duration-300 ease-in-out {$isSidebarExpanded
    ? 'md:ml-64'
    : 'md:ml-16'}"
>
  <div class="flex-grow overflow-hidden p-4 pt-6">
    <ToastPanel />
    {#if !$apiAuthEnabled || ($apiAuthEnabled && $authenticated)}
      <slot />
    {:else}
      <div class="flex flex-col items-center justify-start min-h-screen">
        <h2 class="text-xl mb-4 p-4 dark:text-white pt-0">
          <strong>Could not authenticate</strong>
          : Please make sure you are using the complete link with api token provided by UDS Runtime to connect.
        </h2>
        <img src="/doug.svg" alt="Authentication Failed" class="mx-auto mt-4" style="width: 250px; height: 250px;" />
      </div>
    {/if}
  </div>
</main>
