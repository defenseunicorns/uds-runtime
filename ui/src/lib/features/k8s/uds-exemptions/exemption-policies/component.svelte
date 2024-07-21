<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { Policy } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

  export let policies: Policy[]
</script>

<ul class="mt-4 text-sm">
  {#each policies as policy}
    <li>- {policy}</li>
  {/each}
</ul>
