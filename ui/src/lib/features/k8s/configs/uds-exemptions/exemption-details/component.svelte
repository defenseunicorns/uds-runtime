<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { InformationFilled } from 'carbon-icons-svelte'
  import type { ExemptionElement } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

  export let exemption: ExemptionElement
</script>

<p class=" whitespace-nowrap relative group py-4 flex items-center">
  <span>{exemption.title}</span>

  {#if exemption.description}
    <span> <InformationFilled class="ml-2" /></span>
    <div class="tooltip tooltip-right">
      <div class="text-balance max-w-md text-loose">{exemption.description}</div>
    </div>
  {/if}
</p>
