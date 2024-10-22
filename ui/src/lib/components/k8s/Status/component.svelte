<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { getColorForStatus } from '$lib/features/k8s/helpers'

  export let type
  export let status

  $: statusClass = getColorForStatus(type, status)
</script>

{#if status}
  <span class={`${statusClass} truncate`}>
    {status}
  </span>
{:else}
  -
{/if}
