<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import Row from '$components/Drawer/Row.svelte'
  import { Variants } from '$components/Drawer/Row.svelte'
</script>

<h1 class="text-white text-5xl mb-7">Home</h1>
<Row label="Label" data={{ variant: Variants.TEXT, value: 'something' }} />
<Row label="Label" data={{ variant: Variants.ICON_TEXT, value: 'something' }} />
<Row label="Link" data={{ variant: Variants.EXTERNAL_LINK, value: 'something' }} />
<Row label="Link" data={{ variant: Variants.INTERNAL_LINK, value: 'something' }} />
<Row
  label="Link"
  data={{
    variant: Variants.INPUT_COPY,
    value: ['/etc/link-1', '/etc/bin/link-2', '/var/run/secrets/kubernetes.io/serviceaccount'],
  }}
/>
<Row label="Link" data={{ variant: Variants.BADGES, value: ['badge', 'badge', 'badge', 'badge', 'badge', 'badge'] }} />
