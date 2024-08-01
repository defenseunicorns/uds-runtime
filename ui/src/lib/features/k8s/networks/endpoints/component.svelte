<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [['name', 'emphasize'], ['namespace'], ['endpoints'], ['age']]
  const name = 'Endpoints'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
