<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script>
  import { goto } from '$app/navigation'
  import { onMount } from 'svelte'

  onMount(() => {
    // redirect to overview page
    if (window.location.pathname === '/') {
      goto('/overview')
    }
  })
</script>
