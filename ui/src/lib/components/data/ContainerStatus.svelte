<script lang="ts">
  import { type V1ContainerStatus } from '@kubernetes/client-node/dist/gen/models/V1ContainerStatus'

  export let containers: V1ContainerStatus[] = []
</script>

<div class="flex items-center space-x-1">
  {#each containers as c}
    <span
      class="w-2 h-2 rounded-sm"
      class:bg-green-500={c.state?.running}
      class:bg-yellow-500={c.state?.waiting}
      class:bg-gray-500={c.state?.terminated}
      class:animate-pulse={c.state?.running && !c.ready}
    ></span>
  {/each}
</div>
