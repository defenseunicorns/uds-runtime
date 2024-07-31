<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { type PatchOperation } from './helpers'

  export let details: Record<string, PatchOperation[]>
</script>

<span class="whitespace-nowrap relative group py-4 flex items-center">
  <span class="text-blue-400">Details</span>
  <div class="tooltip tooltip-right">
    {#each details['messages'] as msg}
      <div class="text-pretty tooltip-w text-red-400">{msg}</div>
    {/each}
  </div>
</span>

<style>
  .tooltip-w {
    width: 496px;
  }
</style>
