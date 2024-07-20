<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import type { ExemptionElement } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

  export let exemption: ExemptionElement
</script>

<p class="font-bold tracking-tight">{exemption.title}</p>

{#if exemption.description}
  <p class="mt-4">{exemption.description}</p>
{/if}

<p class="mt-6">
  Matches any <strong>{exemption.matcher.kind}</strong> in the <strong>{exemption.matcher.namespace}</strong> namespace whose
  name matches the following regular expression:
</p>

<p class="mt-2"><span class="text-sm bg-gray-600 p-1 rounded text-white">{exemption.matcher.name}</span></p>
