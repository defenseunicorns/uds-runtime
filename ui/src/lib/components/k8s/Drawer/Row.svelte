<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script context="module" lang="ts">
  export type VariantType = 'text' | 'icon-text' | 'external-link' | 'internal-link' | 'input-copy' | 'badges'
</script>

<script lang="ts">
  import { CopyFile } from 'carbon-icons-svelte'

  import RowItem from './RowItem.svelte'

  type BadgesType = {
    [key: string]: string
  }

  export let label: string = 'text'

  type RowType = {} & (
    | { variant: 'text'; value: string }
    | { variant: 'icon-text'; value: string }
    | { variant: 'external-link'; value: string }
    | { variant: 'internal-link'; value: string }
    | { variant: 'input-copy'; value: string[] }
    | { variant: 'badges'; value: BadgesType }
  )

  export let data: RowType

  const hasValue = (value: string | object) => {
    if (value === undefined || value === null) {
      return false
    }

    if (typeof value === 'object') {
      return Object.entries(value).length > 0
    }

    return value
  }
</script>

{#if hasValue(data.value)}
  {#if data.variant === 'text' || data.variant === 'icon-text'}
    <RowItem {label} variant={data.variant}>
      <div class="text-gray-300 text-base font-extralight leading-normal">{data.value}</div>
    </RowItem>
  {:else if data.variant === 'external-link'}
    <RowItem {label} variant={data.variant}>
      <div class="text-base font-extralight text-blue-600 leading-normal underline">
        <a href={data.value}> Value string </a>
      </div>
    </RowItem>
  {:else if data.variant === 'internal-link'}
    <RowItem {label} variant={data.variant}>
      <div class="text-base font-extralight text-blue-600 leading-normal">
        <button on:click={() => console.log('clicked')}> Value string </button>
      </div>
    </RowItem>
  {:else if data.variant === 'input-copy'}
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
  {:else if data.variant === 'badges'}
    <RowItem {label} variant={data.variant}>
      {#each Object.entries(data.value) as [key, val]}
        <span
          class="bg-gray-100 text-gray-800 text-xs font-medium m-1 px-2.5 py-1 rounded dark:bg-gray-700 dark:text-gray-300"
        >
          {`${key}=${val}`}
        </span>
      {/each}
    </RowItem>
  {/if}
{/if}
