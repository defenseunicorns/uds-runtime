<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { onMount } from 'svelte'

  import { ProgressBarWidget, WithRightIconWidget } from '$components'
  import { Analytics, DataVis_1 } from 'carbon-icons-svelte'
  import { type ChartData, type ChartOptions } from 'chart.js'
  import Chart from 'chart.js/auto'

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

  function formatTicks(tick: string | number) {
    if (typeof tick === 'number') {
      return tick.toFixed(2)
    }
    return tick
  }

  const formatTime = (timestamp: string) => {
    let parts = new Date(timestamp).toISOString().split('T')
    parts.shift()
    return parts.join('').split('.')[0]
  }

  function calculatePercentage(usage: number, capacity: number): number {
    if (capacity <= 0) return 0
    return Math.min(Math.max((usage / capacity) * 100, 0), 100)
  }

  function formatNumber(value: number, decimals: number = 2): string {
    return Number(value.toFixed(decimals)).toString()
  }

  let chartjsOptions: ChartOptions<'line'> = {
    maintainAspectRatio: false,
    elements: {
      point: {
        radius: 0,
      },
    },
    scales: {
      x: {
        grid: {
          display: false,
        },
        ticks: {
          color: 'white',
          maxTicksLimit: 20,
        },
      },
      y: {
        grid: {
          color: 'rgba(255, 255, 255, 0.2)',
        },
        type: 'linear',
        display: true,
        position: 'left',
        title: {
          display: true,
          text: 'CPU Usage (cores)',
          color: 'white',
          padding: {
            bottom: 15,
          },
        },
        ticks: {
          color: 'white',
          callback: (value) => `${formatTicks(value)} cores`,
        },
      },
      y1: {
        grid: {
          display: false,
        },
        type: 'linear',
        display: true,
        position: 'right',
        title: {
          display: true,
          text: 'Memory Usage (GB)',
          color: 'white',
          padding: {
            bottom: 10,
          },
        },
        ticks: {
          color: 'white',
          callback: (value) => `${formatTicks(value)} GB`,
        },
      },
    },
    plugins: {
      legend: {
        position: 'bottom',
        labels: {
          color: 'white',
          boxHeight: 14,
          boxWidth: 14,
          useBorderRadius: true,
          borderRadius: 7,
        },
      },
      tooltip: {
        enabled: true,
        mode: 'index',
        intersect: false,
        backgroundColor: '#1F2937',
        borderColor: 'white',
        borderWidth: 1,
      },
    },
    hover: {
      intersect: true,
    },
  }

  let onMessageCount = 0
  let myChart: Chart
  let chartjsData: ChartData<'line'> = {
    labels: [],
    datasets: [
      {
        label: 'Memory Usage',
        data: [],
        borderColor: '#00D39F',
        backgroundColor: '#00D39F',
        yAxisID: 'y1',
        tension: 0.4,
      },
      {
        label: 'CPU Usage',
        data: [],
        borderColor: '#057FDD',
        backgroundColor: '#057FDD',
        yAxisID: 'y',
        tension: 0.4,
      },
    ],
  }

  onMount(() => {
    let ctx = document.getElementById('chartjs-el') as HTMLCanvasElement
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
        myChart = new Chart(ctx, {
          type: 'line',
          data: chartjsData,
          options: chartjsOptions,
        })
      }

      // on each message manually update the grap
      myChart.data.labels = clusterData.historicalUsage.map((point) => [formatTime(point.Timestamp)])
      myChart.data.datasets[0].data = clusterData.historicalUsage.map((point) => point.Memory / (1024 * 1024 * 1024))
      myChart.data.datasets[1].data = clusterData.historicalUsage.map((point) => point.CPU / 1000)
      myChart.update()
      onMessageCount++
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
      capacity={cpuCapacity}
      progress={cpuUsed}
      statText="CPU Usage"
      unit="Cores"
      value={cpuPercentage.toFixed(2)}
    />

    <ProgressBarWidget
      capacity={gbCapacity}
      progress={gbUsed}
      statText="Memory Usage"
      unit="Cores"
      value={memoryPercentage.toFixed(2)}
    />
  </div>
  <div class="mt-8">
    <h2 class="text-xl font-bold mb-4">Resource Usage Over Time</h2>

    <div class="p-5 bg-gray-800 rounded-lg overflow-hidden shadow" style:position="relative" style:margin="auto">
      <canvas id="chartjs-el" height={350} />
    </div>
  </div>
</div>
