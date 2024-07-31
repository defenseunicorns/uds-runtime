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
    ['desired'],
    ['current'],
    ['ready'],
    ['up_to_date'],
    ['available'],
    ['node_selector'],
    ['age'],
  ]
  const resourceName = 'DaemonSets'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
