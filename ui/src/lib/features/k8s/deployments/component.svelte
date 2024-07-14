<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'

  const columns: Columns = [['name', 'emphasize'], ['namespace'], ['ready'], ['up_to_date'], ['available'], ['age']]
</script>

<DataTable {columns} {createStore} />
