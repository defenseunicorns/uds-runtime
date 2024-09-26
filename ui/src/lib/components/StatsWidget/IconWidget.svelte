<script lang="ts">
  import { goto } from '$app/navigation'
  import { Card } from '$components'
  import { type CarbonIcon } from 'carbon-icons-svelte'

  export let statText: string
  export let helperText: string
  export let icon: typeof CarbonIcon
  export let link: string = ''
</script>

<Card>
  <button on:click={() => goto(link)}>
    <div class="flex items-center justify-start space-x-4">
      <div class="bg-blue-900 p-3 rounded-md">
        <svelte:component this={icon} size={24} color="#3f83f8" />
      </div>

      <div class="flex flex-col items-start">
        <dt
          class="text-3xl font-semibold text-blue-500 dark:text-white truncate"
          data-testid={`resource-count-${helperText.split(' ')[0].toLowerCase()}`}
        >
          {statText}
        </dt>
        <dd class="mt-1 text-sm text-gray-900 dark:text-gray-400">
          {helperText}
        </dd>
      </div>
    </div>
  </button>
</Card>
