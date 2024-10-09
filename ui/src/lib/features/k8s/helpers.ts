import type { K8StatusMapping } from './types'

const statusColors = {
  success: 'text-green-400',
  info: 'text-blue-400',
  warning: 'text-orange-300',
  error: 'text-red-400',
  disabled: 'text-grey-400',
} as const

// Mapping between sections, statuses, and colors
const k8StatusMapping: K8StatusMapping = {
  Pod: {
    Pending: { color: statusColors.warning },
    Running: { color: statusColors.success },
    Succeeded: { color: statusColors.success },
    Failed: { color: statusColors.error },
    Unknown: { color: statusColors.warning },
    Completed: { color: statusColors.disabled },
  },
  Deployments: {
    Available: { color: statusColors.success },
    Progressing: { color: statusColors.info },
    Unavailable: { color: statusColors.error },
  },
  ReplicaSets: {
    Available: { color: statusColors.success },
    Progressing: { color: statusColors.info },
    Unavailable: { color: statusColors.error },
  },
  StatefulSets: {
    Available: { color: statusColors.success },
    Progressing: { color: statusColors.info },
    Unavailable: { color: statusColors.error },
  },
  Services: {
    Pending: { color: statusColors.warning },
    Active: { color: statusColors.success },
    Terminating: { color: statusColors.warning },
  },
  PersistentVolumeClaims: {
    Pending: { color: statusColors.warning },
    Bound: { color: statusColors.success },
    Lost: { color: statusColors.error },
  },
  Nodes: {
    Ready: { color: statusColors.success },
    NotReady: { color: statusColors.warning },
    SchedulingDisabled: { color: statusColors.disabled },
  },
  Jobs: {
    Complete: { color: statusColors.success },
    Failed: { color: statusColors.error },
    Running: { color: statusColors.success },
  },
  CronJobs: {
    Active: { color: statusColors.success },
    Suspended: { color: statusColors.disabled },
  },
  ConfigMaps: {
    Active: { color: statusColors.success },
  },
  Secrets: {
    Active: { color: statusColors.success },
  },
  Namespaces: {
    Active: { color: statusColors.success },
    Terminating: { color: statusColors.warning },
  },
  Logs: {
    Normal: { color: statusColors.success },
    Warning: { color: statusColors.warning },
  },
}

// Function to get the color and status for a specific type and status
export const getColorForStatus = <T extends keyof K8StatusMapping>(type: T, status: keyof K8StatusMapping[T]) => {
  return (k8StatusMapping[type][status] as { color: string }).color || 'Unknown'
}

export const mebibytesToGigabytes = (value: number) => value / (1024 * 1024 * 1024)
export const millicoresToCores = (value: number) => value / 1000

export const formatTicks = (tick: string | number) => {
  if (typeof tick === 'number') {
    return tick.toFixed(2)
  }

  return tick
}

export const formatTime = (timestamp: string) => {
  const parts = new Date(timestamp).toISOString().split('T')
  parts.shift()

  return parts.join('').split('.')[0]
}

export const calculatePercentage = (usage: number, capacity: number): number => {
  if (capacity <= 0) return 0

  return Math.min(Math.max((usage / capacity) * 100, 0), 100)
}
