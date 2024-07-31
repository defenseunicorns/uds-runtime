<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import type { Resource } from '$features/k8s/types'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['schedule'],
    ['suspend'],
    ['last_scheduled'],
    ['age'],
  ]
  const resourceName = 'CronJobs'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
