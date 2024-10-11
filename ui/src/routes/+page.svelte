<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import Unauthenticated from '$components/Auth/component.svelte'
  import { authenticated } from '$features/auth/store'
  import { ClusterOverview } from '$features/k8s'

  // remove token param from URL if it exists
  if (window.location.search.includes('token')) {
    window.history.replaceState({}, '', window.location.pathname)
  }
</script>

<!-- Hide homepage if api auth is enabled and user is not authenticated-->
{#if $authenticated}
  <ClusterOverview />
{:else}
  <Unauthenticated />
{/if}
