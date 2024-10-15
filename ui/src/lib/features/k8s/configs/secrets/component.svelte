<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['type'],
    ['keys', 'line-clamp-5 max-w-screen-md'],
    ['age'],
  ]

  const name = 'Secrets'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
