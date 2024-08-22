<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import { onMount } from 'svelte'
  import { Auth } from '$lib/utils/http'
  import { authenticated } from '$lib/features/api-auth/store'
  import { apiAuthEnabled } from '$lib/features/api-auth/store'
  import { updateApiAuthEnabled } from '$lib/utils/http'

  export let data

  let authFailure = false

  onMount(async () => {
    authenticated.set(false)
    await updateApiAuthEnabled()
    if ($apiAuthEnabled) {
      const url = new URL(window.location.href)
      let token = url.searchParams.get('token') || ''
      if (await Auth.connect(token)) {
        authenticated.set(true)
        goto('/')
      } else {
        authenticated.set(false) // Update the store
        authFailure = true
      }

      //set namespaces
      data.namespaces.start()
    }
  })
</script>

{#if apiAuthEnabled}
  {#if authFailure}
    <div class="flex flex-col items-center justify-start min-h-screen">
      <h2 class="text-xl mb-4 p-4 dark:text-white pt-0">
        <strong>Could not authenticate</strong>
        : Please make sure you are using the complete link with api token provided by UDS Runtime to connect.
      </h2>
      <img src="/doug.svg" alt="Authentication Failed" class="mx-auto mt-4" style="width: 250px; height: 250px;" />
    </div>
  {/if}
{/if}
