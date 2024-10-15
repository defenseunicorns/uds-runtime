<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['namespace', 'w-2/12'],
    ['type', 'w-2/12'],
    ['cluster_ip', 'w-2/12'],
    ['external_ip', 'w-1/12'],
    ['ports', 'w-2/12'],
    ['age', 'w-1/12'],
  ]

  const name = 'Services'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
