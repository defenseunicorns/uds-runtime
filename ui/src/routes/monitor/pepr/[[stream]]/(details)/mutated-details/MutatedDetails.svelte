<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { PeprDetails } from '$lib/types'

  export let details: PeprDetails
</script>

<div class="whitespace-nowrap relative group py-4 flex">
  <span class="text-blue-400">Details</span>
  <div class="tooltip tooltip-right">
    {#each Object.entries(details.operations || {}) as [key, ops], index}
      {#if index > 0}
        <br />
      {/if}
      <h3 class="font-bold">{key}:</h3>
      {#each ops as op}
        <p class="text-pretty tooltip-w">
          {op.path}
          {#if op.value}
            <span class="text-blue-400">={JSON.stringify(op.value)}</span>
          {/if}
        </p>
      {/each}
    {/each}
  </div>
</div>

<style>
  .tooltip-w {
    width: 496px;
  }
</style>
