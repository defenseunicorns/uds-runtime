<script lang="ts">
  import type { TailwindSizeType } from '$components/StatsWidget/types'

  export let size: TailwindSizeType = 32
</script>

<div
  class={`bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg hover:dark:bg-gray-700 h-${size} px-5 sm:px-6 flex items-center`}
>
  <slot />
</div>
