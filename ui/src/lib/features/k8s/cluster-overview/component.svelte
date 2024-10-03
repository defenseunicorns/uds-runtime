<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'

  import { ProgressBarWidget, WithRightIconWidget } from '$components'
  import { Analytics, DataVis_1 } from 'carbon-icons-svelte'
  import Chart from 'chart.js/auto'

  import { calculatePercentage, formatTime, mebibytesToGigabytes, millicoresToCores } from '../helpers'
  import type { ClusterData } from '../types'
  import { chartData, chartOptions } from './chart'

  import './styles.postcss'

  let clusterData: ClusterData = {
    totalPods: 0,
    totalNodes: 0,
    cpuCapacity: 0,
    memoryCapacity: 0,
    currentUsage: {
      CPU: 0,
      Memory: 0,
      Timestamp: new Date().toISOString(),
    },
    historicalUsage: [],
  }

  let cpuPercentage = 0
  let memoryPercentage = 0
  let gbUsed = 0
  let gbCapacity = 0
  let cpuUsed = 0
  let formattedCpuCapacity = 0
  let chartjsOptions = chartOptions
  let onMessageCount = 0
  let myChart: Chart
  let chartjsData = chartData

  onMount(() => {
    let ctx = document.getElementById('chartjs-el') as HTMLCanvasElement
    const path: string = `/api/v1/monitor/cluster-overview`
    const overview = new EventSource(path)

    overview.onmessage = (event) => {
      clusterData = JSON.parse(event.data) as ClusterData

      if (clusterData && Object.keys(clusterData).length > 0) {
        const { cpuCapacity, currentUsage, historicalUsage, memoryCapacity } = clusterData
        const { CPU, Memory } = currentUsage

        cpuPercentage = calculatePercentage(CPU, cpuCapacity)
        memoryPercentage = calculatePercentage(Memory, memoryCapacity)
        gbUsed = mebibytesToGigabytes(Memory)
        gbCapacity = mebibytesToGigabytes(memoryCapacity)
        cpuUsed = millicoresToCores(CPU)
        formattedCpuCapacity = millicoresToCores(cpuCapacity)

        if (onMessageCount === 0) {
          myChart = new Chart(ctx, { type: 'line', data: chartjsData, options: chartjsOptions })
        }

        // on each message manually update the grap
        myChart.data.labels = historicalUsage.map((point) => [formatTime(point.Timestamp)])
        myChart.data.datasets[0].data = historicalUsage.map((point) => point.Memory / (1024 * 1024 * 1024))
        myChart.data.datasets[1].data = historicalUsage.map((point) => point.CPU / 1000)
        myChart.update()
        onMessageCount++
      }
    }

    Chart.register({})

    return () => {
      onMessageCount = 0
      overview.close()
      myChart.destroy()
    }
  })

  // Chart.js settings
  Chart.defaults.datasets.line.tension = 0.4
</script>

<div class="p-4 dark:text-white pt-0">
  <h1 class="text-2xl font-bold mb-4">Cluster Overview</h1>
  <div class="grid grid-cols-1 min-[1024px]:grid-cols-2 min-[1510px]:grid-cols-4 gap-4">
    <WithRightIconWidget
      statText={clusterData.totalPods.toString()}
      helperText="Pods running in cluster"
      icon={Analytics}
      link="/workloads/pods"
    />

    <WithRightIconWidget
      statText={clusterData.totalNodes.toString()}
      helperText="Nodes running in cluster"
      icon={DataVis_1}
      link="/nodes"
    />

    <ProgressBarWidget
      capacity={formattedCpuCapacity}
      progress={cpuUsed}
      statText="CPU Usage"
      unit="Cores"
      value={cpuPercentage.toFixed(2)}
    />

    <ProgressBarWidget
      capacity={gbCapacity}
      progress={gbUsed}
      statText="Memory Usage"
      unit="GB"
      value={memoryPercentage.toFixed(2)}
    />
  </div>

  <div class="mt-8">
    <h2 class="text-xl font-bold mb-4">Resource Usage Over Time</h2>

    <div class="p-5 bg-gray-800 rounded-lg overflow-hidden shadow" style:position="relative" style:margin="auto">
      <canvas id="chartjs-el" height={350} />
    </div>
  </div>

  <div class="bg-white dark:bg-gray-800 w-full relative shadow-md sm:rounded-lg overflow-hidden mt-10 px-6">
    <!-- Header which has Title, Dropdown and Search-->
    <div class="py-6 dark:border-gray-700 flex items-start">
      <div class="w-7/12 flex">
        <div class="flex h-6 items-center space-x-1">
          <h5 class="dark:text-white font-semibold justify-items-start">Event Logs</h5>

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

      <div class="w-5/12 flex space-x-3">
        <div class="w-1/3">
          <button
            id="filterDropdownButton"
            data-dropdown-toggle="filterDropdown"
            class="w-full flex items-center justify-center py-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-primary-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
            type="button"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              aria-hidden="true"
              class="h-4 w-4 mr-2 text-gray-400"
              viewbox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z"
                clip-rule="evenodd"
              />
            </svg>
            Filter
            <svg
              class="-mr-1 ml-1.5 w-5 h-5"
              fill="currentColor"
              viewbox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
              aria-hidden="true"
            >
              <path
                clip-rule="evenodd"
                fill-rule="evenodd"
                d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
              />
            </svg>
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
  </div>
</div>
