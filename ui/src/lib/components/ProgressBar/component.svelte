<script lang="ts">
  type BarSizeType = 'sm' | 'md' | 'lg' | 'xl'
  type UnitType = 'Cores' | 'GB'

  export let size: BarSizeType = 'sm'
  export let progress: number = 0
  export let capacity: number = 0
  export let unit: UnitType = 'GB'

  let calculatedWidth = 0
  let fixedProgress: string
  let fixedCapacity: string

  $: {
    let percentage = (progress / capacity) * 100
    fixedProgress = progress.toFixed(2)
    fixedCapacity = capacity.toFixed(2)
    // Need a minimum width set for rounded corners to look right
    calculatedWidth = percentage < 2 ? 2 : percentage
  }

  $: progressText = `${fixedProgress} ${unit} of ${fixedCapacity} ${unit} used`

  const sizeMapping = {
    sm: 'h-1.5',
    md: 'h-2.5',
    lg: 'h-4',
    xl: 'h-6',
  }
</script>

<div class="flex flex-col">
  <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700 mt-3">
    <div class={`bg-green-600 rounded-full ${sizeMapping[size]}`} style={`width: ${calculatedWidth}%`}></div>
  </div>

  <span class="text-xs mt-1 font-normal text-gray-500 dark:text-gray-400 truncate overflow-ellipsis">
    {progressText}
  </span>
</div>
