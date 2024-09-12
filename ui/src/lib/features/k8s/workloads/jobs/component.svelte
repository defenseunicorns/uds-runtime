<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [['name', 'emphasize'], ['namespace'], ['completions'], ['durations'], ['age']]
  const name = 'Jobs'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
