<script>
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  let name = 'CustomResourceDefinitions'
  const description = resourceDescriptions[name]

  let columns: Columns = [['name', 'emphasize'], ['group'], ['kind'], ['versions'], ['scope'], ['replicas'], ['age']]
</script>

<DataTable {columns} {createStore} {name} {description} />
