<script lang="ts">
  import { goto } from '$app/navigation'
  import { ChevronDown, ChevronRight, Information, Time } from 'carbon-icons-svelte'

  type DropdownProps = {
    title: string
    options: string[]
  }

  export let title: string
  export let dropdown: DropdownProps
</script>

<div class="events">
  <!-- Header which has Title, Dropdown and Search-->
  <div class="events__header">
    <div class="w-3/12 min-[1400px]:w-6/12 flex">
      <div class="flex h-6 items-center space-x-1">
        <h2 class="events__header-title">{title}</h2>

        <div class="relative group">
          <Information class="ml-1 w-4 h-4 text-gray-400" />

          <div class="tooltip tooltip-right min-w-72">
            <div class="whitespace-normal">Showing 1-10 of 6,560 results</div>
          </div>
        </div>
      </div>
    </div>

    <div class="w-9/12 min-[1400px]:w-6/12 flex space-x-5">
      <div class="w-1/3 flex justify-end">
        <button
          id="filterDropdownButton"
          data-dropdown-toggle="filterDropdown"
          class="hover:text-primary-700 flex items-center justify-center rounded-lg border border-gray-200 bg-white px-3 py-2 text-sm font-medium text-gray-900 hover:bg-gray-100 focus:z-10 focus:outline-none focus:ring-4 focus:ring-gray-200 md:w-auto dark:border-gray-600 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:ring-gray-700"
          type="button"
          data-testid="datatable-filter-dropdown"
        >
          <Time class="mr-1.5" style="height: 12px" />
          {dropdown.title}
          <ChevronDown class="ml-1" />
        </button>

        <div id="filterDropdown" class="z-10 hidden w-48 p-3 bg-white rounded-lg shadow dark:bg-gray-700">
          <h6 class="mb-3 text-sm font-medium text-gray-900 dark:text-white">Choose brand</h6>
          <ul class="space-y-2 text-sm" aria-labelledby="filterDropdownButton">
            <li class="flex items-center">
              <input
                id="apple"
                type="checkbox"
                value=""
                class="w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500"
              />
              <label for="apple" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-100">Apple (56)</label>
            </li>
            <li class="flex items-center">
              <input
                id="fitbit"
                type="checkbox"
                value=""
                class="w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500"
              />
              <label for="fitbit" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-100">
                Microsoft (16)
              </label>
            </li>
            <li class="flex items-center">
              <input
                id="razor"
                type="checkbox"
                value=""
                class="w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500"
              />
              <label for="razor" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-100">Razor (49)</label>
            </li>
            <li class="flex items-center">
              <input
                id="nikon"
                type="checkbox"
                value=""
                class="w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500"
              />
              <label for="nikon" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-100">Nikon (12)</label>
            </li>
            <li class="flex items-center">
              <input
                id="benq"
                type="checkbox"
                value=""
                class="w-4 h-4 bg-gray-100 border-gray-300 rounded text-primary-600 focus:ring-primary-500 dark:focus:ring-primary-600 dark:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500"
              />
              <label for="benq" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-100">BenQ (74)</label>
            </li>
          </ul>
        </div>
      </div>

      <div class="w-2/3">
        <form class="flex items-center">
          <label for="simple-search" class="sr-only">Search</label>
          <div class="relative w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <svg
                aria-hidden="true"
                class="w-5 h-5 text-gray-500 dark:text-gray-400"
                fill="currentColor"
                viewbox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  fill-rule="evenodd"
                  d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>

            <input
              type="text"
              id="simple-search"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full pl-10 p-2 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
              placeholder="Search"
              required
            />
          </div>
        </form>
      </div>
    </div>
  </div>

  <!-- Rows -->
  <div class="events__rows">
    <div class="events__rows-header">
      <div class="w-2/12">STATUS</div>
      <div class="w-8/12">EVENT</div>
      <div class="w-2/12">TIMESTAMP</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-green-900">Completed</span></div>
      <div class="w-8/12">Payment from Bonnie Green</div>
      <div class="w-2/12">April 23, 2024</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-green-900">Completed</span></div>
      <div class="w-8/12">Payment refund from #00910</div>
      <div class="w-2/12">April 23, 2024</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-red-900">Cancelled</span></div>
      <div class="w-8/12">Payment refund from #087651</div>
      <div class="w-2/12">April 18, 2024</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-red-900">Cancelled</span></div>
      <div class="w-8/12">Payment refund from #087651</div>
      <div class="w-2/12">April 18, 2024</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-pink-900">In Progress</span></div>
      <div class="w-8/12">Payment from Jose Leos</div>
      <div class="w-2/12">April 18, 2024</div>
    </div>

    <div class="events__rows-item">
      <div class="w-2/12 content-center"><span class="py-1.5 px-2.5 rounded-md bg-green-900">Completed</span></div>
      <div class="w-8/12">Payment from THEMSBERG LLC</div>
      <div class="w-2/12">April 18, 2024</div>
    </div>
  </div>

  <!-- Footer with link-->
  <div class="events__footer">
    <button class="text-sm dark:text-blue-300 flex items-center space-x-1" on:click={() => goto('/monitor/events')}>
      <span>VIEW EVENTS</span>
      <ChevronRight />
    </button>
  </div>
</div>

<style lang="postcss">
  .events {
    @apply bg-white dark:bg-gray-800 w-full relative shadow-md rounded-lg overflow-hidden mt-10;
  }

  .events__header {
    @apply p-6 dark:border-gray-700 flex items-start;
  }

  .events__header-title {
    @apply text-lg dark:text-white font-semibold justify-items-start;
  }

  .events__rows {
    @apply flex flex-col text-xs;
  }

  .events__rows-header {
    @apply flex text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 px-6 py-5;
  }

  .events__rows-item {
    @apply flex p-6 border-b border-b-gray-600;
  }

  .events__footer {
    @apply bg-white dark:bg-gray-800 rounded-b-lg px-10 h-20 flex items-center justify-end;
  }
</style>
