<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { Information } from 'carbon-icons-svelte'
  import type { ExemptionElement } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

  export let exemption: ExemptionElement
</script>

<p class=" whitespace-nowrap relative group py-4 flex items-center">
  <span>{exemption.title}</span>

  {#if exemption.description}
    <div class="flex relative items-center">
      <Information class="ml-2" />
      <div class="tooltip tooltip-right -translate-y-4">
        <div class="text-balance max-w-md text-loose">{exemption.description}</div>
      </div>
    </div>
  {/if}
</p>
