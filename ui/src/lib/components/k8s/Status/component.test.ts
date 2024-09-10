import { render } from '@testing-library/svelte'
import Status from './component.svelte'

describe('Status component', () => {
  test('renders text-green-400 for Running', () => {
    const { container } = render(Status, { props: { type: 'Pod', status: 'Running' } })
    expect(container.firstChild).toHaveClass('text-green-400')
  })
  test('renders text-red-400 for Failed', () => {
    const { container } = render(Status, { props: { type: 'Pod', status: 'Failed' } })
    expect(container.firstChild).toHaveClass('text-red-400')
  })
  test('renders text-orange-300 for Pending', () => {
    const { container } = render(Status, { props: { type: 'Pod', status: 'Pending' } })
    expect(container.firstChild).toHaveClass('text-orange-300')
  })
})
