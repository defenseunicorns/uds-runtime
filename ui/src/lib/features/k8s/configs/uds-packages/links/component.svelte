<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { Link } from '$components'

  export let endpoints: string[]
</script>

{#if endpoints.length < 1}
  <p class="text-gray-500 dark:text-gray-400">-</p>
{/if}

{#each endpoints as endpoint}
  <Link href={`https://${endpoint}`} text={endpoint} />
{/each}
