// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

import {
  expectEqualIgnoringFields,
  TestCreationTimestamp,
  testK8sResourceStore,
  testK8sTableWithCustomColumns,
  testK8sTableWithDefaults,
} from '$features/k8s/test-helper'
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

  testK8sTableWithDefaults(Component, {
    createStore,
    columns: [['name', 'emphasize'], ['namespace'], ['details'], ['matcher'], ['policies'], ['age']],
  })

  testK8sTableWithCustomColumns(Component, { createStore })

  const mockData = [
    {
      apiVersion: 'uds.dev/v1alpha1',
      kind: 'Exemption',
      metadata: {
        creationTimestamp: TestCreationTimestamp,
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

  const expectedTables = [
    {
      name: mockData[0].metadata!.name,
      namespace: mockData[0].metadata?.namespace,
      title: mockData[0].spec?.exemptions[0].title,
      details: {
        component: vi.fn(),
        sort: 'neuvector-enforcer-pod',
        props: {
          exemption: { ...mockData[0].spec?.exemptions[0], resource: mockData[0] },
        },
      },
      matcher: {
        component: vi.fn(),
        props: { matcher: mockData[0].spec?.exemptions[0].matcher },
      },
      policies: {
        component: vi.fn(),
        props: { policies: mockData[0].spec?.exemptions[0].policies },
      },
      age: { text: '1 minute', sort: 1721923882000 },
    },
  ]

  //eslint-disable-next-line @typescript-eslint/no-explicit-any
  const subscribeCallback = (data: any[]) => {
    const { resource, table } = data[0]

    // Assert the data was passed from eventSource to transformer (avoid date time inconsistencies by ignoring creationTimestamp)
    expectEqualIgnoringFields({ ...resource }, { ...mockData[0] }, ['metadata.creationTimestamp'])
    // Assert the data was transformed correctly to create the desired table rows
    expectEqualIgnoringFields(table, expectedTables[0], [
      'details.component',
      'matcher.component',
      'policies.component',
      'creationTimestamp',
    ])
  }

  testK8sResourceStore(
    'UDSExemptions',
    mockData,
    expectedTables,
    `/api/v1/resources/configs/uds-exemptions?dense=true`,
    createStore,
    subscribeCallback,
  )
})
