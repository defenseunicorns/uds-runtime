<script lang="ts">
  import Datatable from '$components/data/Datatable.svelte'
  import { createStore, type Columns } from '$stores/resources/events'

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
