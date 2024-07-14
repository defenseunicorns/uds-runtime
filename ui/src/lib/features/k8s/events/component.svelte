<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2023-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['namespace'],
    ['age'],
    ['type'],
    ['reason'],
    ['object_kind'],
    ['object_name'],
    ['count'],
  ]
</script>

<DataTable {columns} {createStore} />
