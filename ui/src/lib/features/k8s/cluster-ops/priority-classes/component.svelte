<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-4/12'],
    ['value', 'w-1/12'],
    ['global_default', 'w-2/12'],
    ['description', 'w-4/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'PriorityClasses'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} isNamespaced={false} />
