<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'
  import { createStore, type Columns } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['namespace'],
    ['ingress_ports'],
    ['ingress_block'],
    ['egress_ports'],
    ['egress_block'],
    ['age'],
  ]
  const name = 'NetworkPolicies'
  const description = resourceDescriptions[name]
</script>

<DataTable {columns} {createStore} {name} {description} />
