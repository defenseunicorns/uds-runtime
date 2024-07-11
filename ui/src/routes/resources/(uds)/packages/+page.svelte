<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte'
  import { createStore, type Columns } from '$lib/stores/resources/uds/packages'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['status'],
    ['endpoints'],
    ['monitors'],
    ['ssoClients'],
    ['networkPolicies'],
    ['age'],
  ]
</script>

<Datatable {columns} {createStore} />
