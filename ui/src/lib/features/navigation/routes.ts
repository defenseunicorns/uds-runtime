// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

import { ChartPieOutline, CloudArrowUpOutline, EyeOutline, FileOutline, LayersOutline } from 'flowbite-svelte-icons'

export const routes = [
  {
    path: '/',
    name: 'Overview',
    icon: ChartPieOutline,
  },
  {
    path: '/monitor',
    name: 'Monitor',
    icon: EyeOutline,
    children: [
      {
        path: '/monitor/pepr',
        name: 'Pepr',
      },
      {
        path: '/monitor/events',
        name: 'Events',
      },
    ],
  },
  {
    path: '/resources/workloads',
    name: 'Workloads',
    icon: LayersOutline,
    children: [
      {
        path: '/resources/workloads/pods',
        name: 'Pods',
      },
      {
        path: '/resources/workloads/deployments',
        name: 'Deployments',
      },
      {
        path: '/resources/workloads/daemonsets',
        name: 'DaemonSets',
      },
      {
        path: '/resources/workloads/statefulsets',
        name: 'StatefulSets',
      },
    ],
  },
  {
    path: '/resources/config',
    name: 'Config',
    icon: FileOutline,
    children: [
      {
        path: '/resources/config/packages',
        name: 'Packages',
      },
    ],
  },
  {
    path: '/resources/network',
    name: 'Network',
    icon: CloudArrowUpOutline,
    children: [
      {
        path: '/resources/services',
        name: 'Services',
      },
    ],
  },
  {
    path: '/resources/namespaces',
    name: 'Namespaces',
    icon: LayersOutline,
  },
]
