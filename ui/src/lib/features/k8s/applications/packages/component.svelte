<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { DataTable } from '$components'

  import { createStore, type Columns } from './store'

  export let columns: Columns = [
    ['name', 'w-1/12'],
    ['version', ' w-1/12'],
    ['description', 'italic  w-5/12'],
    ['arch', ' w-1/12'],
    ['flavor', ' w-1/12'],
    ['components', ' w-2/12'],
    ['age', ' w-1/12'],
  ]
  const name = 'Application Packages'
  const description = 'Packages are either UDS or Zarf packages that are deployed in the cluster.'
</script>

<DataTable {columns} {createStore} {name} {description} isNamespaced={false} disableRowClick={true} />
