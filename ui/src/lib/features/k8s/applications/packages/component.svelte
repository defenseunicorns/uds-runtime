<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'emphasize w-1/12 truncate'],
    ['version', ' w-1/12 truncate'],
    ['description', 'italic  w-5/12 truncate'],
    ['arch', ' w-1/12 truncate'],
    ['flavor', ' w-1/12 truncate'],
    ['components', ' w-2/12 truncate'],
    ['age', ' w-1/12 truncate'],
  ]
  const name = 'Application Packages'
  const description = 'Packages are either UDS or Zarf packages that are deployed in the cluster.'
</script>

<DataTable {columns} {createStore} {name} {description} isNamespaced={false} disableRowClick={true} />
