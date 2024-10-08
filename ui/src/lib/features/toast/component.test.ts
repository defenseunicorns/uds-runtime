// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { get } from 'svelte/store'

import { fireEvent, render, screen } from '@testing-library/svelte'
import { beforeEach, describe, expect, vi } from 'vitest'

import ToastComponent from './component.svelte'
import { addToast, toast } from './store'

// Mock the icon components
vi.mock('carbon-icons-svelte', () => {
  const mockComponent = () => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })

  return {
    CheckmarkOutline: vi.fn().mockImplementation(mockComponent),
    Close: vi.fn().mockImplementation(mockComponent),
    Information: vi.fn().mockImplementation(mockComponent),
    Warning: vi.fn().mockImplementation(mockComponent),
  }
})

describe('Toast Component', () => {
  beforeEach(() => {
    // Reset the store before each test
    toast.set([])
  })

  test('renders nothing when there are no toasts', () => {
    const { container } = render(ToastComponent)
    expect(container.firstChild?.childNodes.length).toBe(0)
  })

  test('renders a toast message', () => {
    addToast({ message: 'Test toast', timeoutSecs: 3, type: 'info' })
    render(ToastComponent)
    expect(screen.getByText('Test toast')).toBeInTheDocument()
  })

  test('renders multiple toast messages', () => {
    addToast({ message: 'Toast 1', timeoutSecs: 3, type: 'info' })
    addToast({ message: 'Toast 2', timeoutSecs: 3, type: 'success' })
    render(ToastComponent)
    expect(screen.getByText('Toast 1')).toBeInTheDocument()
    expect(screen.getByText('Toast 2')).toBeInTheDocument()
  })

  test('renders the correct number of icons for each toast type', () => {
    addToast({ message: 'Error toast', timeoutSecs: 3, type: 'error' })
    addToast({ message: 'Warning toast', timeoutSecs: 3, type: 'warning' })
    addToast({ message: 'Info toast', timeoutSecs: 3, type: 'info' })
    addToast({ message: 'Success toast', timeoutSecs: 3, type: 'success' })
    const { container } = render(ToastComponent)

    const icons = container.querySelectorAll('.w-8.h-8')
    // 4 toasts * 2 icons per toast (icon + close button)
    expect(icons.length).toBe(8)
  })

  test('removes a toast when the close button is clicked', async () => {
    addToast({ message: 'Test toast', timeoutSecs: 3, type: 'info' })
    render(ToastComponent)

    const closeButton = screen.getByRole('button')
    await fireEvent.click(closeButton)

    expect(get(toast)).toHaveLength(0)
    expect(screen.queryByText('Test toast')).not.toBeInTheDocument()
  })

  test('hides close button if noClose is true', () => {
    addToast({ message: 'Test toast', timeoutSecs: 3, type: 'info', noClose: true })
    render(ToastComponent)

    expect(screen.queryByRole('button')).not.toBeInTheDocument()
  })
})
