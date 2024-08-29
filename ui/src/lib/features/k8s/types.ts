import type { KubernetesObject } from '@kubernetes/client-node'
import { type Writable } from 'svelte/store'

export interface CommonRow {
  name: string
  namespace?: string
  creationTimestamp: Date
  age?: {
    sort: number
    text: string
  }
}

export type ColumnWrapper<T> = [name: keyof T, styles?: string, modifier?: 'link-external' | 'link-internal'][]

export enum SearchByType {
  ANYWHERE = 'Anywhere',
  METADATA = 'Metadata',
  NAME = 'Name',
}

export interface ResourceWithTable<T extends KubernetesObject, U extends CommonRow> {
  resource: T
  table: U
}

export interface ResourceStoreInterface<T extends KubernetesObject, U extends CommonRow> {
  // Start the EventSource and update the resources
  start: () => () => void
  // Sort the table by the key
  sortByKey: (key: keyof U) => void
  // Store for search text
  search: Writable<string>
  // Store for search by type
  searchBy: Writable<SearchByType>
  // Store for sortBy key
  sortBy: Writable<keyof U>
  // Store for sort direction
  sortAsc: Writable<boolean>
  // The list of search types
  searchTypes: SearchByType[]
  // Subscribe to the filtered and sorted resources
  subscribe: (run: (value: ResourceWithTable<T, U>[]) => void) => () => void
  // Store for namespace
  namespace: Writable<string>
  // Store for number of resources
  numResources: Writable<number>
  // The url for the EventSource
  url: string
}

// Define specific status types for each resource
type PodStatus = 'Pending' | 'Running' | 'Succeeded' | 'Failed' | 'Unknown' | 'Completed'
type DeploymentStatus = 'Available' | 'Progressing' | 'Unavailable'
type ServiceStatus = 'Pending' | 'Active' | 'Terminating'
type PVCStatus = 'Pending' | 'Bound' | 'Lost'
type NodeStatus = 'Ready' | 'NotReady' | 'SchedulingDisabled'
type JobStatus = 'Complete' | 'Failed' | 'Running'
type CronJobStatus = 'Active' | 'Suspended'
type ConfigMapStatus = 'Active'
type SecretStatus = 'Active'
type NamespaceStatus = 'Active' | 'Terminating'

// Define a type for the k8StatusMapping
export type K8StatusMapping = {
  Pod: Record<PodStatus, { color: string }>
  Deployments: Record<DeploymentStatus, { color: string }>
  ReplicaSets: Record<DeploymentStatus, { color: string }>
  StatefulSets: Record<DeploymentStatus, { color: string }>
  Services: Record<ServiceStatus, { color: string }>
  PersistentVolumeClaims: Record<PVCStatus, { color: string }>
  Nodes: Record<NodeStatus, { color: string }>
  Jobs: Record<JobStatus, { color: string }>
  CronJobs: Record<CronJobStatus, { color: string }>
  ConfigMaps: Record<ConfigMapStatus, { color: string }>
  Secrets: Record<SecretStatus, { color: string }>
  Namespaces: Record<NamespaceStatus, { color: string }>
}
