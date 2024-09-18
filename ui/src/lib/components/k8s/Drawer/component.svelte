<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

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

  interface TempDataType {
    resource: CoreV1Event
  }

  const tempData: TempDataType[] = [
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          fieldPath: 'spec.containers{podinfo}',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-876rf',
          namespace: 'podinfo',
          resourceVersion: '91375',
          uid: 'c7e2f3a5-14a3-44ed-a427-1d886e7ad618',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message:
          'Readiness probe failed: 2024-09-18T14:50:07.995Z\tINFO\tpodcli/check.go:137\tcheck failed\t{"address": "http://localhost:9898/readyz", "status code": 503}\n',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-876rf.17f65e222bbbba52',
          namespace: 'podinfo',
          resourceVersion: '101934',
          uid: 'a45052c1-e115-4727-b783-62292c040a56',
        },
        reason: 'Unhealthy',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'kubelet',
          host: 'k3d-runtime-server-0',
        },
        type: 'Warning',
      },
    },
    {
      resource: {
        action: 'Binding',
        apiVersion: 'v1',
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-86fjg',
          namespace: 'podinfo',
          resourceVersion: '101919',
          uid: 'caff0343-c3d7-42bf-ab49-6a38f2ad318b',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Successfully assigned podinfo/podinfo-b4d7c7fd5-86fjg to k3d-runtime-server-0',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-86fjg.17f65e21c0312274',
          namespace: 'podinfo',
          resourceVersion: '101923',
          uid: 'e18ee072-6488-4fad-896a-465534021899',
        },
        reason: 'Scheduled',
        reportingComponent: 'default-scheduler',
        reportingInstance: 'default-scheduler-k3d-runtime-server-0',
        source: {},
        type: 'Normal',
      },
    },
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          fieldPath: 'spec.containers{podinfo}',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-86fjg',
          namespace: 'podinfo',
          resourceVersion: '101921',
          uid: 'caff0343-c3d7-42bf-ab49-6a38f2ad318b',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Container image "ghcr.io/stefanprodan/podinfo:6.4.0" already present on machine',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-86fjg.17f65e21d879569b',
          namespace: 'podinfo',
          resourceVersion: '101928',
          uid: 'df645d11-52d3-4414-9646-aac1f15a0db0',
        },
        reason: 'Pulled',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'kubelet',
          host: 'k3d-runtime-server-0',
        },
        type: 'Normal',
      },
    },
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          fieldPath: 'spec.containers{podinfo}',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-86fjg',
          namespace: 'podinfo',
          resourceVersion: '101921',
          uid: 'caff0343-c3d7-42bf-ab49-6a38f2ad318b',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Created container podinfo',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-86fjg.17f65e21d923d387',
          namespace: 'podinfo',
          resourceVersion: '101929',
          uid: '9494dcb0-4248-4118-a3b2-fdf7313a3f58',
        },
        reason: 'Created',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'kubelet',
          host: 'k3d-runtime-server-0',
        },
        type: 'Normal',
      },
    },
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          fieldPath: 'spec.containers{podinfo}',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-86fjg',
          namespace: 'podinfo',
          resourceVersion: '101921',
          uid: 'caff0343-c3d7-42bf-ab49-6a38f2ad318b',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Started container podinfo',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-86fjg.17f65e21dc959efe',
          namespace: 'podinfo',
          resourceVersion: '101930',
          uid: '57053521-ee9d-462a-b2e1-39986d8d383e',
        },
        reason: 'Started',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'kubelet',
          host: 'k3d-runtime-server-0',
        },
        type: 'Normal',
      },
    },
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'v1',
          fieldPath: 'spec.containers{podinfo}',
          kind: 'Pod',
          name: 'podinfo-b4d7c7fd5-876rf',
          namespace: 'podinfo',
          resourceVersion: '91375',
          uid: 'c7e2f3a5-14a3-44ed-a427-1d886e7ad618',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Stopping container podinfo',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5-876rf.17f65e21b7e3470f',
          namespace: 'podinfo',
          resourceVersion: '101916',
          uid: '5daf4e02-d95e-4eff-9098-3279d905966b',
        },
        reason: 'Killing',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'kubelet',
          host: 'k3d-runtime-server-0',
        },
        type: 'Normal',
      },
    },
    {
      resource: {
        apiVersion: 'v1',
        count: 1,
        eventTime: new Date(),
        firstTimestamp: new Date(),
        involvedObject: {
          apiVersion: 'apps/v1',
          kind: 'ReplicaSet',
          name: 'podinfo-b4d7c7fd5',
          namespace: 'podinfo',
          resourceVersion: '91392',
          uid: '194dd699-0480-481a-a2b4-b85df8174655',
        },
        kind: 'Event',
        lastTimestamp: new Date(),
        message: 'Created pod: podinfo-b4d7c7fd5-86fjg',
        metadata: {
          creationTimestamp: new Date(),
          name: 'podinfo-b4d7c7fd5.17f65e21bfde302c',
          namespace: 'podinfo',
          resourceVersion: '101920',
          uid: 'e014093a-186c-4182-b6d5-937cc56a8d44',
        },
        reason: 'SuccessfulCreate',
        reportingComponent: '',
        reportingInstance: '',
        source: {
          component: 'replicaset-controller',
        },
        type: 'Normal',
      },
    },
  ]

  onMount(() => {
    // initialize highlight language
    hljs.registerLanguage('yaml', yaml)

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
        <EventList events={tempData} />
      {:else if activeTab === 'yaml'}
        <!-- YAML tab -->
        <div class="text-gray-200 p-4">
          <code class="text-sm text-gray-500 dark:text-gray-400 whitespace-pre w-full block">
            <!-- We turned off svelte/no-at-html-tags eslint rule because we are using DOMPurify to sanitize -->
            {@html DOMPurify.sanitize(hljs.highlight(YAML.stringify(resource), { language: 'yaml' }).value)}
          </code>
        </div>
      {/if}
    </div>
  </div>
</div>
