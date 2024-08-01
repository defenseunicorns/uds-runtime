<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { KubernetesObject } from '@kubernetes/client-node'

  import { page } from '$app/stores'
  import { DataTable } from '$components'
  import type { ResourceStoreInterface } from '$features/k8s/types'
  import { type Columns, type Row } from './store'
  import { resourceDescriptions } from '$lib/utils/descriptions'

  export let columns: Columns = [['name', 'emphasize'], ['status'], ['age']]

  const { namespaces } = $page.data

  const createStore = (): ResourceStoreInterface<KubernetesObject, Row> => {
    return namespaces
  }

  const name = 'Namespaces'
  const description = resourceDescriptions[name]
</script>

{#if $namespaces}
  <DataTable {columns} {createStore} isNamespaced={false} {name} {description} />
{/if}
