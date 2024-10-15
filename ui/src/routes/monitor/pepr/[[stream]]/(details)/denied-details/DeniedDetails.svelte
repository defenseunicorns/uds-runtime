<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import type { PeprDetails } from '$lib/types'

  export let details: PeprDetails
</script>

<span class="whitespace-nowrap relative group py-4 flex items-center">
  <span class="text-blue-400">Details</span>
  <div class="tooltip tooltip-left">
    {#each details.messages || [] as msg}
      <div class="text-red-400">{msg}</div>
    {/each}
  </div>
</span>
