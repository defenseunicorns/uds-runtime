<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['namespace', 'w-1/12'],
    ['status', 'w-1/12'],
    ['endpoints', 'w-2/12'],
    ['monitors', 'w-2/12'],
    ['ssoClients', 'w-1/12'],
    ['networkPolicies', 'w-2/12'],
    ['age', 'w-1/12'],
  ]

  const name = 'Packages'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
