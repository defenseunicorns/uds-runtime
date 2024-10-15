<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script>
  export let message = 'Could not authenticate'
  export let details =
    'Please make sure you are using the complete link with api token provided by UDS Runtime to connect.'
</script>

<div class="flex flex-col items-center justify-start min-h-screen">
  <h2 class="text-xl mb-4 p-4 dark:text-white pt-0">
    <strong>{message}</strong>
    : {details}
  </h2>
  <img src="/doug.svg" alt="Authentication Failed" class="mx-auto mt-4" style="width: 250px; height: 250px;" />
</div>
