<script lang="ts">
  export let endpoints: string[]
</script>

{#each endpoints as endpoint}
  <a
    href={`https://${endpoint}`}
    target="_blank"
    rel="noopener noreferrer"
    on:click|stopPropagation
    class="font-medium text-blue-600 dark:text-blue-500 hover:underline pr-4">{endpoint}</a
  >
{/each}
