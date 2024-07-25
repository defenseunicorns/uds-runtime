<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { type Columns, createStore } from './store'

  export let columns: Columns = [['name', 'emphasize'], ['storage_class'], ['capacity'], ['claim'], ['age'], ['status']]
</script>

<DataTable {columns} {createStore} />
