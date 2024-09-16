import type { SvelteComponent } from 'svelte'

import { render, screen } from '@testing-library/svelte'

import MutatedDetails from './MutatedDetails.svelte'

// Mock the carbon-icons-svelte module

const comp = vi.fn().mockImplementation(() => ({
  $$: {
    on_mount: [],
    on_destroy: [],
    before_update: [],
    after_update: [],
  },
})) as unknown as SvelteComponent

describe('Denied Details', () => {
  test('renders exemption title', () => {
    const peprDetails = {
      component: comp,
      operations: {
        ADDED: [{ op: 'add', path: '/path', value: 'value' }],
        REPLACED: [{ op: 'add', path: '/path', value: 'value' }],
        REMOVED: [{ op: 'add', path: '/path', value: '' }],
      },
    }
    render(MutatedDetails, { props: { details: peprDetails } })

    expect(screen.getByText('Details')).toBeInTheDocument()
    //todo: figure out assertions for split up text
  })
})
