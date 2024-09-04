<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import { onMount } from 'svelte'
  import { Auth, updateApiAuthEnabled } from '$lib/utils/api-auth'
  import { apiAuthEnabled, authenticated } from '$lib/features/api-auth/store'

  export let data

  onMount(async () => {
    await updateApiAuthEnabled()
    if ($apiAuthEnabled) {
      const url = new URL(window.location.href)
      let token = url.searchParams.get('token') || ''
      if (await Auth.connect(token)) {
        authenticated.set(true)
        goto('/')
      } else {
        authenticated.set(false) // Update the store
      }

      //set namespaces
      data.namespaces.start()
    }
  })
</script>
