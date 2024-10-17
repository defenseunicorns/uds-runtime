<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { onMount } from 'svelte'

  import type { CoreV1Event, KubernetesObject } from '@kubernetes/client-node'
  import { goto } from '$app/navigation'
  import { EventList } from '$components'
  import { Close } from 'carbon-icons-svelte'
  import DOMPurify from 'dompurify'
  import hljs from 'highlight.js/lib/core'
  import yaml from 'highlight.js/lib/languages/yaml'
  import * as YAML from 'yaml'

  import './styles.postcss'

  export let resource: KubernetesObject
  export let baseURL: string

  type Tab = 'metadata' | 'yaml' | 'events'

  let events: CoreV1Event[] = []

  onMount(() => {
    // initialize highlight language
    hljs.registerLanguage('yaml', yaml)

    const path: string = '/api/v1/resources/events?fields=.count,.involvedObject,.message,.source,.type'
    const eventSource = new EventSource(path)

    eventSource.onmessage = (event) => {
      events = JSON.parse(event.data) as CoreV1Event[]
    }

    const handleKeydown = (e: KeyboardEvent) => {
      const tabList: Tab[] = ['metadata', 'events', 'yaml']
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
      eventSource.close()
    }
  })

  function formatDate(dateString: string) {
    return new Date(dateString).toLocaleString()
  }

  $: details = [
    { label: 'Created', value: formatDate(resource.metadata?.creationTimestamp as unknown as string) },
    { label: 'Name', value: resource.metadata?.name },
    { label: 'Namespace', value: resource.metadata?.namespace },
  ]

  if ((resource.metadata?.ownerReferences?.length && details) || 0 > 0) {
    details.push({
      label: 'Controlled By',
      value: `${resource.metadata?.ownerReferences?.[0]?.kind} ${resource.metadata?.ownerReferences?.[0]?.name}`,
    })
  }

  let activeTab: Tab = 'metadata'

  function setActiveTab(evt: MouseEvent) {
    const target = evt.target as HTMLButtonElement
    activeTab = target.id as Tab
  }
</script>

<div
  data-testid="drawer"
  class="fixed top-16 right-0 z-40 h-screen overflow-y-auto w-1/2 dark:bg-gray-800 shadow-2xl shadow-black/80 transform transition-transform duration-300 ease-in-out"
>
  <div class="flex flex-col h-full">
    <!-- Dark header -->
    <div class="bg-gray-900 text-white p-4 pb-0">
      <div class="flex justify-between items-center">
        <h2 class="text-xl">
          <span class="font-semibold">{resource.kind}:</span>
          <span>{resource.metadata?.name}</span>
        </h2>
        <button
          type="button"
          class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 absolute top-2.5 end-2.5 inline-flex items-center justify-center dark:hover:bg-gray-600 dark:hover:text-white"
          on:click={() => goto(baseURL)}
        >
          <Close />
        </button>
      </div>

      <!-- Tabs -->
      <div class="flex font-medium pt-3">
        <ul class="flex w-full" id="drawer-tabs">
          <li class="flex-1">
            <button id="metadata" class:active={activeTab === 'metadata'} on:click={setActiveTab}>Metadata</button>
          </li>
          <li class="flex-1">
            <button id="events" class:active={activeTab === 'events'} on:click={setActiveTab}>Events</button>
          </li>
          <li class="flex-1">
            <button id="yaml" class:active={activeTab === 'yaml'} on:click={setActiveTab}>YAML</button>
          </li>
        </ul>
      </div>
    </div>

    <!-- Content -->

    <div class="flex-grow overflow-y-auto dark:text-gray-300 pb-20">
      {#if activeTab === 'metadata'}
        <!-- Metadata tab -->
        <div class="bg-gray-800 text-gray-200 p-6 rounded-lg">
          <dl class="space-y-4">
            {#each details as { label, value }}
              <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
                <dt class="font-bold text-sm flex-none w-[180px]">{label}</dt>
                <dd class="text-gray-400">{value || 'N/A'}</dd>
              </div>
            {/each}

            {#if resource.metadata?.labels}
              <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
                <dt class="font-bold text-sm flex-none w-[180px]">Labels</dt>
                <dd class="overflow-x-auto">
                  <div class="flex flex-wrap gap-2">
                    {#each Object.entries(resource.metadata?.labels || {}) as [key, value]}
                      <span class="bg-gray-600 px-2 py-0.5 rounded text-white text-xs">{key}: {value}</span>
                    {/each}
                  </div>
                </dd>
              </div>
            {/if}

            {#if resource.metadata?.annotations}
              <div class="flex flex-col sm:flex-row gap-9">
                <dt class="font-bold text-sm flex-none w-[180px]">Annotations</dt>
                <dd class="overflow-x-auto">
                  <div class="flex flex-wrap gap-2">
                    {#each Object.entries(resource.metadata?.annotations || {}) as [key, value]}
                      <span class="bg-gray-600 px-2 py-0.5 rounded text-white text-xs">{key}: {value}</span>
                    {/each}
                  </div>
                </dd>
              </div>
            {/if}
          </dl>
        </div>
      {:else if activeTab === 'events'}
        <EventList {events} {resource} />
      {:else if activeTab === 'yaml'}
        <!-- YAML tab -->
        <div class="text-gray-200 p-4">
          <code class="text-sm text-gray-500 dark:text-gray-400 whitespace-pre w-full block">
            <!-- Disable svelte/no-at-html-tags eslint rule here because we are using DOMPurify to sanitize -->
            <!-- eslint-disable-next-line svelte/no-at-html-tags -->
            {@html DOMPurify.sanitize(hljs.highlight(YAML.stringify(resource), { language: 'yaml' }).value)}
          </code>
        </div>
      {/if}
    </div>
  </div>
</div>
