<script lang="ts">
  import { goto } from '$app/navigation'
  import { Card } from '$components'

  export let path: string
</script>

<Card>
  <button class="flex flex-col" on:click={() => goto(path)}>
    <slot />
  </button>
</Card>
