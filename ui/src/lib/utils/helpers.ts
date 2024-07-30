export const stringToSnakeCase = (name: string) => name.split(' ').join('-').toLocaleLowerCase()
export const routeToTitle = (route: string): string =>
  (route.split('/').pop() || '').replace(/-/g, ' ').replace(/\b\w/g, (char) => char.toUpperCase())
