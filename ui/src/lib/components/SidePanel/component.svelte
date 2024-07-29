<script lang="ts">
  import { goto } from '$app/navigation'
  import type { KubernetesObject } from '@kubernetes/client-node'
  import { Close } from 'carbon-icons-svelte'

  export let resource: KubernetesObject
  export let baseURL: string

  console.log(resource)
</script>

<div
  class="fixed top-14 right-0 z-40 h-screen overflow-y-auto transition-transform w-2/5 dark:bg-gray-800 shadow-lg transform transition-transform duration-300 ease-in-out"
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

    <!-- Content area -->
    <div class="flex-grow overflow-y-auto p-4 dark:text-gray-300">
      <pre><code>{JSON.stringify(resource, null, 2)}</code></pre>
    </div>
  </div>
</div>
