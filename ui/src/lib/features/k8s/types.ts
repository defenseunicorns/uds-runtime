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
