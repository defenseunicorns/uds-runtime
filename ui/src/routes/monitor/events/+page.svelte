<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte'
  import { createStore, type Columns } from '$lib/stores/resources/events'

  export let columns: Columns = [
    ['namespace'],
    ['age'],
    ['type'],
    ['reason'],
    ['object_kind'],
    ['object_name'],
    ['count'],
  ]
</script>

<Datatable {columns} {createStore} />
