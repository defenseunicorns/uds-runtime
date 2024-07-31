<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'
  import type { NameAndDesc } from '$features/k8s/types'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['type'],
    ['keys', 'line-clamp-5 max-w-screen-md'],
    ['age'],
  ]

  const resourceName = 'Secrets'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
