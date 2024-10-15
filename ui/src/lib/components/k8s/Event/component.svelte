<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import type { CoreV1Event } from '@kubernetes/client-node'
  import { ChevronRight } from 'carbon-icons-svelte'

  export let event: CoreV1Event

  let toggled = false

  const renderSource = () => {
    if (event.source && Object.keys(event.source).length !== 0) {
      return `${event.source.component} ${event.source.host}`
    }

    return '<unknown>'
  }
</script>

<div class="flex flex-col">
  <div class="flex">
    <button class="mr-2" on:click={() => (toggled = !toggled)}>
      <ChevronRight class="expanded-only h-6 w-6 transition {toggled ? 'rotate-90 transform' : ''} duration-300" />
    </button>
    <div>
      <span class={`${event.type === 'Warning' ? 'text-red-600' : ''}`}>{event.message}</span>
    </div>
  </div>

  <div class={`${toggled ? '' : 'hidden'}`}>
    <dl class="mt-3 space-y-4 ml-6">
      <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
        <dt class="font-bold text-sm flex-none w-[180px]">Source</dt>
        <dd class="text-gray-400">{renderSource()}</dd>
      </div>

      <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
        <dt class="font-bold text-sm flex-none w-[180px]">Count</dt>
        <dd class="text-gray-400">{event?.count || '-'}</dd>
      </div>

      <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
        <dt class="font-bold text-sm flex-none w-[180px]">Sub-object</dt>
        <dd class="text-gray-400">{event.involvedObject?.fieldPath || '-'}</dd>
      </div>
    </dl>
  </div>
</div>
