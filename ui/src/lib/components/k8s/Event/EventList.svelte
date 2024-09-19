<script lang="ts">
  import type { CoreV1Event, KubernetesObject } from '@kubernetes/client-node'
  import { Event } from '$components'

  export let events: CoreV1Event[]
  export let resource: CoreV1Event

  let filteredEvents = events.filter((event: CoreV1Event) => event.involvedObject.name === resource.metadata?.name)
</script>

<div class="m-6">
  {#if filteredEvents.length === 0}
    <span>No events.</span>
  {:else}
    {#each filteredEvents as event}
      <div class="my-6">
        <Event {event} />
      </div>
    {/each}
  {/if}
</div>
