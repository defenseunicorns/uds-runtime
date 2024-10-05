import { get } from 'svelte/store'

import { toast } from '$features/toast/store'
import type { Mock } from 'vitest'

import { checkClusterConnection } from './cluster-check'

/* eslint-disable @typescript-eslint/no-explicit-any */

// Vitest type redeclared cause it's not exported from vitest
type Procedure = (...args: any[]) => any

const urlAssertionMock = vi.fn()

class ClusterCheckEventSource {
  // Placeholder for the 'onmessage' event handler
  onmessage: ((event: MessageEvent) => void) | null = null
  onerror: ((event: Event) => void) | null = null

  constructor(url: string, urlAssertionMock: Mock<Procedure>, triggers: { msg: string; timer: number }[]) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)

    for (const trigger of triggers) {
      setTimeout(() => {
        this.onmessage?.(new MessageEvent('message', { data: JSON.stringify(trigger.msg) }))
      }, trigger.timer)
    }
  }
}

describe('cluster check', () => {
  beforeEach(() => {
    vi.useFakeTimers()
    toast.set([])
  })

  afterEach(() => {
    vi.unstubAllGlobals()
    vi.useRealTimers()
  })

  test('cluster check success, then error, then restored', async () => {
    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementationOnce(
        (url: string) =>
          new ClusterCheckEventSource(url, urlAssertionMock, [
            { msg: 'success', timer: 1000 },
            { msg: 'error', timer: 2000 },
            { msg: 'success', timer: 3000 },
          ]),
      ),
    )
    checkClusterConnection()

    expect(urlAssertionMock).toHaveBeenCalledWith('/health')

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(0)

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster health check failed: no connection')

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster connection restored')
  })

  test('multiple errors only show one toast', async () => {
    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementationOnce(
        (url: string) =>
          new ClusterCheckEventSource(url, urlAssertionMock, [
            { msg: 'error', timer: 1000 },
            { msg: 'error', timer: 2000 },
          ]),
      ),
    )
    checkClusterConnection()

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster health check failed: no connection')

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster health check failed: no connection')
  })

  test('event disptached on reconnection', async () => {
    const eventSpy = vi.spyOn(window, 'dispatchEvent')
    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementationOnce(
        (url: string) =>
          new ClusterCheckEventSource(url, urlAssertionMock, [
            { msg: 'error', timer: 1000 },
            { msg: 'success', timer: 2000 },
          ]),
      ),
    )
    checkClusterConnection()

    vi.advanceTimersByTime(1000)
    expect(eventSpy).toHaveBeenCalledTimes(0)

    vi.advanceTimersByTime(1000)
    expect(eventSpy).toHaveBeenCalledTimes(1)
    expect(eventSpy).toHaveBeenCalledWith(
      expect.objectContaining({
        detail: { message: 'Cluster connection restored' },
      }),
    )
  })
})
