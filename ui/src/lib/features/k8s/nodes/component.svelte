<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import type { NameAndDesc } from '$features/k8s/types'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [['name', 'emphasize'], ['status'], ['roles'], ['taints'], ['version'], ['age']]
  const resourceName = 'Nodes'
  const resource: NameAndDesc = {
    name: resourceName,
    desc: resourceDescriptions[resourceName],
  }
</script>

<DataTable {columns} {createStore} isNamespaced={false} {resource} />
