// eslint-disable @typescript-eslint/no-explicit-any
import { render } from '@testing-library/svelte'
import type { Policy } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import ExemptionPolicies from './component.svelte'

describe('Policies', () => {
  test('renders empty list when no policies are provided', () => {
    const { container } = render(ExemptionPolicies, { props: { policies: [] } })
    const list = container.querySelector('ul')
    expect(list).toBeInTheDocument()
    expect(list?.children.length).toBe(0)
  })

  test('renders correct number of policy items', () => {
    const policies = ['DisallowHostNamespaces', 'DisallowNodePortServices', 'DisallowPrivileged'] as Policy[]
    const { container } = render(ExemptionPolicies, { props: { policies } })
    const listItems = container.querySelectorAll('li')
    expect(listItems.length).toBe(3)
  })

  test('renders policy names correctly', () => {
    const policies = ['DisallowHostNamespaces', 'DisallowNodePortServices', 'DisallowPrivileged'] as Policy[]
    const { getByText } = render(ExemptionPolicies, { props: { policies } })
    policies.forEach((policy) => {
      expect(getByText(`- ${policy}`)).toBeInTheDocument()
    })
  })
})
