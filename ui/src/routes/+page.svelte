<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import Unauthenticated from '$components/Auth/component.svelte'
  import { authenticated } from '$features/api-auth/store'
  import { ClusterOverview } from '$features/k8s'
</script>

<!-- Hide homepage if api auth is enabled and user is not authenticated-->
{#if $authenticated}
  <ClusterOverview />
{:else}
  <Unauthenticated />
{/if}
