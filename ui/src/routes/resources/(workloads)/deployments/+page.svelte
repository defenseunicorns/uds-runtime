<script lang="ts">
  import Datatable from '$lib/components/data/Datatable.svelte';
  import { createStore, type Columns } from '$lib/stores/resources/workloads/deployments';

  const columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['ready'],
    ['up_to_date'],
    ['available'],
    ['age']
  ];
</script>

<Datatable {columns} {createStore} />
