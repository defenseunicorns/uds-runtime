<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { KubernetesObject } from '@kubernetes/client-node'
  import { Close } from 'carbon-icons-svelte'
  import { onMount } from 'svelte'

  import { goto } from '$app/navigation'
  import './styles.postcss'

  export let resource: KubernetesObject
  export let baseURL: string

  type Tab = 'metadata' | 'yaml' | 'events'

  onMount(() => {
    const handleKeydown = (e: KeyboardEvent) => {
      const tabList: Tab[] = ['metadata', 'yaml', 'events']
      let targetTab: string | undefined

      switch (e.key) {
        // If the Escape key is pressed, close the panel by navigating to the base URL
        case 'Escape':
          goto(baseURL)
          return

        // If the left arrow key is pressed, move to the previous tab
        case 'ArrowLeft':
          targetTab = tabList[tabList.indexOf(activeTab) - 1]
          break

        // If the right arrow key is pressed, move to the next tab
        case 'ArrowRight':
          targetTab = tabList[tabList.indexOf(activeTab) + 1]
          break
      }

      // Only update the active tab if the target tab is valid
      if (targetTab) {
        activeTab = targetTab as Tab
      }
    }

    // Add the event listener when the component is mounted
    window.addEventListener('keydown', handleKeydown)

    // Clean up the event listener when the component is destroyed
    return () => {
      window.removeEventListener('keydown', handleKeydown)
    }
  })

  function formatDate(dateString: string) {
    return new Date(dateString).toLocaleString()
  }

  const details = [
    { label: 'Created', value: formatDate(resource.metadata?.creationTimestamp as unknown as string) },
    { label: 'Name', value: resource.metadata?.name },
    { label: 'Namespace', value: resource.metadata?.namespace },
    { label: 'Controller', value: resource.metadata?.ownerReferences?.[0]?.name },
  ]

  let activeTab: Tab = 'metadata'

  function setActiveTab(evt: Event) {
    const target = evt.target as HTMLButtonElement
    activeTab = target.id as Tab
  }
</script>

<div
  class="fixed top-16 right-0 z-40 h-screen overflow-y-auto transition-transform w-1/2 dark:bg-gray-800 shadow-2xl shadow-black/80 transform transition-transform duration-300 ease-in-out"
>
  <div class="flex flex-col h-full">
    <!-- Dark header -->
    <div class="bg-gray-900 text-white p-4">
      <div class="flex justify-between items-center">
        <h2 class="text-xl font-semibold">{resource.metadata?.name}</h2>
        <button
          type="button"
          class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 absolute top-2.5 end-2.5 inline-flex items-center justify-center dark:hover:bg-gray-600 dark:hover:text-white"
          on:click={() => goto(baseURL)}
        >
          <Close />
        </button>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-gray-800 text-sm font-medium text-center text-gray-400">
      <ul class="flex" id="drawer-tabs">
        <li class="mr-2">
          <button id="metadata" class:active={activeTab === 'metadata'} on:click={setActiveTab}>Metadata</button>
        </li>
        <li class="mr-2">
          <button id="yaml" class:active={activeTab === 'yaml'} on:click={setActiveTab}>YAML</button>
        </li>
        <li class="mr-2">
          <button id="events" class:active={activeTab === 'events'} on:click={setActiveTab}>Events</button>
        </li>
      </ul>
    </div>

    <!-- Content -->

    <div class="flex-grow overflow-y-auto dark:text-gray-300">
      {#if activeTab === 'metadata'}
        <!-- Metadata tab -->
        <div class="bg-gray-800 text-gray-200 p-6 rounded-lg shadow-lg">
          <dl class="space-y-4">
            {#each details as { label, value }}
              <div class="flex flex-col sm:flex-row sm:justify-between border-b border-gray-700 pb-2">
                <dt class="font-bold text-sm sm:w-1/3">{label}</dt>
                <dd class="text-gray-400 sm:w-2/3">{value || 'N/A'}</dd>
              </div>
            {/each}

            <div class="flex flex-col sm:flex-row sm:justify-between border-b border-gray-700 pb-2">
              <dt class="font-bold text-sm sm:w-1/3">Labels</dt>
              <dd class="sm:w-2/3">
                <div class="flex flex-wrap gap-2">
                  {#each Object.entries(resource.metadata?.labels || {}) as [key, value]}
                    <span class="bg-gray-600 px-2 py-0.5 rounded text-white text-xs">{key}: {value}</span>
                  {/each}
                </div>
              </dd>
            </div>

            <div class="flex flex-col sm:flex-row sm:justify-between">
              <dt class="font-bold text-sm sm:w-1/3">Annotations</dt>
              <dd class="sm:w-2/3">
                <div class="flex flex-wrap gap-2">
                  {#each Object.entries(resource.metadata?.annotations || {}) as [key, value]}
                    <span class="bg-gray-600 px-2 py-0.5 rounded text-white text-xs">{key}: {value}</span>
                  {/each}
                </div>
              </dd>
            </div>
          </dl>
        </div>
      {:else if activeTab === 'yaml'}
        <!-- YAML tab -->
        <pre class="bg-gray-800 p-6 rounded-lg shadow-lg">{JSON.stringify(resource, null, 2)}</pre>
      {:else if activeTab === 'events'}
        <!-- Events tab -->
        <div class="bg-gray-800 text-gray-200 p-6 rounded-lg shadow-lg">
          <p>Events tab content</p>
        </div>
      {/if}
    </div>
  </div>
</div>
