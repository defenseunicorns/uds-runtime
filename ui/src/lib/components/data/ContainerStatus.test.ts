import { describe, it, expect } from 'vitest'
import { render } from '@testing-library/svelte'
import ContainerStatus from './ContainerStatus.svelte'
import type { V1ContainerStatus } from '@kubernetes/client-node/dist/gen/models/V1ContainerStatus'

describe('ContainerStatus', () => {
  it('renders nothing when no containers are provided', () => {
    const { container } = render(ContainerStatus, { props: { containers: [] } })
    expect(container.querySelector('.flex')).toBeTruthy()
    expect(container.querySelectorAll('.w-2')).toHaveLength(0)
  })

  it('renders correct number of status indicators', () => {
    const containers: V1ContainerStatus[] = [
      { state: { running: {} } },
      { state: { waiting: {} } },
      { state: { terminated: {} } },
    ] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    // Double due to the tooltip
    expect(container.querySelectorAll('.w-2')).toHaveLength(3 * 2)
  })

  it('applies correct classes for running containers', () => {
    const containers: V1ContainerStatus[] = [{ state: { running: {} } }] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    const indicator = container.querySelector('.w-2')
    expect(indicator).toHaveClass('bg-green-500')
  })

  it('applies correct classes for waiting containers', () => {
    const containers: V1ContainerStatus[] = [{ state: { waiting: {} } }] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    const indicator = container.querySelector('.w-2')
    expect(indicator).toHaveClass('bg-yellow-500')
  })

  it('applies correct classes for terminated containers', () => {
    const containers: V1ContainerStatus[] = [{ state: { terminated: {} } }] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    const indicator = container.querySelector('.w-2')
    expect(indicator).toHaveClass('bg-gray-500')
  })

  it('applies pulsing class for containers that are not ready', () => {
    const containers: V1ContainerStatus[] = [{ ready: false, state: { running: {} } }] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    const indicator = container.querySelector('.w-2')
    expect(indicator).toHaveClass('animate-pulse')
  })

  it('handles mixed container states correctly', () => {
    const containers: V1ContainerStatus[] = [
      { state: { running: {} }, ready: true },
      { state: { running: {} }, ready: false },
      { state: { waiting: {} } },
      { state: { terminated: {} } },
    ] as V1ContainerStatus[]
    const { container } = render(ContainerStatus, { props: { containers } })
    const indicators = container.querySelectorAll('.w-2')
    expect(indicators[0]).toHaveClass('bg-green-500')
    expect(indicators[0]).not.toHaveClass('animate-pulse')
    expect(indicators[1]).toHaveClass('bg-green-500', 'animate-pulse')
    expect(indicators[2]).toHaveClass('bg-yellow-500')
    expect(indicators[3]).toHaveClass('bg-gray-500')
  })
})
