<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'

  import { goto } from '$app/navigation'
  import Unauthenticated from '$components/Auth/component.svelte'
  import { apiAuthEnabled, authenticated } from '$lib/features/api-auth/store'
  import { Auth } from '$lib/utils/api-auth'
  import { updateApiAuthEnabled } from '$lib/utils/helpers'

  export let data

  onMount(async () => {
    await updateApiAuthEnabled()
    if ($apiAuthEnabled) {
      const url = new URL(window.location.href)
      let token = url.searchParams.get('token') || ''
      if (await Auth.connect(token)) {
        // Update the store
        authenticated.set(true)
        // Update the session storage
        sessionStorage.setItem('authenticated', JSON.stringify(true))
        goto('/')
      } else {
        // Update the store
        authenticated.set(false)
        // Update the session storage
        sessionStorage.setItem('authenticated', JSON.stringify(false))
      }

      //set namespaces
      data.namespaces.start()
    }
  })
</script>

{#if $apiAuthEnabled && !$authenticated}
  <Unauthenticated />
{/if}
