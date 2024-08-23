<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'
  // @ts-expect-error types don't exist for svelte-apexcharts
  import { chart } from 'svelte-apexcharts'
  import type { ApexOptions } from 'apexcharts'

  import './styles.postcss'

  type ClusterData = {
    totalPods: number
    totalNodes: number
    cpuCapacity: number
    memoryCapacity: number
    currentUsage: {
      CPU: number
      Memory: number
      Timestamp: string
    }
    historicalUsage: {
      CPU: number
      Memory: number
      Timestamp: string
    }[]
  }

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

  function calculatePercentage(usage: number, capacity: number): number {
    if (capacity <= 0) return 0
    return Math.min(Math.max((usage / capacity) * 100, 0), 100)
  }

  function formatNumber(value: number, decimals: number = 2): string {
    return Number(value.toFixed(decimals)).toString()
  }

  function formatCPU(value: number): string {
    return formatNumber(value) + ' cores'
  }

  function formatMemory(value: number): string {
    return formatNumber(value) + ' GB'
  }

  let options: ApexOptions = {
    series: [
      {
        name: 'CPU Usage',
        data: [],
      },
      {
        name: 'Memory Usage',
        data: [],
      },
    ],
    chart: {
      type: 'line',
      height: 350,
      animations: {
        enabled: true,
        easing: 'linear',
        dynamicAnimation: {
          speed: 1000,
        },
      },
      background: '#1f2937',
      foreColor: '#e5e7eb',
    },
    stroke: {
      curve: 'smooth',
      width: 3,
    },
    xaxis: {
      type: 'datetime',
      labels: {
        style: {
          colors: '#e5e7eb',
        },
      },
    },
    yaxis: [
      {
        title: {
          text: 'CPU Usage (cores)',
          style: {
            color: '#e5e7eb',
          },
        },
        labels: {
          formatter: function (value: number) {
            return formatCPU(value)
          },
          style: {
            colors: '#e5e7eb',
          },
        },
      },
      {
        title: {
          text: 'Memory Usage (GB)',
          style: {
            color: '#e5e7eb',
          },
        },
        opposite: true,
        labels: {
          formatter: function (value: number) {
            return formatMemory(value)
          },
          style: {
            colors: '#e5e7eb',
          },
        },
      },
    ],
    legend: {
      show: true,
      labels: {
        colors: '#e5e7eb',
      },
    },
    tooltip: {
      theme: 'dark',
      x: {
        format: 'dd MMM yyyy HH:mm:ss',
      },
      y: [
        {
          formatter: function (value: number) {
            return formatCPU(value)
          },
        },
        {
          formatter: function (value: number) {
            return formatMemory(value)
          },
        },
      ],
    },
    grid: {
      borderColor: '#4b5563',
    },
  }

  $: options = {
    ...options,
    series: [
      {
        name: 'CPU Usage',
        data: clusterData.historicalUsage.map((point) => ({
          x: new Date(point.Timestamp).getTime(),
          y: point.CPU / 1000, // Convert millicores to cores
        })),
      },
      {
        name: 'Memory Usage',
        data: clusterData.historicalUsage.map((point) => ({
          x: new Date(point.Timestamp).getTime(),
          y: point.Memory / (1024 * 1024 * 1024), // Convert bytes to GB
        })),
      },
    ],
  }

  onMount(() => {
    const overview = new EventSource(`/api/v1/monitor/cluster-overview`)

    overview.onmessage = (event) => {
      clusterData = JSON.parse(event.data) as ClusterData

      cpuPercentage = calculatePercentage(clusterData.currentUsage.CPU, clusterData.cpuCapacity)
      memoryPercentage = calculatePercentage(clusterData.currentUsage.Memory, clusterData.memoryCapacity)
    }

    return () => {
      overview.close()
    }
  })
</script>

<div class="p-4 dark:text-white pt-0">
  <h1 class="text-2xl font-bold mb-4">Cluster Overview</h1>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Running Pods</dt>
        <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
          {clusterData.totalPods}
        </dd>
      </div>
    </div>
    <div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Running Nodes</dt>
        <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white" data-testid="node-count">
          {clusterData.totalNodes}
        </dd>
      </div>
    </div>
    <div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">CPU Usage</dt>
        <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
          {cpuPercentage.toFixed(2)}%
        </dd>
        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700 mt-2">
          <div class="bg-blue-600 h-2.5 rounded-full" style="width: {cpuPercentage}%"></div>
        </div>
      </div>
    </div>
    <div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
      <div class="px-4 py-5 sm:p-6">
        <dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">Memory Usage</dt>
        <dd class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
          {memoryPercentage.toFixed(2)}%
        </dd>
        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700 mt-2">
          <div class="bg-green-600 h-2.5 rounded-full" style="width: {memoryPercentage}%"></div>
        </div>
      </div>
    </div>
  </div>
  <div class="mt-8">
    <h2 class="text-xl font-bold mb-4">Resource Usage Over Time</h2>
    <div class="h-96 bg-gray-800 rounded-lg overflow-hidden shadow">
      <div use:chart={options} />
    </div>
  </div>
</div>
