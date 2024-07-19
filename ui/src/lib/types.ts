export interface CommonRow {
  name: string
  namespace?: string
  creationTimestamp: Date
  age?: {
    sort: number
    text: string
  }
}

export type PeprEvent = {
  _name: string
  count: number
  event: string
  header: string
  repeated?: number
  ts?: string
  epoch: number
}
