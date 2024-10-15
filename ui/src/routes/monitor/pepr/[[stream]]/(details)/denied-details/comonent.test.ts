// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import type { SvelteComponent } from 'svelte'

import { render } from '@testing-library/svelte'

import DeniedDetails from './DeniedDetails.svelte'

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
  test('renders denied messages', () => {
    const peprDetails = { component: comp, messages: ['Authorized: test', 'Found: test'] }
    const { getByText } = render(DeniedDetails, { props: { details: peprDetails } })
    expect(getByText('Details')).toBeInTheDocument()
    expect(getByText(peprDetails.messages[0])).toBeInTheDocument()
    expect(getByText(peprDetails.messages[1])).toBeInTheDocument()
  })
})
