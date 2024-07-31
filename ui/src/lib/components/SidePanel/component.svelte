<script lang="ts">
  import { goto } from '$app/navigation'
  import type { KubernetesObject } from '@kubernetes/client-node'
  import { Close } from 'carbon-icons-svelte'
  import { onMount } from 'svelte'

  export let resource: KubernetesObject
  export let baseURL: string

  // If the Escape key is pressed, close the panel by navigating to the base URL
  onMount(() => {
    const handleKeydown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        goto(baseURL)
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
</script>

<div
  class="fixed top-16 right-0 z-40 h-screen overflow-y-auto transition-transform w-2/5 dark:bg-gray-800 shadow-2xl shadow-black/80 transform transition-transform duration-300 ease-in-out"
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
      <ul class="flex">
        <li class="mr-2">
          <a
            href="#a"
            class="inline-block p-4 text-blue-500 border-b-2 border-blue-500 rounded-t-lg active"
            aria-current="page">Details</a
          >
        </li>
        <li class="mr-2">
          <a
            href="#a"
            class="inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-300 hover:border-gray-300"
            >Tab 2</a
          >
        </li>
        <li class="mr-2">
          <a
            href="#a"
            class="inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-300 hover:border-gray-300"
            >Tab 3</a
          >
        </li>
      </ul>
    </div>

    <div class="flex-grow overflow-y-auto dark:text-gray-300">
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

      {#if resource.spec}
        <div class="bg-gray-800 text-gray-200 p-6">
          <h3 class="text-lg font-semibold">Spec</h3>
          <pre class="text-sm overflow-x-auto">{JSON.stringify(resource.spec, null, 2)}</pre>
        </div>
      {/if}

      {#if resource.status}
        <div class="bg-gray-800 text-gray-200 p-6">
          <h3 class="text-lg font-semibold">Status</h3>
          <pre class="text-sm overflow-x-auto">{JSON.stringify(resource.status, null, 2)}</pre>
        </div>
      {/if}
    </div>
  </div>
</div>
