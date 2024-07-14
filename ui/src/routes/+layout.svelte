<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2023-Present The UDS Authors -->

<script lang="ts">
  import 'flowbite'
  import { initFlowbite } from 'flowbite'
  import 'flowbite/dist/flowbite.css'
  import { onMount } from 'svelte'

  import { afterNavigate } from '$app/navigation'

  import { Breadcrumb, isSidebarExpanded, Navbar, Sidebar } from '$features/navigation'
  import '../app.postcss'

  onMount(initFlowbite)
  afterNavigate(initFlowbite)
</script>

<Navbar />

<Sidebar />

<main
  class="flex h-screen flex-col pt-16 transition-all duration-300 ease-in-out {$isSidebarExpanded
    ? 'md:ml-64'
    : 'md:ml-16'}"
>
  <div class="p-5">
    <Breadcrumb />
  </div>
  <div class="flex-grow overflow-hidden p-4 pt-0">
    <slot />
  </div>
</main>
