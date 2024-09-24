<script lang="ts">
  import IconWidget from './IconWidget.svelte'
  import ProgressBarWidget from './ProgressBarWidget.svelte'
  import type { StatType } from './types'

  export let variant: 'with_right_icon' | 'progress_bar'
  export let props: StatType
</script>

{#if variant === 'progress_bar' && 'capacity' in props}
  <svelte:component this={ProgressBarWidget} {...props} />
{:else if variant === 'with_right_icon' && 'icon' in props}
  <svelte:component this={IconWidget} {...props} />
{/if}
