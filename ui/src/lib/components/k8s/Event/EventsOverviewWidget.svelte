<script lang="ts">
  import { goto } from '$app/navigation'
  import { ChevronDown, ChevronRight, Filter, Time } from 'carbon-icons-svelte'

  type DropdownProps = {
    title: string
    options: string[]
  }

  export let title: string
  export let dropdown: DropdownProps
</script>

<div class="flex flex-col">
  <div class="bg-white dark:bg-gray-800 w-full relative shadow-md rounded-lg overflow-hidden mt-10">
    <!-- Header which has Title, Dropdown and Search-->
    <div class="p-6 dark:border-gray-700 flex items-start">
      <div class="w-3/12 min-[1400px]:w-6/12 flex">
        <div class="flex h-6 items-center space-x-1">
          <h5 class="dark:text-white font-semibold justify-items-start">{title}</h5>

          <div data-tooltip-target="results-tooltip">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 text-gray-400"
              viewbox="0 0 20 20"
              fill="currentColor"
              aria-hidden="true"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                clip-rule="evenodd"
              />
            </svg>

            <span class="sr-only">More info</span>

            <div
              id="results-tooltip"
              role="tooltip"
              class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700"
            >
              Showing 1-10 of 6,560 results
              <div class="tooltip-arrow" data-popper-arrow=""></div>
            </div>
          </div>
        </div>
      </div>

      <div class="w-9/12 min-[1400px]:w-6/12 flex space-x-3">
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
    <div class="flex flex-col text-xs">
      <div class="row-header flex text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 px-6 py-5">
        <div class="w-2/12">STATUS</div>
        <div class="w-8/12">EVENT</div>
        <div class="w-2/12">TIMESTAMP</div>
      </div>

      <div class="rows-cells flex p-6">
        <div class="w-2/12 content-center"><span class="py-1.5 px-4 rounded-md bg-red-900">Completed</span></div>
        <div class="w-8/12">Payment from John Doe</div>
        <div class="w-2/12">April 23, 2024</div>
      </div>
    </div>

    <!-- Footer with link-->
    <div class="bg-white dark:bg-gray-800 rounded-b-lg px-10 h-20 flex items-center justify-end">
      <button class="text-sm dark:text-blue-300 flex items-center space-x-1" on:click={() => goto('/monitor/events')}>
        <span>VIEW EVENTS</span>
        <ChevronRight />
      </button>
    </div>
  </div>
</div>
