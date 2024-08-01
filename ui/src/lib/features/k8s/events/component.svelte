<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['namespace'],
    ['age'],
    ['type'],
    ['reason'],
    ['object_kind'],
    ['object_name'],
    ['count'],
  ]

  const name = 'Events'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
