<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import type { CoreV1Event, KubernetesObject } from '@kubernetes/client-node'
  import { Event } from '$components'

  export let events: CoreV1Event[]
  export let resource: KubernetesObject

  let filteredEvents: CoreV1Event[] = []

  $: filteredEvents =
    events?.filter((event: CoreV1Event) => event.involvedObject.name === resource.metadata?.name) || []
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
