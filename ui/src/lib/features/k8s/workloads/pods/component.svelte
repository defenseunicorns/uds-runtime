<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-3/12'],
    ['namespace', 'w-1/12'],
    ['controlled_by', 'w-2/12'],
    ['containers', 'w-1/12'],
    ['status', 'w-1/12'],
    ['restarts', 'w-1/12'],
    ['usage', 'w-1/12'],
    ['node', 'w-1/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'Pods'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
