<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { Db2Database } from 'carbon-icons-svelte'

  import { type VariantType } from './Row.svelte'
  import { fieldNameMap } from './Rows.svelte'

  export let label: string
  export let variant: VariantType
</script>

<div class="py-1 border-b border-gray-700 gap-9 flex mt-3">
  <div class="flex-none w-1/4">
    {#if variant === 'icon-text'}
      <Db2Database color="white" />
    {/if}
    <div class="text-white text-base font-semibold leading-normal">
      {fieldNameMap.get(label) || label}
    </div>
  </div>

  <div class="flex w-3/4">
    <div class="flex flex-col">
      <div class="flex flex-wrap">
        <slot />
      </div>
    </div>
  </div>
</div>
