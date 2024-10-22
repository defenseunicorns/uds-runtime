<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { Card, ProgressBar } from '$components'
  import { Information } from 'carbon-icons-svelte'

  import type { BarSizeType, UnitType } from './types.ts'

  export let capacity: number = 0
  export let progress: number = 0
  export let barSize: BarSizeType = 'sm'
  export let statText: string
  export let unit: UnitType
  export let value: number | string
  export let deactivated: boolean = false
</script>

<Card>
  <div class="w-full">
    <div class="w-full">
      {#if deactivated}
        <div class="flex justify-end">
          <span
            class="bg-gray-100 text-gray-500 text-xs font-medium inline-flex items-center px-2.5 py-0.5 rounded me-2 dark:bg-gray-100 dark:text-gray-500 border border-gray-500"
            data-testid="unavailable-tag"
          >
            <div class="relative group mr-2">
              <Information class="w-4 h-4 text-grey-500" />
              <div class="tooltip tooltip-left min-w-56">
                <div class="whitespace-normal">
                  Metrics Server is unavailable.
                  <br />
                  Ensure Metrics Server is running in the cluster.
                </div>
              </div>
            </div>
            Unavailable
          </span>
        </div>
      {/if}
      <div class="flex items-center">
        <dt class="text-sm font-medium text-gray-500 dark:text-gray-500 truncate">{statText}</dt>
      </div>
      <dd
        class="mt-1 text-3xl font-semibold"
        class:text-gray-900={!deactivated}
        class:text-gray-500={deactivated}
        class:dark:text-white={!deactivated}
      >
        {value.toString()}
      </dd>
    </div>

    {#if !deactivated}
      <ProgressBar size={barSize} {progress} {capacity} {unit} />
    {/if}
  </div>
</Card>
