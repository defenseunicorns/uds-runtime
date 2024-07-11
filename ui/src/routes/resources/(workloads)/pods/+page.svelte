<script lang="ts">
  import Datatable from '$components/data/Datatable.svelte'
  import { createStore, type Columns } from '$stores/resources/workloads/pods'

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
