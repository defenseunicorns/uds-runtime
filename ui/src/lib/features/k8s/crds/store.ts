import type { V1CustomResourceDefinition as Resource } from '@kubernetes/client-node'
import { ResourceStore, transformResource } from '$features/k8s/store'
import type { ColumnWrapper, CommonRow, ResourceStoreInterface } from '$features/k8s/types'

interface Row extends CommonRow {
  group: string
  kind: string
  versions: string
  scope: string
}

export type Columns = ColumnWrapper<Row>

export function createStore(): ResourceStoreInterface<Resource, Row> {
  const url = `/api/v1/resources/custom-resource-definitions?fields=metadata,spec.group,spec.names.kind,spec.versions[].name,spec.scope`

  const transform = transformResource<Resource, Row>((r) => {
    return {
      group: r.spec?.group,
      kind: r.spec?.names.kind ?? '',
      versions: r.spec?.versions?.map((v) => v.name).join(', ') ?? '',
      scope: r.spec?.scope,
    }
  })

  const store = new ResourceStore<Resource, Row>(url, transform, 'namespace')

  return {
    ...store,
    start: store.start.bind(store),
    sortByKey: store.sortByKey.bind(store),
  }
}
