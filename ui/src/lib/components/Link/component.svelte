<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  export let href: string
  export let target: string = '_blank'
  export let text: string
</script>

<a
  {href}
  {target}
  on:click|stopPropagation
  rel="noopener noreferrer"
  class="font-medium text-blue-600 dark:text-blue-500 hover:underline pr-4">{text}</a
>
