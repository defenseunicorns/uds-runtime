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
    <div class="w-full">
      <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">{statText}</dt>
      <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
        {value.toString()}%
      </dd>
    </div>

    <ProgressBar size={barSize} {progress} {capacity} {unit} />
  </div>
</Card>
