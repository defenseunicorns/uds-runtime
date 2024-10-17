<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { type PVCStatus } from '$features/k8s/types'
  import { getColorForStatus } from '$lib/features/k8s/helpers'

  export let status: PVCStatus

  $: statusClass = getColorForStatus('PersistentVolumeClaims', status)
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
