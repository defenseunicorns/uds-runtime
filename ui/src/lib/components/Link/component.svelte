<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  export let href: string
  export let target: string = '_blank'
  export let text: string
</script>

{#if href !== ''}
  <a
    {href}
    {target}
    on:click|stopPropagation
    rel="noopener noreferrer"
    class="font-medium text-blue-600 dark:text-blue-500 hover:underline pr-4"
  >
    {text}
  </a>
{:else}
  <span class="dark:text-gray-400 mt-4 text-sm">{text}</span>
{/if}
