<script lang="ts">
  import { Db2Database } from 'carbon-icons-svelte'

  import { type VariantType } from './Row.svelte'

  export let label: string
  export let variant: VariantType
</script>

<div class="py-1 border-b border-gray-700 gap-9 flex">
  <div class="h-6 justify-start items-center gap-1.5 flex my-1 w-[180px]">
    {#if variant === 'icon-text'}
      <Db2Database color="white" />
    {/if}
    <div class="text-white text-base font-semibold leading-normal">{label}</div>
  </div>
  <div class="justify-start items-center gap-1.5 flex">
    <slot />
  </div>
</div>
