import { type Writable } from 'svelte/store'

import type { KubernetesObject } from '@kubernetes/client-node'

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
