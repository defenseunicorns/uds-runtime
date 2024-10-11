<script lang="ts">
  import { goto } from '$app/navigation'
  import { ChevronRight, Cube, Information } from 'carbon-icons-svelte'

  const services: Record<string, string>[] = [
    {
      name: 'Authorization',
      status: 'Running',
      path: '',
    },
    {
      name: 'Monitoring',
      status: 'Running',
      path: '',
    },
    {
      name: 'Service Mesh',
      status: 'Running',
      path: '',
    },
    {
      name: 'Healthy Access Management',
      status: 'Running',
      path: '',
    },
    {
      name: 'Log Aggregation',
      status: 'Running',
      path: '',
    },
    {
      name: 'Metrics',
      status: 'Running',
      path: '',
    },
    {
      name: 'Container Security',
      status: 'Running',
      path: '',
    },
    {
      name: 'Policy Engine & Operator',
      status: 'Running',
      path: '',
    },
    {
      name: 'Backup & Restore',
      status: 'Running',
      path: '',
    },
  ]
</script>

<div class="core-services">
  <div class="core-services__header">
    <h2 class="core-services__header-title">Core Services</h2>

    <Information class="ml-2 w-4 h-4 dark:text-gray-400 text-blue-500" />
  </div>

  <div class="core-services__rows">
    {#each services as { name, status, path }}
      <div class="core-services__rows-item">
        <div class="w-7/12 flex items-center space-x-2">
          <span class="core-services__name-icon">
            <Cube size={16} class="text-gray-400" />
          </span>

          <span>{name}</span>
        </div>

        <div class="w-2/12">{status}</div>

        <div class="w-3/12 flex justify-end">
          <button class="core-services__button-link" on:click={() => goto(path)}>
            <span>View Service</span>
            <ChevronRight />
          </button>
        </div>
      </div>
    {/each}
  </div>
</div>

<style lang="postcss">
  .core-services {
    @apply px-6 pb-6 w-full rounded-md dark:bg-gray-800 flex flex-col;
  }

  .core-services__header {
    @apply flex items-center h-14;
  }

  .core-services__header-title {
    @apply text-lg dark:text-white font-semibold;
  }

  .core-services__rows {
    @apply flex flex-col text-sm;
  }

  .core-services__rows-item {
    @apply flex dark:border-gray-700 items-center;
    border-bottom-width: 1px;
  }

  .core-services__name-icon {
    @apply h-7 w-7 rounded-lg bg-gray-700 flex items-center justify-center my-2;
  }

  .core-services__button-link {
    @apply text-sm text-blue-500 dark:text-blue-400 flex items-center space-x-1;
  }
</style>
