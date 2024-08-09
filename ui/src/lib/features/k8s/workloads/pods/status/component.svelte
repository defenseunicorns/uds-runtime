<script lang="ts">
  export let status: string

  const statusClasses = {
    Running: 'text-green-400',
    Stopped: 'text-red-400',
    Pending: 'text-orange-300',
  }

  $: statusClass = statusClasses[status as keyof typeof statusClasses]
</script>

{#if status}
  <span class={statusClass}>
    {status}
  </span>
{:else}
  -
{/if}
