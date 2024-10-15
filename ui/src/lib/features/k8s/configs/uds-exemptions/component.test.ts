// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

/* eslint-disable @typescript-eslint/no-explicit-any */
import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  MockResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
import type { ResourceWithTable } from '$features/k8s/types'
import { resourceDescriptions } from '$lib/utils/descriptions'
import type { Exemption } from 'uds-core-types/src/pepr/operator/crd/generated/exemption-v1alpha1'

import Component from './component.svelte'
import { createStore } from './store'

// @todo: had to mock these components because they were causing major
// slow-downs in the transform process for vitest
vi.mock('./exemption-details/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

vi.mock('./exemption-matcher/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

vi.mock('./exemption-policies/component.svelte', () => ({
  default: vi.fn().mockImplementation(() => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })),
}))

suite('UDSExemptionTable Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  const name = 'Exemptions'
  const description = resourceDescriptions[name]

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [
      ['name', 'emphasize w-2/12'],
      ['namespace', 'w-3/12'],
      ['details', 'w-1/12'],
      ['matcher', 'w-2/12'],
      ['policies', 'w-3/12'],
      ['age', 'w-1/12'],
    ],
    name,
    description,
  })

  testK8sTableWithCustomColumns(Component, { createStore, name, description })

  vi.mock('../../store.ts', async (importOriginal) => {
    const mockData = [
      {
        apiVersion: 'uds.dev/v1alpha1',
        kind: 'Exemption',
        metadata: {
          creationTimestamp: '2024-09-29T20:00:00Z',
          name: 'neuvector',
          namespace: 'uds-policy-exemptions',
        },
        spec: {
          exemptions: [
            {
              description:
                "Neuvector requires HostPath volume types Neuvector mounts the following hostPaths: `/var/neuvector`: (as writable) for Neuvector's buffering and persistent state `/var/run`: communication to docker daemon `/proc`: monitoring of processes for malicious activity `/sys/fs/cgroup`: important files the controller wants to monitor for malicious content https://github.com/neuvector/neuvector-helm/blob/master/charts/core/templates/enforcer-daemonset.yaml#L108",
              matcher: { kind: 'pod', name: '^neuvector-enforcer-pod.*', namespace: 'neuvector' },
              policies: [
                'DisallowHostNamespaces',
                'DisallowPrivileged',
                'DropAllCapabilities',
                'RequireNonRootUser',
                'RestrictHostPathWrite',
                'RestrictVolumeTypes',
              ],
              title: 'neuvector-enforcer-pod',
            },
          ],
        },
      },
    ] as unknown as Exemption[]

    const original: Record<string, unknown> = await importOriginal()
    return {
      ...original,
      ResourceStore: vi.fn().mockImplementation((url, transform) => new MockResourceStore(url, transform, mockData)),
    }
  })

  const expectedTables = [
    {
      name: 'neuvector',
      namespace: 'uds-policy-exemptions',
      title: 'neuvector-enforcer-pod',
      details: {
        component: vi.fn(),
        sort: 'neuvector-enforcer-pod',
        props: {
          exemption: {
            resource: {},
            title: 'neuvector-enforcer-pod',
            description:
              "Neuvector requires HostPath volume types Neuvector mounts the following hostPaths: `/var/neuvector`: (as writable) for Neuvector's buffering and persistent state `/var/run`: communication to docker daemon `/proc`: monitoring of processes for malicious activity `/sys/fs/cgroup`: important files the controller wants to monitor for malicious content https://github.com/neuvector/neuvector-helm/blob/master/charts/core/templates/enforcer-daemonset.yaml#L108",
            matcher: {
              kind: 'pod',
              name: '^neuvector-enforcer-pod.*',
              namespace: 'neuvector',
            },
            policies: [
              'DisallowHostNamespaces',
              'DisallowPrivileged',
              'DropAllCapabilities',
              'RequireNonRootUser',
              'RestrictHostPathWrite',
              'RestrictVolumeTypes',
            ],
          },
        },
      },
      matcher: {
        component: vi.fn(),
        props: {
          matcher: {
            kind: 'pod',
            name: '^neuvector-enforcer-pod.*',
            namespace: 'neuvector',
          },
        },
      },
      policies: {
        list: [
          'DisallowHostNamespaces',
          'DisallowPrivileged',
          'DropAllCapabilities',
          'RequireNonRootUser',
          'RestrictHostPathWrite',
          'RestrictVolumeTypes',
        ],
      },
    },
  ]

  const store = createStore()
  const start = store.start as unknown as () => ResourceWithTable<Exemption, any>[]
  expect(store.url).toEqual(`/api/v1/resources/configs/uds-exemptions?dense=true`)
  console.log(start()[0].table)
  // ignore creationTimestamp because age is not calculated at this point and added to the table
  expectEqualIgnoringFields(start()[0].table, expectedTables[0] as unknown, [
    'creationTimestamp',
    'details.props.exemption.resource',
    'details.component',
    'matcher.component',
    'policies.component',
  ])
})
