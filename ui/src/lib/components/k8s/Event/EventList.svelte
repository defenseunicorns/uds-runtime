<script lang="ts">
  import type { CoreV1Event } from '@kubernetes/client-node'
  import { Event } from '$components'

  export let events: { resource: CoreV1Event }[]
</script>

<div class="m-6">
  {#each events as event}
    <div class="my-6">
      <Event resource={event.resource} />
    </div>
  {/each}
</div>
