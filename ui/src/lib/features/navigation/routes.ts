// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import {
  AudioConsole,
  BareMetalServer,
  ChartCombo,
  DataTableReference,
  Db2Database,
  KubernetesPod,
  Layers,
  Network_2,
  SearchLocate,
  WorkflowAutomation,
} from 'carbon-icons-svelte'

import type { BaseRoute, Route } from './types'

const baseRoutes: BaseRoute[] = [
  {
    name: 'Overview',
    icon: ChartCombo,
  },
  {
    name: 'Applications',
    icon: DataTableReference,
    children: ['Zarf Packages'],
  },
  {
    name: 'Monitor',
    icon: SearchLocate,
    children: ['Pepr', 'Events'],
  },
  {
    name: 'Workloads',
    icon: KubernetesPod,
    class: 'top-border',
    children: ['Pods', 'Deployments', 'DaemonSets', 'StatefulSets', 'Jobs', 'CronJobs'],
  },
  {
    name: 'Configs',
    icon: AudioConsole,
    class: 'icon-rotate-90',
    children: ['UDS Packages', 'UDS Exemptions', 'ConfigMaps', 'Secrets'],
  },
  {
    name: 'Cluster Ops',
    icon: WorkflowAutomation,
    children: [
      'Mutating Webhooks',
      'Validating Webhooks',
      'HPA',
      'Pod Disruption Budgets',
      'Resource Quotas',
      'Limit Ranges',
      'Priority Classes',
      'Runtime Classes',
    ],
  },
  {
    name: 'Networks',
    icon: Network_2,
    children: ['Services', 'Virtual Services', 'Network Policies', 'Endpoints'],
  },
  {
    name: 'Storage',
    icon: Db2Database,
    children: ['Persistent Volumes', 'Persistent Volume Claims', 'Storage Classes'],
  },
  {
    name: 'Namespaces',
    icon: Layers,
  },
  {
    name: 'Nodes',
    icon: BareMetalServer,
  },
]

// Convert the path to a URL-friendly format
const createPath = (name: string) => `/${name.replace(/\s+/g, '-').toLowerCase()}`

// Convert the base routes to routes
export const routes: Route[] = baseRoutes.map(({ name, children, ...rest }) => ({
  ...rest,
  name,
  path: createPath(name),
  children: children?.map((name) => ({ name, path: createPath(name) })),
}))
