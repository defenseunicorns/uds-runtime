<script lang="ts">
  import { type V1ContainerStatus } from '@kubernetes/client-node/dist/gen/models/V1ContainerStatus'

  export let containers: V1ContainerStatus[] = []
</script>

<div class="flex items-center space-x-1">
  {#each containers as c}
    <div
      class="w-2 h-2 rounded-sm relative group"
      class:bg-green-500={c.state?.running}
      class:bg-yellow-500={c.state?.waiting}
      class:bg-gray-500={c.state?.terminated}
      class:animate-pulse={c.state?.running && !c.ready}
    >
      <div class="tooltip">
        <div class="font-bold">{c.name}</div>
        <div class="text-xs leading-loose">
          {#if c.state?.running}
            <div>
              Running
              {#if c.ready}
                (Ready)
              {/if}
            </div>
            <div>Started at: {c.state?.running?.startedAt}</div>
          {/if}
          {#if c.state?.waiting}
            <div>Waiting ({c.state.waiting.reason})</div>
          {/if}
          {#if c.state?.terminated}
            <div>Terminated ({c.state.terminated.reason})</div>
            <div>Exit code: {c.state.terminated.exitCode}</div>
            <div>Started at: {c.state.terminated.startedAt}</div>
            <div>Finished at: {c.state.terminated.finishedAt}</div>
          {/if}
        </div>
      </div>
    </div>
  {/each}
</div>
