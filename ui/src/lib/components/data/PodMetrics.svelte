<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { ContainerMetric } from '@kubernetes/client-node'
  import { formatCPU, formatMemory, parseCPU, parseMemory } from './PodMetrics'

  export let containers: ContainerMetric[] = []

  $: totalCPU = containers.reduce((sum, container) => sum + parseCPU(container.usage.cpu), 0)
  $: totalMemory = containers.reduce((sum, container) => sum + parseMemory(container.usage.memory), 0)
</script>

{#if containers.length > 0}
  <div class="relative group">
    <div class="text-xs leading-loose text-nowrap">
      <div>CPU: {formatCPU(totalCPU)}</div>
      <div>Mem: {formatMemory(totalMemory)}</div>
    </div>

    <div class="tooltip">
      {#each containers as container}
        <div class="mb-1 leading-loose">
          <h3 class="font-bold">{container.name}:</h3>
          CPU: {formatCPU(parseCPU(container.usage.cpu))}<br />
          Mem: {formatMemory(parseMemory(container.usage.memory))}
          <br />
        </div>
      {/each}
    </div>
  </div>
{/if}
