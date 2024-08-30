<script lang="ts">
  import { type PodStatus } from '$features/k8s/types'
  import { getColorAndStatus } from '$lib/features/k8s/helpers'
  export let status: PodStatus

  $: statusClass = getColorAndStatus('Pod', status)
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
