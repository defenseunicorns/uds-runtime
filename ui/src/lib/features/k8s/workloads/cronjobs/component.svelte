<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['namespace', 'w-2/12'],
    ['schedule', 'w-2/12'],
    ['suspend', 'w-2/12'],
    ['last_scheduled', 'w-2/12'],
    ['age', 'w-2/12'],
  ]
  const name = 'CronJobs'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
