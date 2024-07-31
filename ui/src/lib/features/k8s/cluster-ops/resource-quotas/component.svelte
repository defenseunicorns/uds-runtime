<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'
  import type { Resource } from '$features/k8s/types'

  export let columns: Columns = [['name', 'emphasize'], ['namespace'], ['age']]
  const resourceName = 'ResourceQuotas'
  const resource: Resource = {
    name: resourceName,
    description: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} {resource} />
