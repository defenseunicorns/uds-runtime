<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

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
