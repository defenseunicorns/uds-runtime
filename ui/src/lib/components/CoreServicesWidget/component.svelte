<script lang="ts">
  import type { V1Pod } from '@kubernetes/client-node'
  import { type CoreServiceType } from '$lib/types'
  import { Cube, Information } from 'carbon-icons-svelte'

  export let coreServices: CoreServiceType[] = []
  export let pods: V1Pod[] = []

  const coreServicesMapping: Record<string, string> = {
    authservice: 'Authorization',
    grafana: 'Monitoring',
    istio: 'Service Mesh',
    keycloak: 'Identity Access Management',
    loki: 'Log Aggregation',
    'metrics-server': 'Metrics',
    neuvector: 'Container Security',
    'prometheus-stack': 'Monitoring',
    vector: 'Log Aggregation',
    velero: 'Backup & Restore',
    'uds-runtime': 'Frontend Views & Insights',
  }

  const coreServiceKeys = Object.keys(coreServicesMapping)

  let hasNoCoreServices: boolean = false
  let uniqueServiceList: string[] = []
  let hasPolicyEngineOperator: boolean = false

  $: hasNoCoreServices = coreServices.every((service) => !coreServiceKeys.includes(service.metadata.name))
  $: {
    hasPolicyEngineOperator = pods.filter((pod: V1Pod) => pod?.metadata?.name?.match(/^pepr-uds-core/)).length > 0

    coreServices.forEach((service) => {
      let serviceName = coreServicesMapping[service.metadata.name]

      uniqueServiceList.push(serviceName)
      uniqueServiceList = Array.from(new Set([...uniqueServiceList]))

      if (hasPolicyEngineOperator && !uniqueServiceList.includes('Policy Engine & Operator')) {
        uniqueServiceList.push('Policy Engine & Operator')
      }
    })
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
    {@const sortedServices = uniqueServiceList.sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0))}
    <div class="core-services__rows">
      {#each sortedServices as serviceName}
        <div class="core-services__rows-item">
          <div class="w-10/12 flex items-center space-x-2">
            <div class="core-services__name-icon">
              <Cube size={16} class="text-gray-400" />
            </div>

            <div class="truncate">{serviceName}</div>
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
</style>
