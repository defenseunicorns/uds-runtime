<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import type { Matcher } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

  export let matcher: Matcher
</script>

<div class="p-2 text-sm">
  <div class="grid grid-cols-[auto,1fr] gap-x-4 items-baseline">
    <span class="text-gray-500">Kind:</span>
    <span class="font-medium">{matcher.kind}</span>
    <span class="text-gray-500">Namespace:</span>
    <span class="font-medium">{matcher.namespace}</span>
    <span class="text-gray-500">Name Pattern:</span>
    <div class="overflow-hidden text-ellipsis">
      <span
        class="bg-gray-600 px-2 py-0.5 rounded text-white text-xs max-w-full overflow-hidden text-ellipsis whitespace-nowrap"
      >
        {matcher.name}
      </span>
    </div>
  </div>
</div>
