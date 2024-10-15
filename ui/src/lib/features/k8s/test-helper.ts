// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

/* eslint-disable @typescript-eslint/no-explicit-any */

import type { ComponentType } from 'svelte'

import type { KubernetesObject } from '@kubernetes/client-node'
import { render } from '@testing-library/svelte'
import * as components from '$components'
import _ from 'lodash'
import type { Mock } from 'vitest'

import type { CommonRow, ResourceWithTable } from './types'

// Vitest type redeclared cause it's not exported from vitest
type Procedure = (...args: any[]) => any

export function testK8sTableWithDefaults(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with correct props and default columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    render(Component)

    // Ensure name and tooltip desc aren't empty
    const name: string = props.name as string
    expect(name).toBeTruthy()
    const description: string = props.description as string
    expect(description).toBeTruthy()

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

export function testK8sTableWithCustomColumns(Component: ComponentType, props: Record<string, unknown>) {
  test('creates component with custom columns', () => {
    // Access the mocked DataTable
    const { DataTable } = components

    props.columns = [['blah'], ['blah2']]

    render(Component, { columns: props.columns })

    // Check if DataTable was called
    expect(DataTable).toHaveBeenCalled()

    expect(DataTable).toHaveBeenCalledWith({
      $$inline: true,
      props,
    })
  })
}

// Helper function to compare two objects while ignoring certain fields; can ignore nested fields (eg 'metadata.creationTimestamp')
export function expectEqualIgnoringFields<T>(actual: T, expected: T, fieldsToIgnore: string[]) {
  const removeFields = (obj: T, fields: string[]) => {
    const result = _.cloneDeep(obj)
    fields.forEach((field) => _.unset(result, field))
    return result
  }

  const expectedWithoutFields = removeFields(expected, fieldsToIgnore)
  const actualWithoutFields = removeFields(actual, fieldsToIgnore)

  expect(actualWithoutFields).toEqual(expectedWithoutFields)
}

// MockResourceStore is a lightweight mock of ResourceStore for testing URLs and data transformation callbacks
export class MockResourceStore {
  url: string
  data: KubernetesObject[]
  #tableCallback: (data: KubernetesObject[]) => ResourceWithTable<KubernetesObject, CommonRow>[]

  constructor(
    url: string,
    transform: <R extends KubernetesObject, U extends CommonRow>(resources: R[]) => ResourceWithTable<R, U>[],
    data: KubernetesObject[],
  ) {
    this.url = url
    this.data = data
    this.#tableCallback = transform
  }

  // call the given transform function, imitating the store.start()
  start() {
    return this.#tableCallback(this.data)
  }

  // added to satisfy store.sortByKey.bind(store),
  sortByKey() {}
}

export class MockEventSource {
  constructor(url: string, urlAssertionMock: Mock<Procedure>) {
    // Used for testing the correct URL was passed to the EventSource
    urlAssertionMock(url)
  }
  // satisfies store.stopCallback = metricsEvents.close.bind(metricsEvents)
  close() {}
}
