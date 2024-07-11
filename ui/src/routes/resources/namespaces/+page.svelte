<script lang="ts">
  import type { KubernetesObject } from '@kubernetes/client-node'
  
  import { page } from '$app/stores'
  import Datatable from '$components/data/Datatable.svelte'
  import type { ResourceStoreInterface } from '$stores/resources/common'
  import { type Columns } from '$stores/resources/namespaces'

  export let columns: Columns = [['name', 'emphasize'], ['status'], ['age']]

  const { namespaces } = $page.data

  const createStore = (): ResourceStoreInterface<KubernetesObject, any> => {
    return namespaces
  }
</script>

{#if $namespaces}
  <Datatable {columns} {createStore} />
{/if}
