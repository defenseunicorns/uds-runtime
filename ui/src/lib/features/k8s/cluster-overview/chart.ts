import type { ChartData, ChartOptions } from 'chart.js'

import { formatTicks } from '../helpers'

export const chartOptions: ChartOptions<'line'> = {
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
        callback: (value: number | string) => `${formatTicks(value)} cores`,
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
        callback: (value: number | string) => `${formatTicks(value)} GB`,
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

export const chartData: ChartData<'line'> = {
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
