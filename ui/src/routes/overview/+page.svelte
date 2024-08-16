<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script>
  import { goto } from '$app/navigation'
  import { onMount } from 'svelte'

  onMount(() => {
    // redirect to root because that's where the overview page is
    goto('/')
  })
</script>
