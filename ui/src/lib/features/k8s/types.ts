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

export type ColumnWrapper<T> = [name: keyof T, styles?: string][]

export interface ResourceWithTable<T extends KubernetesObject, U extends CommonRow> {
  resource: T
  table: U
}
