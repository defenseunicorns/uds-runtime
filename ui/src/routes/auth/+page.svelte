<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import Unauthenticated from '$components/Auth/component.svelte'
  import { authenticated } from '$features/api-auth/store'

  if (authenticated) {
    goto('/')
  }
</script>

{#if !$authenticated}
  <Unauthenticated />
{/if}
