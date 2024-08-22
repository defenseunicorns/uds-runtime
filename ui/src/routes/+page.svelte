<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { ClusterOverview } from '$features/k8s'
  import { authenticated } from '$lib/features/api-auth/store'
  import { apiAuthEnabled } from '$lib/features/api-auth/store'
  import { updateApiAuthEnabled } from '$lib/utils/http'
  import { onMount } from 'svelte'

  onMount(async () => {
    updateApiAuthEnabled()
  })

  // Redirect to /auth if api auth is enabled and user is not authenticated
  $: if ($apiAuthEnabled && !$authenticated) {
    window.location.href = '/auth'
  }
</script>

<!-- Hide homepage if api auth is enabled and user is not authenticated-->
{#if !$apiAuthEnabled || ($apiAuthEnabled && $authenticated)}
  <ClusterOverview />
{/if}
