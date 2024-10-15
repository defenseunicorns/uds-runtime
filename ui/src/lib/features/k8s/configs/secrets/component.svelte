<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-3/12'],
    ['namespace', 'w-2/12'],
    ['type', 'w-2/12'],
    ['keys', 'line-clamp-5 max-w-screen-md w-4/12'],
    ['age', 'w-1/12'],
  ]

  const name = 'Secrets'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
