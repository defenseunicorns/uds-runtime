<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import Row from '$components/k8s/Drawer/Row.svelte'
</script>

<h1 class="text-white text-5xl mb-7">Home</h1>
<Row label="Label" data={{ variant: 'text', value: 'something' }} />
<Row label="Label" data={{ variant: 'icon-text', value: 'something' }} />
<Row label="Link" data={{ variant: 'external-link', value: 'something' }} />
<Row label="Link" data={{ variant: 'internal-link', value: 'something' }} />
<Row
  label="Link"
  data={{
    variant: 'input-copy',
    value: ['/etc/link-1', '/etc/bin/link-2', '/var/run/secrets/kubernetes.io/serviceaccount'],
  }}
/>
<Row label="Link" data={{ variant: 'badges', value: ['badge', 'badge', 'badge', 'badge', 'badge', 'badge'] }} />
