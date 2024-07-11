<script lang="ts">
  import { page } from '$app/stores'
  import Datatable from '$lib/components/data/Datatable.svelte'
  import type { ResourceStoreInterface } from '$lib/stores/resources/common'
  import { type Columns } from '$lib/stores/resources/namespaces'
  import type { KubernetesObject } from '@kubernetes/client-node'

  export let columns: Columns = [['name', 'emphasize'], ['status'], ['age']]

  const { namespaces } = $page.data

  const createStore = (): ResourceStoreInterface<KubernetesObject, any> => {
    return namespaces
  }
</script>

{#if $namespaces}
  <Datatable {columns} {createStore} />
{/if}
