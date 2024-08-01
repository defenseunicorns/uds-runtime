<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['keys', 'line-clamp-3 max-w-screen-md'],
    ['age'],
  ]
  const name = 'ConfigMaps'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
