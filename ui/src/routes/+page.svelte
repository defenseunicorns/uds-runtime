<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import { ClusterOverview } from '$features/k8s'
  import { apiAuthEnabled, authenticated } from '$lib/features/api-auth/store'
  import { updateApiAuthEnabled } from '$lib/utils/helpers'
  import { onMount } from 'svelte'

  onMount(async () => {
    updateApiAuthEnabled()
  })

  // Redirect to /auth if api auth is enabled and user is not authenticated
  $: if ($apiAuthEnabled && !$authenticated) {
    goto('/auth')
  }
</script>

<!-- Hide homepage if api auth is enabled and user is not authenticated-->
{#if !$apiAuthEnabled || ($apiAuthEnabled && $authenticated)}
  <ClusterOverview />
{/if}
