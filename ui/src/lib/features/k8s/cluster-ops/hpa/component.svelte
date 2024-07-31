<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import type { NameAndDesc } from '$features/k8s/types'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['metrics'],
    ['min_pods'],
    ['max_pods'],
    ['replicas'],
    ['age'],
    ['status'],
  ]

  const resourceName = 'HorizontalPodAutoscalers'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
