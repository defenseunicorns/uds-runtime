export type PeprEvent = {
  _name: string
  count: number
  event: string
  header: string
  repeated?: number
  ts?: string
  epoch: number
  msg: string
  res?: Record<string, unknown>
  details?: Record<string, any[]> | undefined
}
