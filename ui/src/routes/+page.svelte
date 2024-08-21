<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'
  import { ClusterOverview } from '$features/k8s'
  import { authenticated } from '$lib/features/api-auth/store'

  let isAuthenticated: boolean
  const apiAuthEnabled = import.meta.env.VITE_API_AUTH?.toLowerCase() === 'true'

  // If API_AUTH is enabled, subscribe to the authenticated store to check if the user is authenticated
  if (apiAuthEnabled) {
    authenticated.subscribe((value) => {
      isAuthenticated = value
    })
  }

  // Redirect to /auth if api auth is enabled and user is not authenticated
  onMount(() => {
    if (apiAuthEnabled && !isAuthenticated) {
      window.location.href = '/auth'
    }
  })
</script>

<!-- Hide homepage if api auth is enabled and user is not authenticated-->
{#if !apiAuthEnabled || (apiAuthEnabled && $authenticated)}
  <ClusterOverview />
{/if}
