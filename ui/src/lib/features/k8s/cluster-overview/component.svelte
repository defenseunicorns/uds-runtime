<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'

  import { StatsWidget } from '$components'
  import ApexCharts from 'apexcharts'
  import type { ApexOptions } from 'apexcharts'
  import { Analytics, DataVis_1 } from 'carbon-icons-svelte'

  import { mebibytesToGigabytes, millicoresToCores } from '../helpers'

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
  let gbUsed = 0
  let gbCapacity = 0
  let cpuUsed = 0
  let cpuCapacity = 0

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

  let onMessageCount = 0
  let el: HTMLDivElement | undefined = undefined
  let apexChart: ApexCharts

  $: {
    options = updateClusterData(clusterData)
    apexChart?.updateOptions(options)
  }

  function updateClusterData(clusterData: ClusterData): ApexOptions {
    return {
      ...options,
      series: [
        {
          name: 'CPU Usage',
          data: (clusterData.historicalUsage ?? []).map((point) => ({
            x: new Date(point.Timestamp).getTime(),
            y: millicoresToCores(point.CPU), // Convert millicores to cores
          })),
        },
        {
          name: 'Memory Usage',
          data: (clusterData.historicalUsage ?? []).map((point) => ({
            x: new Date(point.Timestamp).getTime(),
            y: mebibytesToGigabytes(point.Memory), // Convert bytes to GB
          })),
        },
      ],
    }
  }

  onMount(() => {
    const path: string = `/api/v1/monitor/cluster-overview`
    const overview = new EventSource(path)

    overview.onmessage = (event) => {
      clusterData = JSON.parse(event.data) as ClusterData

      cpuPercentage = calculatePercentage(clusterData.currentUsage.CPU, clusterData.cpuCapacity)
      memoryPercentage = calculatePercentage(clusterData.currentUsage.Memory, clusterData.memoryCapacity)
      gbUsed = mebibytesToGigabytes(clusterData.currentUsage.Memory)
      gbCapacity = mebibytesToGigabytes(clusterData.memoryCapacity)
      cpuUsed = millicoresToCores(clusterData.currentUsage.CPU)
      cpuCapacity = millicoresToCores(clusterData.cpuCapacity)

      if (onMessageCount === 0) {
        onMessageCount++
        apexChart = new ApexCharts(el, updateClusterData(clusterData))
        apexChart?.render()
      }
    }

    return () => {
      onMessageCount = 0
      overview.close()
      apexChart?.destroy()
    }
  })
</script>

<div class="p-4 dark:text-white pt-0">
  <h1 class="text-2xl font-bold mb-4">Cluster Overview</h1>
  <div class="grid grid-cols-1 min-[1024px]:grid-cols-2 min-[1510px]:grid-cols-4 gap-4">
    <StatsWidget
      variant="with_right_icon"
      props={{
        statText: clusterData.totalPods.toString(),
        helperText: 'Pods running in cluster',
        icon: Analytics,
        link: '/workloads/pods',
      }}
    />

    <StatsWidget
      type="with_right_icon"
      props={{
        statText: clusterData.totalNodes.toString(),
        helperText: 'Nodes running in cluster',
        icon: DataVis_1,
        link: '/nodes',
      }}
    />

    <StatsWidget
      type="progress_bar"
      props={{
        capacity: cpuCapacity,
        progress: cpuUsed,
        statText: 'CPU Usage',
        unit: 'Cores',
        value: cpuPercentage.toFixed(2),
      }}
    />

    <StatsWidget
      type="progress_bar"
      props={{
        capacity: gbCapacity,
        progress: gbUsed,
        statText: 'Memory Usage',
        unit: 'GB',
        value: memoryPercentage.toFixed(2),
      }}
    />
  </div>
  <div class="mt-8">
    <h2 class="text-xl font-bold mb-4">Resource Usage Over Time</h2>
    <div class="h-96 bg-gray-800 rounded-lg overflow-hidden shadow">
      <div bind:this={el} />
    </div>
  </div>
</div>
