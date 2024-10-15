<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['status', 'w-2/12'],
    ['roles', 'w-2/12'],
    ['taints', 'w-2/12'],
    ['version', 'w-2/12'],
    ['age', 'w-2/12'],
  ]
  const name = 'Nodes'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} isNamespaced={false} {name} {description} />
