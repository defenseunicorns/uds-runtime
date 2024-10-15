<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['url', 'w-3/12', 'link-external'],
    ['namespace', 'w-3/12'],
    ['status', 'w-3/12'],
    ['age', 'w-3/12'],
  ]

  const name = 'Application Endpoints'
  const description = 'Application Endpoints are exposed by the UDS Operator via a UDS Package CR.'
</script>

<DataTable {columns} {createStore} {name} {description} disableRowClick={true} />
