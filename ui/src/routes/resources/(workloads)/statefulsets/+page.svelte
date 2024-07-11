<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte'
  import { createStore, type Columns } from '$lib/stores/resources/workloads/statefulsets'

  const columns: Columns = [['name', 'emphasize'], ['namespace'], ['ready'], ['service'], ['age']]
</script>

<Datatable {columns} {createStore} />
