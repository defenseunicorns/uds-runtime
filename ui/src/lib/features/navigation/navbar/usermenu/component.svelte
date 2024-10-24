<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { onMount } from 'svelte'
  import { fade } from 'svelte/transition'

  import { goto } from '$app/navigation'
  import type { UserData } from '$features/navigation/types'
  import { ChevronDown, ChevronUp, UserAvatarFilled } from 'carbon-icons-svelte'

  export let userData: UserData

  let dropdownContainer: HTMLElement
  let isOpen = false

  onMount(() => {
    const handleWindowClick = (event: MouseEvent) => {
      if (dropdownContainer && !dropdownContainer.contains(event.target as Node)) {
        isOpen = false
      }
    }

    window.addEventListener('click', handleWindowClick)

    return () => {
      window.removeEventListener('click', handleWindowClick)
    }
  })

  function toggleMenu() {
    isOpen = !isOpen
  }

  function signOut() {
    goto('/logout')
  }
</script>

<div class="relative inline-block text-left" bind:this={dropdownContainer}>
  <button
    on:click={toggleMenu}
    class="inline-flex bg-gray-800 items-center justify-center w-full rounded-md px-4 py-2 text-sm font-medium text-white focus:outline-none transition-colors duration-200 ease-in-out"
    class:bg-gray-800={!isOpen}
    class:hover:bg-gray-700={!isOpen}
    class:bg-gray-700={isOpen}
  >
    <UserAvatarFilled class="h-5 w-5 mr-2" />
    <span>{userData.name}</span>
    {#if isOpen}
      <ChevronUp class="h-4 w-4 ml-2" />
    {:else}
      <ChevronDown class="h-4 w-4 ml-2" />
    {/if}
  </button>

  {#if isOpen}
    <div
      transition:fade={{ duration: 100 }}
      class="origin-top-right absolute right-0 mt-1 w-56 rounded-md shadow-lg bg-gray-700 focus:outline-none ring-1 ring-black ring-opacity-5"
    >
      <div class="py-1">
        <div class="px-4 py-2 text-sm text-white font-medium border-b border-gray-600 truncate">
          <p>{userData.preferredUsername}</p>
          <p class="text-xs text-gray-400 mt-1 truncate">{userData.group}</p>
        </div>
        <button on:click={signOut} class="block w-full text-left px-4 py-2 text-sm text-white hover:bg-gray-600">
          Sign Out
        </button>
      </div>
    </div>
  {/if}
</div>
