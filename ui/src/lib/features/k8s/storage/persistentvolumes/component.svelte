<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [['name', 'emphasize'], ['storage_class'], ['capacity'], ['claim'], ['age'], ['status']]
  const name = 'PersistentVolumes'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} isNamespaced={false} />
