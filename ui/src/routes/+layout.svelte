<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import 'flowbite'

  import { onDestroy, onMount } from 'svelte'

  import { afterNavigate } from '$app/navigation'
  import { page } from '$app/stores'
  import { authenticated } from '$features/api-auth/store'
  import { isSidebarExpanded, Navbar, Sidebar } from '$features/navigation'
  import { ToastPanel } from '$features/toast'
  import { initFlowbite } from 'flowbite'

  import '../app.postcss'

  import { ClusterCheck } from '$lib/utils/cluster-check/cluster-check'

  let clusterCheck: ClusterCheck
  let currRoute: string

  // These initiFlowbite calls help load the js necessary to target components which use flowbite js
  // i.e. data-dropdown-toggle
  onMount(() => {
    initFlowbite()
  })

  onDestroy(() => {
    if (clusterCheck) clusterCheck.close()
  })

  afterNavigate(initFlowbite)

  $: if ($authenticated) {
    clusterCheck = new ClusterCheck()
  }

  $: {
    currRoute = $page.route?.id || '/'
  }
</script>

<Navbar />

<!-- Hide Sidebar if api auth is enabled and user is not authenticated-->
{#if $authenticated}
  <Sidebar />
{/if}

<main
  class="flex h-screen flex-col pt-16 transition-all duration-300 ease-in-out {$isSidebarExpanded
    ? 'md:ml-64'
    : 'md:ml-16'}"
>
  <div class="flex-grow {currRoute !== '/' ? 'overflow-hidden' : ''} p-4 pt-6">
    <ToastPanel />
    <slot />
  </div>
</main>
