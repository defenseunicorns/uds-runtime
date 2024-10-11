<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import Unauthenticated from '$components/Auth/component.svelte'
  import { authenticated } from '$features/auth/store'

  if (authenticated) {
    goto('/')
  }
</script>

{#if !$authenticated}
  <Unauthenticated />
{/if}
