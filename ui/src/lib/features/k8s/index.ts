// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

export { default as EventTable } from './events/component.svelte'
export { default as NamespaceTable } from './namespaces/component.svelte'
export { default as NodeTable } from './nodes/component.svelte'

export { default as CronJobTable } from './workloads/cronjobs/component.svelte'
export { default as DaemonSetsTable } from './workloads/daemonsets/component.svelte'
export { default as DeploymentTable } from './workloads/deployments/component.svelte'
export { default as JobTable } from './workloads/jobs/component.svelte'
export { default as PodTable } from './workloads/pods/component.svelte'
export { default as StatefulsetTable } from './workloads/statefulsets/component.svelte'

export { default as UDSExemptionTable } from './configs/uds-exemptions/component.svelte'
export { default as UDSPackageTable } from './configs/uds-packages/component.svelte'

export { default as ServiceTable } from './networks/services/component.svelte'
export { default as VirtualServiceTable } from './networks/virtualservices/component.svelte'
