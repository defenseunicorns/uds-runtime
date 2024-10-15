<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

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
