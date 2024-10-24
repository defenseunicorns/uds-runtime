<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { Card, ProgressBar } from '$components'

  import type { BarSizeType, UnitType } from './types.ts'

  export let capacity: number = 0
  export let progress: number = 0
  export let barSize: BarSizeType = 'sm'
  export let statText: string
  export let unit: UnitType
  export let value: number | string
</script>

<Card>
  <div class="w-full">
    <div class="flex flex-col items-start">
      <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">{statText}</dt>
      <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
        {value.toString()}%
      </dd>
    </div>

    <ProgressBar size={barSize} {progress} {capacity} {unit} />
  </div>
</Card>
