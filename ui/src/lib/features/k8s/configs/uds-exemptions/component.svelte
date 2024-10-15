<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['namespace', 'w-3/12'],
    ['details', 'w-1/12'],
    ['matcher', 'w-2/12'],
    ['policies', 'w-3/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'Exemptions'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
