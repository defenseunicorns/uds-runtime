<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { type PatchOperation } from './helpers'

  export let details: Record<string, PatchOperation[]>
</script>

<span class="whitespace-nowrap relative group py-4 flex items-center">
  <span class="text-blue-400">&nbsp;Details</span>
  <div class="tooltip tooltip-right">
    {#each Object.entries(details) as [key, operations]}
      <h3 class="font-bold">{key}:</h3>
      {#each operations as op}
        <div class="text-pretty tooltip-w">{op.path}=<span class="text-blue-400">{op.value}</span></div>
      {/each}
    {/each}
  </div>
</span>

<style>
  .tooltip-w {
    width: 496px;
  }
</style>
