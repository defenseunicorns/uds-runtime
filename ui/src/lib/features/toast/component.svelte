<script lang="ts">
  import { CheckmarkOutline, Close, Information, Warning } from 'carbon-icons-svelte'

  import { removeToast, toast } from './store'
</script>

<div class="fixed top-11 right-5 p-4 z-40">
  {#each $toast as toast}
    <div
      class="flex items-center justify-between shadow-gray-900 w-full max-w-xs p-4 text-gray-500 bg-white rounded-lg shadow-lg dark:text-gray-400 dark:bg-gray-800 mb-4"
      class:border-red-500={toast.type === 'error'}
      class:border-yellow-500={toast.type === 'warning'}
      class:border-blue-500={toast.type === 'info'}
      class:border-green-500={toast.type === 'success'}
      class:border-2={'true'}
    >
      <div class="flex items-center space-x-4">
        <div class="flex-shrink-0 w-8 h-8 rounded-lg">
          {#if toast.type === 'error'}
            <Warning class="w-8 h-8 text-red-500" />
          {:else if toast.type === 'warning'}
            <Warning class="w-8 h-8 text-yellow-500" />
          {:else if toast.type === 'info'}
            <Information class="w-8 h-8 text-blue-500" />
          {:else if toast.type === 'success'}
            <CheckmarkOutline class="w-8 h-8 text-green-500" />
          {/if}
        </div>
        <div class="text-sm font-normal">{toast.message}</div>
      </div>
      {#if !toast.noClose}
        <button
          type="button"
          class="flex-shrink-0 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700"
          on:click={() => removeToast(toast.id)}
        >
          <Close class="w-5 h-5" />
        </button>
      {/if}
    </div>
  {/each}
</div>
