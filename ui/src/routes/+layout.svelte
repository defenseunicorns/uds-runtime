<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import 'flowbite'
  import { initFlowbite } from 'flowbite'
  import { onMount } from 'svelte'
  import { afterNavigate } from '$app/navigation'
  import { isSidebarExpanded, Navbar, Sidebar } from '$features/navigation'
  import { ToastPanel } from '$features/toast'
  import '../app.postcss'
  import { authenticated } from '$lib/features/api-auth/store'
  import { apiAuthEnabled } from '$lib/features/api-auth/store'
  import { addToast } from '$features/toast/store'
  import Unauthenticated from '$components/Auth/component.svelte'

  let path = ''
  // These initiFlowbite calls help load the js necessary to target components which use flowbite js
  // i.e. data-dropdown-toggle
  onMount(() => {
    initFlowbite()
    path = window.location.pathname
  })

  afterNavigate(initFlowbite)

  $: if (!$apiAuthEnabled || ($apiAuthEnabled && $authenticated)) {
    const healthCheck = new EventSource('/health')
    healthCheck.onerror = () => {
      addToast({
        type: 'error',
        message: 'Cluster health check failed: no connection',
        timeoutSecs: 10,
      })
    }
    healthCheck.onmessage = (msg) => {
      const data = JSON.parse(msg.data) as Record<'version' | 'error', string>
      if (data['error']) {
        addToast({
          type: 'error',
          message: 'Cluster health check failed: no connection',
          timeoutSecs: 10,
        })
      }
    }
  }
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
    {#if $apiAuthEnabled && !$authenticated && path !== '/auth'}
      <Unauthenticated />
    {:else}
      <slot />
    {/if}
  </div>
</main>
