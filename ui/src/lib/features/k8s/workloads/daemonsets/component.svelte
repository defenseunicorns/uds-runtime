<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-3/12'],
    ['namespace', 'w-2/12'],
    ['desired', 'w-1/12'],
    ['current', 'w-1/12'],
    ['ready', 'w-1/12'],
    ['up_to_date', 'w-1/12'],
    ['available', 'w-1/12'],
    ['node_selector', 'w-2/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'DaemonSets'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
