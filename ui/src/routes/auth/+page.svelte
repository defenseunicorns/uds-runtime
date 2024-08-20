<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { goto } from '$app/navigation'
  import { onMount } from 'svelte'
  import { Auth } from '$lib/utils/http'
  import { authenticated } from '$lib/features/api-auth/store'

  export let data

  let apiAuthSet = false
  let authFailure = false

  // Check if the API_AUTH environment variable is set to true (case-insensitive)
  const apiAuthEnabled = import.meta.env.VITE_API_AUTH?.toLowerCase() === 'true'
  if (apiAuthEnabled) {
    apiAuthSet = true
    onMount(async () => {
      authenticated.set(false)

      const url = new URL(window.location.href)
      let token = url.searchParams.get('token') || ''
      console.log('Token:', token)
      if (await Auth.connect(token)) {
        authenticated.set(true)
        console.log('Authenticated')
        goto('/')
      } else {
        authenticated.set(false) // Update the store
        authFailure = true
        console.log('Failed to authenticate')
      }

      //set namespaces
      data.namespaces.start()
    })
  }
</script>

{#if apiAuthSet}
  {#if authFailure}
    <h2 class="text-xl font-bold mb-4 p-4 dark:text-white pt-0">
      Could not authenticate: Please make sure you are using the complete link to connect provided by UDS Runtime.
    </h2>
  {/if}
{/if}
