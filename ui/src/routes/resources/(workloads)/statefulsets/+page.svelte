<script lang="ts">
  import Datatable from '$components/data/Datatable.svelte'
  import { createStore, type Columns } from '$stores/resources/workloads/statefulsets'

  const columns: Columns = [['name', 'emphasize'], ['namespace'], ['ready'], ['service'], ['age']]
</script>

<Datatable {columns} {createStore} />
