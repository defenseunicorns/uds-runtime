// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { ChartData, ChartOptions } from 'chart.js'

import { formatTicks } from '../helpers'

const greyedOutStyles = {
  gridColor: 'rgba(128, 128, 128, 0.2)',
  textColor: 'rgba(128, 128, 128, 0.8)',
  borderColor: 'rgba(128, 128, 128, 0.8)',
  backgroundColor: 'rgba(128, 128, 128, 0.2)',
}

const normalStyles = {
  gridColor: 'rgba(255, 255, 255, 0.2)',
  textColor: 'white',
  borderColor: 'rgba(75, 192, 192, 1)',
  backgroundColor: 'rgba(75, 192, 192, 0.2)',
}

export const getChartOptions = (metricsServerAvailable: boolean): ChartOptions<'line'> => {
  const styles = metricsServerAvailable ? normalStyles : greyedOutStyles

  return {
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
          color: styles.textColor,
          maxTicksLimit: 20,
        },
      },
      y: {
        grid: {
          color: styles.gridColor,
        },
        type: 'linear',
        display: true,
        position: 'left',
        title: {
          display: true,
          text: 'CPU Usage (cores)',
          color: styles.textColor,
          padding: {
            bottom: 15,
          },
        },
        ticks: {
          color: styles.textColor,
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
          color: styles.textColor,
          padding: {
            bottom: 10,
          },
        },
        ticks: {
          color: styles.textColor,
          callback: (value: number | string) => `${formatTicks(value)} GB`,
        },
      },
    },
    plugins: {
      legend: {
        position: 'bottom',
        labels: {
          color: styles.textColor,
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
        borderColor: styles.textColor,
        borderWidth: 1,
      },
    },
    hover: {
      intersect: true,
    },
  }
}

const greyedOutDatasetStyles = {
  memoryUsage: {
    borderColor: 'rgba(128, 128, 128, 0.8)',
    backgroundColor: 'rgba(128, 128, 128, 0.2)',
  },
  cpuUsage: {
    borderColor: 'rgba(128, 128, 128, 0.8)',
    backgroundColor: 'rgba(128, 128, 128, 0.2)',
  },
}

const normalDatasetStyles = {
  memoryUsage: {
    borderColor: '#00D39F',
    backgroundColor: '#00D39F',
  },
  cpuUsage: {
    borderColor: '#057FDD',
    backgroundColor: '#057FDD',
  },
}

export const getChartData = (metricsServerAvailable: boolean): ChartData<'line'> => {
  const styles = metricsServerAvailable ? normalDatasetStyles : greyedOutDatasetStyles

  return {
    labels: [],
    datasets: [
      {
        label: 'Memory Usage',
        data: [],
        borderColor: styles.memoryUsage.borderColor,
        backgroundColor: styles.memoryUsage.backgroundColor,
        yAxisID: 'y1',
        tension: 0.4,
      },
      {
        label: 'CPU Usage',
        data: [],
        borderColor: styles.cpuUsage.borderColor,
        backgroundColor: styles.cpuUsage.backgroundColor,
        yAxisID: 'y',
        tension: 0.4,
      },
    ],
  }
}
