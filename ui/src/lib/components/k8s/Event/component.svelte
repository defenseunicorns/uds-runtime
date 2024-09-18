<script lang="ts">
  import type { CoreV1Event } from '@kubernetes/client-node'
  import { ChevronRight } from 'carbon-icons-svelte'

  export let resource: CoreV1Event

  let toggled = false

  const renderSource = () => {
    if (resource.source && Object.keys(resource.source).length !== 0) {
      return `${resource.source.component} ${resource.source.host}`
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
      <span class={`${resource.type === 'Warning' ? 'text-red-600' : ''}`}>{resource.message}</span>
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
        <dd class="text-gray-400">{resource.count || '-'}</dd>
      </div>

      <div class="flex flex-col sm:flex-row gap-9 border-b border-gray-700 pb-2">
        <dt class="font-bold text-sm flex-none w-[180px]">Sub-object</dt>
        <dd class="text-gray-400">{resource.involvedObject.fieldPath || '-'}</dd>
      </div>
    </dl>
  </div>
</div>
