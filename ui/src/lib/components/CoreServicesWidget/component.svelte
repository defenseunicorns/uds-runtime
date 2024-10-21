<script lang="ts">
  import { type CoreServiceType } from '$lib/types'
  import { Cube, Information } from 'carbon-icons-svelte'

  export let services: CoreServiceType[] = []
  export let sortedCoreServices: CoreServiceType[] = []

  const coreServicesMapping: Record<string, string> = {
    authservice: 'Authservice',
    grafana: 'Grafana',
    keycloak: 'KeyCloak',
    loki: 'Loki',
    'metrics-server': 'Metrics Server',
    neuvector: 'Neuvector',
    'prometheus-stack': 'Prometheus Stack',
    vector: 'Vector',
    velero: 'Velero',
    'uds-runtime': 'UDS Runtime',
  }

  const coreServiceKeys = Object.keys(coreServicesMapping)

  let hasNoCoreServices: boolean = false
  $: hasNoCoreServices = services.every((service) => !coreServiceKeys.includes(service.metadata.name))
  $: {
    sortedCoreServices = services
      .filter((service) => coreServiceKeys.includes(service.metadata.name))
      .sort((a, b) => a.metadata.name.charCodeAt(0) - b.metadata.name.charCodeAt(0))
  }
</script>

<div class="core-services">
  <div class="core-services__header">
    <h2 class="core-services__header-title">Core Services</h2>

    <Information class="ml-2 w-4 h-4 dark:text-gray-400 text-blue-500" />
  </div>

  {#if hasNoCoreServices}
    <span class="flex self-center">No Core Services running</span>
  {:else}
    <div class="core-services__rows">
      {#each sortedCoreServices as { metadata: { name } }}
        <div class="core-services__rows-item">
          <div class="w-10/12 flex items-center space-x-2">
            <div class="core-services__name-icon">
              <Cube size={16} class="text-gray-400" />
            </div>

            <div class="truncate">{coreServicesMapping[name]}</div>
          </div>

          <div class="w-2/12 text-green-400">Running</div>
        </div>
      {/each}
    </div>
  {/if}
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
