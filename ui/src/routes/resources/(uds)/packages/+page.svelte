<script lang="ts">
  import Datatable from '$components/data/Datatable.svelte'
  import { createStore, type Columns } from '$stores/resources/uds/packages'

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
