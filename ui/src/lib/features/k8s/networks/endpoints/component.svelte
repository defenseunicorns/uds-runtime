<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize w-3/12'],
    ['namespace', 'w-2/12'],
    ['endpoints', 'w-6/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'Endpoints'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
