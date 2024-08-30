import type { K8StatusMapping } from './types'

export type StatusColorMapping = {
  success: string
  info: string
  warning: string
  error: string
  disabled: string
}

const statusColors: StatusColorMapping = {
  success: 'text-green-400',
  info: 'text-blue-400',
  warning: 'text-orange-300',
  error: 'text-red-400',
  disabled: 'text-grey-400',
}

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
}

// Function to get the color and status for a specific type and status
export const getColorAndStatus = <T extends keyof K8StatusMapping>(type: T, status: keyof K8StatusMapping[T]) => {
  return (k8StatusMapping[type][status] as { color: string }).color || 'Unknown'
}
