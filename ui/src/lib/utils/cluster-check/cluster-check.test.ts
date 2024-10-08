import { get } from 'svelte/store'

import { toast } from '$features/toast/store'
import type { Mock } from 'vitest'

import { ClusterCheck } from './cluster-check'

/* eslint-disable @typescript-eslint/no-explicit-any */

// Vitest type redeclared cause it's not exported from vitest
type Procedure = (...args: any[]) => any

const urlAssertionMock = vi.fn()

class ClusterCheckEventSource {
  // Placeholder for the 'onmessage' event handler
  onmessage: ((event: MessageEvent) => void) | null = null
  onerror: ((event: Event) => void) | null = null
  closeEvtHandler: (() => void) | null = null

  constructor(
    url: string,
    urlAssertionMock: Mock<Procedure>,
    triggers: { msg: string; timer: number; closeEvt?: boolean }[],
  ) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)

    for (const trigger of triggers) {
      if (trigger.closeEvt) {
        // let addEventListener get set before calling handler
        setTimeout(() => {
          this.closeEvtHandler && this.closeEvtHandler()
          return
        }, 1000)
      }

      setTimeout(() => {
        this.onmessage?.(new MessageEvent('message', { data: trigger.msg }))
      }, trigger.timer)
    }
  }

  close() {}

  addEventListener(_: string, handler: () => void) {
    this.closeEvtHandler = handler
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

    new ClusterCheck()

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

    new ClusterCheck()

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
    new ClusterCheck()

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

  test('close eventSource on message containing "close" (in-cluster case)', async () => {
    const closeSpy = vi.spyOn(ClusterCheckEventSource.prototype, 'close')
    vi.stubGlobal(
      'EventSource',
      vi
        .fn()
        .mockImplementationOnce(
          (url: string) => new ClusterCheckEventSource(url, urlAssertionMock, [{ msg: '', timer: 0, closeEvt: true }]),
        ),
    )
    new ClusterCheck()

    vi.advanceTimersByTime(1000)
    expect(closeSpy).toHaveBeenCalledTimes(1)
  })
})
