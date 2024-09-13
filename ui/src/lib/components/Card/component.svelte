<script lang="ts">
  import { goto } from '$app/navigation'

  export let title: string = ''
  export let text: string = ''
  export let route: string = '/'
</script>

<button
  class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg hover:dark:bg-gray-700 flex"
  on:click={() => goto(route)}
>
  <div class="px-4 py-5 sm:p-6 flex flex-col items-start">
    <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">{title}</dt>
    <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
      {text}
    </dd>
  </div>
</button>
