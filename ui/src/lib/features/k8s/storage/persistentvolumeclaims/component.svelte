<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import type { Resource } from '$features/k8s/types'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['storage_class'],
    ['capacity'],
    ['age'],
    ['status'],
  ]
  const resourceName = 'PersistentVolumeClaims'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
