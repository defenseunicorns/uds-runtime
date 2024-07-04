import { ChartPieOutline, EyeOutline, LayersOutline } from 'flowbite-svelte-icons'

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
    path: '/resources',
    name: 'Resources',
    icon: LayersOutline,
    children: [
      {
        path: '/resources/namespaces',
        name: 'Namespaces',
      },
      {
        path: '/resources/pods',
        name: 'Pods',
      },
      {
        path: '/resources/deployments',
        name: 'Deployments',
      },
      {
        path: '/resources/daemonsets',
        name: 'DaemonSets',
      },
      {
        path: '/resources/statefulsets',
        name: 'StatefulSets',
      },
      {
        path: '/resources/packages',
        name: 'Packages',
      },
      {
        path: '/resources/services',
        name: 'Services',
      },
    ],
  },
]
