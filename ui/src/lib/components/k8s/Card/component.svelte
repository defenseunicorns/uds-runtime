<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import type { TailwindSizeType } from '$components/StatsWidget/types'

  export let size: TailwindSizeType = 32
  export let link: string = ''
</script>

<button
  on:click={() => goto(link)}
  data-testid="card-container"
  class={`bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg hover:dark:bg-gray-700 h-${size} px-5 sm:px-6 flex items-center`}
>
  <slot />
</button>
