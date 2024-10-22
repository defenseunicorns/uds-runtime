<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { onMount } from 'svelte'

  import type { V1Pod } from '@kubernetes/client-node'
  import { CoreServicesWidget, ProgressBarWidget, WithRightIconWidget } from '$components'
  import EventsOverviewWidget from '$components/k8s/Event/EventsOverviewWidget.svelte'
  import { createStore } from '$lib/features/k8s/events/store'
  import { type CoreServiceType } from '$lib/types'
  import { resourceDescriptions } from '$lib/utils/descriptions'
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
  let onMessageCount = 0
  let myChart: Chart
  const description = resourceDescriptions['Events']
  let coreServices: CoreServiceType[] = []
  let pods: V1Pod[] = []

  onMount(() => {
    let ctx = document.getElementById('chartjs-el') as HTMLCanvasElement
    const overviewPath: string = '/api/v1/monitor/cluster-overview'
    const coreServicesPath: string = '/api/v1/resources/configs/uds-packages?fields=.metadata.name,.status.phase'
    const podsPath: string = '/api/v1/resources/workloads/pods?fields=.metadata.name'

    const overview = new EventSource(overviewPath)
    const coreServicesEvent = new EventSource(coreServicesPath)
    const podsEvent = new EventSource(podsPath)

    coreServicesEvent.onmessage = (event) => {
      coreServices = JSON.parse(event.data) as CoreServiceType[]
    }

    podsEvent.onmessage = (event) => {
      pods = JSON.parse(event.data) as V1Pod[]
    }

    overview.onmessage = (event) => {
      clusterData = JSON.parse(event.data) as ClusterData

      if (clusterData && Object.keys(clusterData).length > 0) {
        const { cpuCapacity, currentUsage, historicalUsage, memoryCapacity } = clusterData
        let { CPU, Memory } = currentUsage

        // Handle case where CPU or Memory is -1 indicating metrics server is not available. Don't want to display negative values
        if (CPU == -1) {
          CPU = 0
        }
        if (Memory == -1) {
          Memory = 0
        }

        cpuPercentage = calculatePercentage(CPU, cpuCapacity)
        memoryPercentage = calculatePercentage(Memory, memoryCapacity)
        gbUsed = mebibytesToGigabytes(Memory)
        gbCapacity = mebibytesToGigabytes(memoryCapacity)
        cpuUsed = millicoresToCores(CPU)
        formattedCpuCapacity = millicoresToCores(cpuCapacity)

        if (onMessageCount === 0) {
          myChart = new Chart(ctx, { type: 'line', data: chartData, options: chartOptions })
        }

        // on each message manually update the graph
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
      coreServicesEvent.close()
      podsEvent.close()
      myChart.destroy()
    }
  })

  // Chart.js settings
  Chart.defaults.datasets.line.tension = 0.4
</script>

<div class="p-4 dark:text-white pt-0">
  <h1 class="text-xl font-bold mb-4">Cluster Overview</h1>
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

  <div class="mt-8 flex flex-col xl:flex-row xl:space-x-4">
    <div class="w-full mt-4">
      <div class="p-5 bg-gray-800 rounded-lg overflow-hidden shadow">
        <CoreServicesWidget {coreServices} {pods} />
      </div>
    </div>
  </div>

  <div class="mt-8">
    <h2 class="text-xl font-bold mb-4">Resource Usage Over Time</h2>

    <div class="p-5 bg-gray-800 rounded-lg overflow-hidden shadow" style:position="relative" style:margin="auto">
      <canvas id="chartjs-el" height={350} />
    </div>
  </div>

  <EventsOverviewWidget {createStore} {description} />
</div>
