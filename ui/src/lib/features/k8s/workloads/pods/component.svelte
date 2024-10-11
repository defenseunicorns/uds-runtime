<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize w-3/12 truncate'],
    ['namespace', 'w-1/12 truncate'],
    ['controlled_by', 'w-2/12 truncate'],
    ['containers', 'w-1/12 truncate'],
    ['status', 'w-1/12 truncate'],
    ['restarts', 'w-1/12 truncate'],
    ['usage', 'w-1/12 truncate'],
    ['node', 'w-1/12 truncate'],
    ['age', 'w-1/12 truncate'],
  ]
  const name = 'Pods'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
