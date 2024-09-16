<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { type PVCStatus } from '$features/k8s/types'
  import { getColorAndStatus } from '$lib/features/k8s/helpers'

  export let status: PVCStatus

  $: statusClass = getColorAndStatus('PersistentVolumeClaims', status)
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
