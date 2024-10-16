<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['namespace', 'w-2/12'],
    ['age', 'w-1/12'],
    ['type', 'w-1/12'],
    ['reason', 'w-2/12'],
    ['object_kind', 'w-2/12'],
    ['object_name', 'w-3/12'],
    ['count', 'w-1/12'],
  ]

  const name = 'Events'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
