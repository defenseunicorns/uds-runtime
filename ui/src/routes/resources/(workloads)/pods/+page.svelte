<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte'
  import { createStore } from '$lib/features/pods/pods'
  import { type Columns } from '$lib/features/pods/types'

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
