import { render } from '@testing-library/svelte'
import { InformationFilled } from 'carbon-icons-svelte'
import type { ExemptionElement } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import ExemptionDetails from './component.svelte'

// Mock the carbon-icons-svelte module
vi.mock('carbon-icons-svelte', () => ({
  InformationFilled: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

describe('ExemptionElement', () => {
  test('renders exemption title', () => {
    const exemption = { title: 'Test Exemption' } as ExemptionElement
    const { getByText } = render(ExemptionDetails, { props: { exemption } })
    expect(getByText('Test Exemption')).toBeInTheDocument()
  })

  it('does not render information icon when description is missing', () => {
    const exemption = { title: 'No Description Exemption' } as ExemptionElement
    const { container } = render(ExemptionDetails, { props: { exemption } })
    expect(InformationFilled).not.toHaveBeenCalled()
    expect(container.querySelector('.tooltip')).toBeNull()
  })

  it('renders information icon and tooltip when description is present', () => {
    const exemption = { title: 'With Description', description: 'Test description' } as ExemptionElement
    const { container } = render(ExemptionDetails, { props: { exemption } })
    expect(InformationFilled).toHaveBeenCalled()
    const tooltip = container.querySelector('.tooltip')
    expect(tooltip).not.toBeNull()
    expect(tooltip?.textContent).toBe('Test description')
  })
})
