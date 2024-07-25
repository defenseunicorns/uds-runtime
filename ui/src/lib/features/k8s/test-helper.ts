import { render } from '@testing-library/svelte'

import * as components from '$components'
import type { V1Deployment as Resource } from '@kubernetes/client-node'
import type { ComponentType } from 'svelte'
import type { Mock } from 'vitest'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type Procedure = (...args: any[]) => any

export function testK8sTableWithDefaults(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

export function testK8sTableWithCustomColumns(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with custom columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    props.columns = [['blah'], ['blah2']]

    render(Component, { columns: props.columns })

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

// Mocking EventSource globally
export class MockEventSource {
  onmessage: (event: MessageEvent) => void | null = () => {}
  constructor(url: string, data: Resource[], urlAssertionMock: Mock<Procedure>) {
    urlAssertionMock(url)

    setTimeout(() => {
      const messageEvent = {
        data: JSON.stringify(data),
        origin: url,
        lastEventId: '',
      }

      if (typeof this.onmessage === 'function') {
        this.onmessage(messageEvent as MessageEvent)
      }
    }, 500)
  }
}
