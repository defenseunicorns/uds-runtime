import { render } from '@testing-library/svelte'

import EndpointLinks from './component.svelte'

describe('ExemptionElement', () => {
  test('renders exemption title', () => {
    const endpoints = ['grafana', 'keycloak']
    const { getAllByRole } = render(EndpointLinks, { props: { endpoints } })
    const links = getAllByRole('link')
    expect(links).toHaveLength(2)
    expect(links[0]).toHaveTextContent('grafana')
    expect(links[1]).toHaveTextContent('keycloak')
  })
})
