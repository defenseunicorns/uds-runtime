// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import type { ContainerMetric } from '@kubernetes/client-node'
import { render, screen, within } from '@testing-library/svelte'

import ContainerMetrics from './component.svelte'

// Mock the utility functions
vi.mock('./utils', () => ({
  formatCPU: vi.fn((cpu) => `${cpu} cpu`),
  formatMemory: vi.fn((memory) => `${memory} Mi`),
  parseCPU: vi.fn((cpu) => parseInt(cpu)),
  parseMemory: vi.fn((memory) => parseInt(memory)),
}))

suite('ContainerMetrics', () => {
  const mockContainers: ContainerMetric[] = [
    {
      name: 'container1',
      usage: {
        cpu: '100',
        memory: '100',
      },
    },
    {
      name: 'container2',
      usage: {
        cpu: '200',
        memory: '200',
      },
    },
  ]

  beforeEach(() => {
    render(ContainerMetrics, { props: { containers: mockContainers } })
  })

  test('displays total CPU usage', () => {
    expect(screen.getByText('CPU: 300 cpu')).toBeInTheDocument()
  })

  test('displays total memory usage', () => {
    expect(screen.getByText('Mem: 300 Mi')).toBeInTheDocument()
  })

  test('displays individual container information in tooltip', () => {
    const tooltip = screen.getByText('container1:').closest('.tooltip')
    expect(tooltip).not.toBeNull()

    if (tooltip) {
      mockContainers.forEach((container) => {
        const containerSection = within(tooltip as HTMLElement)
          .getByText(container.name + ':')
          .closest('div')
        expect(containerSection).not.toBeNull()

        if (containerSection) {
          expect(containerSection).toHaveTextContent(`CPU: ${container.usage.cpu} cpu`)
          expect(containerSection).toHaveTextContent(`Mem: ${container.usage.memory} Mi`)
        }
      })
    }
  })

  test('does not render anything when containers array is empty', () => {
    render(ContainerMetrics, { props: { containers: [] } })
    expect(screen.queryByText('CPU:')).not.toBeInTheDocument()
    expect(screen.queryByText('Mem:')).not.toBeInTheDocument()
  })
})
