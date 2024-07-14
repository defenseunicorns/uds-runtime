// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { ChartCombo, KubernetesPod, Layers, Network_2, SearchLocate, TextAlignLeft } from 'carbon-icons-svelte'

import type { Route } from './types'

export const routes: Route[] = [
  {
    path: '/',
    name: 'Overview',
    icon: ChartCombo,
  },
  {
    path: '/monitor',
    name: 'Monitor',
    icon: SearchLocate,
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
    icon: KubernetesPod,
    class: 'top-border',
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
    icon: TextAlignLeft,
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
    icon: Network_2,
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
    icon: Layers,
  },
]
