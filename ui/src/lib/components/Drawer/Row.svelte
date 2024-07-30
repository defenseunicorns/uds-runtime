<script context="module" lang="ts">
  export const enum Variants {
    TEXT = 'text',
    ICON_TEXT = 'icon-text',
    EXTERNAL_LINK = 'external-link',
    INTERNAL_LINK = 'internal-link',
    INPUT_COPY = 'input-copy',
    BADGES = 'badges',
  }
</script>

<script lang="ts">
  import { CopyFile } from 'carbon-icons-svelte'

  import RowItem from './RowItem.svelte'

  export let label: string = 'text'

  type Props = {} & (
    | { variant: Variants.TEXT; value: string }
    | { variant: Variants.ICON_TEXT; value: string }
    | { variant: Variants.EXTERNAL_LINK; value: string }
    | { variant: Variants.INTERNAL_LINK; value: string }
    | { variant: Variants.INPUT_COPY; value: string[] }
    | { variant: Variants.BADGES; value: string[] }
  )

  export let data: Props
</script>

{#if data.variant === Variants.TEXT || data.variant === Variants.ICON_TEXT}
  <RowItem {label} variant={data.variant}>
    <div class="text-gray-300 text-base font-extralight leading-normal">Value string</div>
  </RowItem>
{:else if data.variant === Variants.EXTERNAL_LINK}
  <RowItem {label} variant={data.variant}>
    <div class="text-base font-extralight text-blue-600 leading-normal underline">
      <a href={data.value}> Value string </a>
    </div>
  </RowItem>
{:else if data.variant === Variants.INTERNAL_LINK}
  <RowItem {label} variant={data.variant}>
    <div class="text-base font-extralight text-blue-600 leading-normal">
      <button on:click={() => console.log('clicked')}> Value string </button>
    </div>
  </RowItem>
{:else if data.variant === Variants.INPUT_COPY}
  <RowItem {label} variant={data.variant}>
    <div class="text-base leading-normal text-white flex flex-col">
      {#each data.value as path}
        <div class="w-[364px] justify-start items-start gap-1 mb-2">
          <div
            class="self-stretch px-4 py-1 bg-gray-700 rounded-md border border-gray-600 justify-start items-center gap-2.5 inline-flex mb-2"
          >
            <div class="grow shrink basis-0 h-[21px] justify-start items-center gap-2.5 flex w-[450px]">
              <div class="grow shrink basis-0 text-gray-400 text-sm font-normal leading-[21px]">
                {path}
              </div>

              <div class="w-2.5 h-3.5 relative bg-white/0 text-gray-400">
                <CopyFile />
              </div>
            </div>
          </div>

          <div class="self-stretch text-gray-400 text-xs font-normal leading-[15px]">from what resource</div>
        </div>
      {/each}
    </div>
  </RowItem>
{:else if data.variant === Variants.BADGES}
  <RowItem {label} variant={data.variant}>
    <div class="text-base leading-normal text-white flex">
      {#each data.value as badge}
        <span
          class="bg-gray-100 text-gray-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300"
        >
          {badge}
        </span>
      {/each}
    </div>
  </RowItem>
{/if}
