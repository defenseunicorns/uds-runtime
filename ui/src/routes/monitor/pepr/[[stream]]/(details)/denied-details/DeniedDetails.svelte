<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { PeprDetails } from '$lib/types'

  export let details: PeprDetails
</script>

<span class="whitespace-nowrap relative group py-4 flex items-center">
  <span class="text-blue-400">Details</span>
  <div class="tooltip tooltip-right">
    {#each details.messages || [] as msg}
      <div class="text-red-400">{msg}</div>
    {/each}
  </div>
</span>
