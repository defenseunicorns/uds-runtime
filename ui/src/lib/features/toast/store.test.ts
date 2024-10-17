// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { get } from 'svelte/store'

import { afterEach, beforeEach, describe, expect, vi } from 'vitest'

import { addToast, removeToast, toast, type Toast } from './store'

describe('Toast Store', () => {
  beforeEach(() => {
    // Reset the store before each test
    toast.set([])
  })

  afterEach(() => {
    // Clear all timers after each test
    vi.useRealTimers()
  })

  test('should initialize with an empty array', () => {
    expect(get(toast)).toEqual([])
  })

  test('should add a toast to the store', () => {
    const newToast: Toast = {
      message: 'Test toast',
      timeoutSecs: 3,
      type: 'info',
    }
    addToast(newToast)
    const toasts = get(toast)
    expect(toasts).toHaveLength(1)
    expect(toasts[0]).toMatchObject(newToast)
    expect(toasts[0].id).toBeDefined()
  })

  test('should remove a toast from the store', () => {
    const newToast: Toast = {
      message: 'Test toast',
      timeoutSecs: 3,
      type: 'info',
    }
    addToast(newToast)
    const toasts = get(toast)
    const toastId = toasts[0].id

    removeToast(toastId)
    expect(get(toast)).toHaveLength(0)
  })

  test('should automatically remove a toast after the specified timeout', () => {
    vi.useFakeTimers()
    const newToast: Toast = {
      message: 'Test toast',
      timeoutSecs: 3,
      type: 'info',
    }
    addToast(newToast)
    expect(get(toast)).toHaveLength(1)

    vi.advanceTimersByTime(3000)
    expect(get(toast)).toHaveLength(0)
  })

  test('should not remove a toast if timeout is not specified', () => {
    vi.useFakeTimers()
    const newToast: Toast = {
      message: 'Test toast',
      timeoutSecs: 0,
      type: 'info',
    }
    addToast(newToast)
    expect(get(toast)).toHaveLength(1)

    vi.advanceTimersByTime(10000)
    expect(get(toast)).toHaveLength(1)
  })

  test('should add multiple toasts and maintain their order', () => {
    const toast1: Toast = { message: 'Toast 1', timeoutSecs: 3, type: 'info' }
    const toast2: Toast = { message: 'Toast 2', timeoutSecs: 3, type: 'success' }
    const toast3: Toast = { message: 'Toast 3', timeoutSecs: 3, type: 'error' }

    addToast(toast1)
    addToast(toast2)
    addToast(toast3)

    const toasts = get(toast)
    expect(toasts).toHaveLength(3)
    expect(toasts[0].message).toBe('Toast 1')
    expect(toasts[1].message).toBe('Toast 2')
    expect(toasts[2].message).toBe('Toast 3')
  })

  test('should not remove other toasts when removing a specific toast', () => {
    const toast1: Toast = { message: 'Toast 1', timeoutSecs: 3, type: 'info' }
    const toast2: Toast = { message: 'Toast 2', timeoutSecs: 3, type: 'success' }

    addToast(toast1)
    addToast(toast2)

    const toasts = get(toast)
    removeToast(toasts[0].id)

    const updatedToasts = get(toast)
    expect(updatedToasts).toHaveLength(1)
    expect(updatedToasts[0].message).toBe('Toast 2')
  })

  test('should do nothing when trying to remove a non-existent toast', () => {
    const toast1: Toast = { message: 'Toast 1', timeoutSecs: 3, type: 'info' }
    addToast(toast1)

    removeToast(12345) // Non-existent ID

    expect(get(toast)).toHaveLength(1)
  })
})
