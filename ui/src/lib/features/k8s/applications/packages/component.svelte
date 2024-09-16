<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize'],
    ['version'],
    ['description', 'italic'],
    ['arch'],
    ['flavor'],
    ['components'],
    ['age'],
  ]
  const name = 'Application Packages'
  const description = 'Packages are either UDS or Zarf packages that are deployed in the cluster.'
</script>

<DataTable {columns} {createStore} {name} {description} isNamespaced={false} disableRowClick={true} />
