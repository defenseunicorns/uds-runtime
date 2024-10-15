<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-2/12'],
    ['namespace', 'w-2/12'],
    ['metrics', 'w-2/12'],
    ['min_pods', 'w-1/12'],
    ['max_pods', 'w-1/12'],
    ['replicas', 'w-1/12'],
    ['age', 'w-1/12'],
    ['status', 'w-2/12'],
  ]

  const name = 'HorizontalPodAutoscalers'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
