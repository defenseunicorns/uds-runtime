<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import 'flowbite'
  import { initFlowbite } from 'flowbite'
  import 'flowbite/dist/flowbite.css'
  import { onMount } from 'svelte'

  import { afterNavigate } from '$app/navigation'
  import { isSidebarExpanded, Navbar, Sidebar } from '$features/navigation'
  import { ToastPanel } from '$features/toast'
  import '../app.postcss'
  import { authenticated } from '$lib/features/api-auth/store'

  let isAuthenticated: boolean
  const apiAuthEnabled = import.meta.env.VITE_API_AUTH?.toLowerCase() === 'true'

  // If API_AUTH is enabled, subscribe to the authenticated store to check if the user is authenticated
  if (apiAuthEnabled) {
    authenticated.subscribe((value) => {
      isAuthenticated = value
    })
  }

  onMount(() => {
    initFlowbite()
  })

  afterNavigate(initFlowbite)
  console.log('apiAuthEnabled:' + apiAuthEnabled)
  console.log('authenticated:' + authenticated)
</script>

<Navbar />

<!-- Hide navbar if api auth is enabled and user is not authenticated-->
{#if !apiAuthEnabled || (apiAuthEnabled && isAuthenticated)}
  <Sidebar />
{/if}

<main
  class="flex h-screen flex-col pt-16 transition-all duration-300 ease-in-out {$isSidebarExpanded
    ? 'md:ml-64'
    : 'md:ml-16'}"
>
  <div class="flex-grow overflow-hidden p-4 pt-6">
    <ToastPanel />
    <slot />
  </div>
</main>
