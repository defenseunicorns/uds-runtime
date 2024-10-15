<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize w-4/12 truncate'],
    ['namespace', 'w-4/12 truncate'],
    ['ready', 'w-1/12 truncate'],
    ['up_to_date', 'w-1/12 truncate'],
    ['available', 'w-1/12 truncate'],
    ['age', 'w-1/12 truncate'],
  ]
  const name = 'Deployments'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
