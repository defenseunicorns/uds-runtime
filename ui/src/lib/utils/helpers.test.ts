import { routeToTitle } from '$lib/utils/helpers'

describe('routeToTitle', () => {
  it('should convert route to title case with spaces', () => {
    expect(routeToTitle('/persistent-volume-claims')).toBe('Persistent Volume Claims')
  })

  it('should handle routes with multiple segments', () => {
    expect(routeToTitle('/some-route/storage/persistent-volume-claims')).toBe('Persistent Volume Claims')
  })

  it('should handle routes without dashes', () => {
    expect(routeToTitle('/pods')).toBe('Pods')
  })
})
