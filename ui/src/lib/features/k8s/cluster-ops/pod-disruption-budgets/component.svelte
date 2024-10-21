<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-3/12'],
    ['namespace', 'w-1/12'],
    ['min_available', 'w-2/12'],
    ['max_unavailable', 'w-2/12'],
    ['current_healthy', 'w-2/12'],
    ['desired_healthy', 'w-2/12'],
    ['age', 'w-1/12'],
  ]
  const name = 'PodDisruptionBudgets'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
