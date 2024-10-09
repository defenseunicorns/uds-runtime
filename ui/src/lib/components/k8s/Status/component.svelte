<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { getColorForStatus } from '$lib/features/k8s/helpers'

  export let type
  export let status

  $: statusClass = getColorForStatus(type, status)
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
