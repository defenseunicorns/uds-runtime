<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script context="module" lang="ts">
  export type VariantType = 'text' | 'icon-text' | 'external-link' | 'internal-link' | 'input-copy' | 'badges'
</script>

<script lang="ts">
  import Row from './Row.svelte'
  export let data: any
  console.log('data')
  console.log(data)
</script>

{#each Object.entries(data) as [key, val]}
  {#if typeof val === 'string'}
    <div class="mb-7">
      <Row label={key} data={{ variant: 'text', value: val }} />
    </div>
  {:else}
    <Row label={key} data={{ variant: 'badges', value: val }} />
  {/if}
{/each}
