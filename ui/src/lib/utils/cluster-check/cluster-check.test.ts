import { get } from 'svelte/store'

import { toast } from '$features/toast/store'
import type { Mock } from 'vitest'

import { checkClusterConnection } from './cluster-check'

/* eslint-disable @typescript-eslint/no-explicit-any */

// Vitest type redeclared cause it's not exported from vitest
type Procedure = (...args: any[]) => any

const urlAssertionMock = vi.fn().mockImplementation((url: string) => {
  console.log(url)
})

class ClusterCheckEventSource {
  // Placeholder for the 'onmessage' event handler
  onmessage: ((event: MessageEvent) => void) | null = null
  onerror: ((event: Event) => void) | null = null

  constructor(
    url: string,
    urlAssertionMock: Mock<Procedure>,
    triggers?: { onError?: number; msg?: { success?: number; error?: number; reconnected?: number } },
  ) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)

    const msg = triggers && triggers?.msg

    if (triggers?.onError) {
      setTimeout(() => {
        this.onerror?.(new Event('error'))
      }, triggers.onError)
    }

    if (msg?.success) {
      setTimeout(() => {
        this.onmessage?.(new MessageEvent('message', { data: JSON.stringify({ success: 'success' }) }))
      }, msg.success)
    }

    if (msg?.error) {
      setTimeout(() => {
        this.onmessage?.(new MessageEvent('message', { data: JSON.stringify({ error: 'error' }) }))
      }, msg.error)
    }

    if (msg?.reconnected) {
      setTimeout(() => {
        this.onmessage?.(new MessageEvent('message', { data: JSON.stringify({ reconnected: 'reconnected' }) }))
      }, msg.reconnected)
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

  test('cluster check initial error then restored', async () => {
    vi.stubGlobal(
      'EventSource',
      vi
        .fn()
        .mockImplementationOnce(
          (url: string) =>
            new ClusterCheckEventSource(url, urlAssertionMock, { onError: 1000, msg: { reconnected: 2000 } }),
        ),
    )
    checkClusterConnection()
    expect(urlAssertionMock).toHaveBeenCalledWith('/health')
    expect(get(toast)).toHaveLength(0)

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster health check failed: no connection')

    vi.advanceTimersByTime(1000)
    expect(get(toast)).toHaveLength(1)
    expect(get(toast)[0].message).toBe('Cluster connection restored')
  })

  test('cluster check success, then error, then restored', async () => {
    vi.stubGlobal(
      'EventSource',
      vi.fn().mockImplementationOnce(
        (url: string) =>
          new ClusterCheckEventSource(url, urlAssertionMock, {
            msg: { success: 1000, error: 2000, reconnected: 3000 },
          }),
      ),
    )
    checkClusterConnection()

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
      vi
        .fn()
        .mockImplementationOnce(
          (url: string) => new ClusterCheckEventSource(url, urlAssertionMock, { onError: 1000, msg: { error: 2000 } }),
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
      vi
        .fn()
        .mockImplementationOnce(
          (url: string) =>
            new ClusterCheckEventSource(url, urlAssertionMock, { onError: 1000, msg: { reconnected: 2000 } }),
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