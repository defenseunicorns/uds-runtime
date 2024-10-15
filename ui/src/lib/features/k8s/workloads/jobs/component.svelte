<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize w-4/12 truncate'],
    ['namespace', 'w-3/12 truncate'],
    ['completions', 'w-2/12 truncate'],
    ['durations', 'w-2/12 truncate'],
    ['age', 'w-1/12 truncate'],
  ]
  const name = 'Jobs'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
