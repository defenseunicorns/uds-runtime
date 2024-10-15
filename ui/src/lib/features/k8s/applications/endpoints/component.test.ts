// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */
import type { ZarfPackage } from '$features/k8s/applications/packages/types'
import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { beforeEach, vi } from 'vitest'

import Component from './component.svelte'
import { createStore } from './store'

suite('StatefulsetTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Application Endpoints'
  const description = 'Application Endpoints are exposed by the UDS Operator via a UDS Package CR.'

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['url', 'emphasize', 'link-external'], ['namespace'], ['status'], ['age']],
    name,
    description,
    disableRowClick: true,
  })

  testK8sTableWithCustomColumns(Component, {
    createStore,
    name,
    description,
    disableRowClick: true,
  })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'uds.dev/v1alpha1',
        kind: 'Package',
        metadata: {
          creationTimestamp: '2021-09-29T20:00:00Z',
          name: 'foo',
          namespace: 'foo',
        },
        spec: {
          network: {
            expose: [
              {
                service: 'foo',
                selector: {
                  ' "app.kubernetes.io/name"': 'foo',
                },
                gateway: 'tenant',
                host: 'foo',
                port: 9898,
              },
            ],
          },
        },
        status: {
          phase: 'Pending',
        },
      },
    ] as unknown as ZarfPackage[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'foo',
      url: {
        href: '',
        sort: '',
        text: 'Pending',
      },
      namespace: 'foo',
      status: 'Pending',
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<ZarfPackage, any>[]
  expect(store.url).toEqual('/api/v1/resources/configs/uds-packages?dense=true')
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, ['creationTimestamp'])
})
