<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte'
  import { createStore, type Columns } from '$lib/stores/resources/workloads/pods'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['controller'],
    ['containers'],
    ['status'],
    ['restarts'],
    ['metrics'],
    ['node'],
    ['age'],
  ]
</script>

<Datatable {columns} {createStore} />
