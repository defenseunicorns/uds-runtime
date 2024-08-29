<script lang="ts">
  import { getColorAndStatus } from '$lib/features/k8s/helpers.ts'
  export let status: string

  $: statusClass = getColorAndStatus('Pod', status)
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
