<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['namespace', 'w-2/12 truncate'],
    ['age', 'w-1/12 truncate'],
    ['type', 'w-2/12 truncate'],
    ['reason', 'w-3/12 truncate'],
    ['object_kind', 'w-2/12 truncate'],
    ['object_name', 'w-3/12 truncate'],
    ['count', 'w-1/12 truncate'],
  ]

  const name = 'Events'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
