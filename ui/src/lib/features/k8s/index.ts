// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

// Core Resources
export { default as EventTable } from './events/component.svelte'
export { default as NamespaceTable } from './namespaces/component.svelte'
export { default as NodeTable } from './nodes/component.svelte'

// Workload resources
export { default as CronJobTable } from './workloads/cronjobs/component.svelte'
export { default as DaemonSetsTable } from './workloads/daemonsets/component.svelte'
export { default as DeploymentTable } from './workloads/deployments/component.svelte'
export { default as JobTable } from './workloads/jobs/component.svelte'
export { default as PodTable } from './workloads/pods/component.svelte'
export { default as StatefulsetTable } from './workloads/statefulsets/component.svelte'

// Config resources
export { default as ConfigMapTable } from './configs/configmaps/component.svelte'
export { default as SecretTable } from './configs/secrets/component.svelte'
export { default as UDSExemptionTable } from './configs/uds-exemptions/component.svelte'
export { default as UDSPackageTable } from './configs/uds-packages/component.svelte'

// Cluster ops resources
export { default as LimitRangesTable } from './cluster-ops/limit-ranges/component.svelte'
export { default as MutatingWebhooksTable } from './cluster-ops/mutatingwebhooks/component.svelte'
export { default as PodDisruptionBudgetsTable } from './cluster-ops/pod-disruption-budgets/component.svelte'
export { default as PriorityClassesTable } from './cluster-ops/priority-classes/component.svelte'
export { default as RuntimeClassesTable } from './cluster-ops/runtime-classes/component.svelte'

// Network resources
export { default as EndpointTable } from './networks/endpoints/component.svelte'
export { default as NetworkPolicyTable } from './networks/network-policies/component.svelte'
export { default as ServiceTable } from './networks/services/component.svelte'
export { default as VirtualServiceTable } from './networks/virtualservices/component.svelte'

// Storage resources
export { default as PersistentVolumeClaimTable } from './storage/persistentvolumeclaims/component.svelte'
export { default as PersistentVolumeTable } from './storage/persistentvolumes/component.svelte'
export { default as StorageClassesTable } from './storage/storageclasses/component.svelte'
