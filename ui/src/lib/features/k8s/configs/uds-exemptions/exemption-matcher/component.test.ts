import { render } from '@testing-library/svelte'
import type { Matcher } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import ExemptionMatcher from './component.svelte'

describe('Matcher', () => {
  const defaultMatcher = {
    kind: 'pod',
    namespace: 'test-namespace',
    name: 'test-name-pattern',
  } as Matcher

  test('renders matcher details correctly', () => {
    const { getByText } = render(ExemptionMatcher, { props: { matcher: defaultMatcher } })
    expect(getByText('Kind:')).toBeInTheDocument()
    expect(getByText('pod')).toBeInTheDocument()
    expect(getByText('Namespace:')).toBeInTheDocument()
    expect(getByText('test-namespace')).toBeInTheDocument()
    expect(getByText('Name Pattern:')).toBeInTheDocument()
    expect(getByText('test-name-pattern')).toBeInTheDocument()
  })
})
