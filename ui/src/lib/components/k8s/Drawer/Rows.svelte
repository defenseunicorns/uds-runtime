<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script context="module" lang="ts">
  export type VariantType = 'text' | 'icon-text' | 'external-link' | 'internal-link' | 'input-copy' | 'badges'

  export const fieldNameMap: Map<string, string> = new Map([
    ['creationTimestamp', 'Created'],
    ['namespace', 'Namespace'],
    ['controlled_by', 'Controlled'],
    ['node', 'Node'],
    ['labels', 'Labels'],
    ['annotations', 'Annotations'],
  ])
</script>

<script lang="ts">
  import Row from './Row.svelte'
  export let data: any
</script>

{#each Object.entries(data) as [key, val]}
  {#if fieldNameMap.has(key)}
    {#if typeof val === 'string'}
      <div class="mb-7">
        <Row label={key} data={{ variant: 'text', value: val }} />
      </div>
    {:else}
      <Row label={key} data={{ variant: 'badges', value: val }} />
    {/if}
  {/if}
{/each}
