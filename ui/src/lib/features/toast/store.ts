// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { get, writable } from 'svelte/store'

export type Toast = {
  id?: number
  message: string
  timeoutSecs: number
  type: 'success' | 'info' | 'warning' | 'error'
}

export const toast = writable<Toast[]>([])

export const addToast = (newToast: Toast) => {
  toast.update((toasts) => {
    // don't show duplicate toasts
    if (toasts.some((toast) => toast.message === newToast.message)) {
      return toasts
    }

    const id = Date.now() + Math.random()
    const toast = { id, ...newToast }

    if (toast.timeoutSecs) {
      setTimeout(() => removeToast(id), toast.timeoutSecs * 1000)
    }
    return [...toasts, toast]
  })
}

export const removeToast = (id?: number) => {
  toast.update((toasts) => toasts.filter((toast) => toast.id !== id))
}

export const getIdByMessage = (message: string) => {
  let id: number | undefined
  const toasts = get(toast)
  const found = toasts.find((toast) => toast.message === message)
  if (found) {
    id = found.id
  }
  return id
}
