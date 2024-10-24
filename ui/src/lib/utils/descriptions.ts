// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

// Generated by robots and the K8s docs
export const resourceDescriptions: { [key: string]: string } = {
  ClusterRoleBindings:
    'A ClusterRoleBinding grants the permissions defined in a ClusterRole to a user or set of users across the entire cluster.',
  ClusterRoles:
    'A ClusterRole is a logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.',
  ConfigMaps: 'A ConfigMap is an API object used to store non-confidential data in key-value pairs.',
  CronJobs: 'A CronJob creates Jobs on a repeating schedule.',
  CustomResourceDefinitions:
    'A CustomResourceDefinition (CRD) allows you to define your own resource types in the Kubernetes API.',
  DaemonSets: 'A DaemonSet ensures that all (or some) Nodes run a copy of a Pod.',
  Deployments: 'A Deployment provides declarative updates for Pods and ReplicaSets.',
  EndpointSlices:
    'An EndpointSlice is a more scalable way to handle endpoints in large clusters. It provides a simple way to track network endpoints.',
  Endpoints:
    'An Endpoint is a collection of endpoints that implement the actual service. An endpoint is typically a set of IP addresses and ports.',
  Events: 'An Event is a report of an event that has occurred in the cluster.',
  Exemptions: 'A UDS Exemption allows you to exempt a resource from UDS Core policies.',
  HorizontalPodAutoscalers:
    'A HorizontalPodAutoscaler automatically scales the number of pods in a replication controller, deployment, replica set or stateful set based on observed CPU utilization or other select metrics.',
  Ingresses: 'An Ingress is an API object that manages external access to the services in a cluster, typically HTTP.',
  Jobs: 'A Job creates one or more Pods and ensures that a specified number of them successfully terminate.',
  LimitRanges: 'A LimitRange is a policy to constrain resource allocations (to Pods or Containers) in a Namespace.',
  MutatingWebhookConfigurations: 'Configures when and how to call a Mutating Webhook.',
  Namespaces: 'A Namespace is a way to divide cluster resources between multiple users.',
  NetworkPolicies:
    'A NetworkPolicy is an API object that allows you to control the traffic flow at the IP address or port level.',
  Nodes: 'A Node is a worker machine in Kubernetes, previously known as a minion.',
  Packages:
    'A UDS Package allows you to define a set of network policies, ingresses, metrics targets and SSO configurations for a workload.',
  PersistentVolumeClaims: 'A PersistentVolumeClaim (PVC) is a request for storage by a user. It is similar to a Pod.',
  PersistentVolumes:
    'A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes.',
  PodDisruptionBudgets:
    'A PodDisruptionBudget is an API object that specifies the minimum number or percentage of replicas that must be up at a time for the application to continue to function.',
  Pods: 'A Pod is the smallest and simplest Kubernetes object. A Pod represents a set of running containers on your cluster.',
  PriorityClasses:
    'A PriorityClass defines the scheduling priority for Pods. PriorityClass names are unique within a cluster.',
  ReplicaSets: 'A ReplicaSet ensures that a specified number of pod replicas are running at any given time.',
  ReplicationControllers:
    'A ReplicationController ensures that a specified number of pod replicas are running at any given time.',
  ResourceQuotas: 'A ResourceQuota provides constraints that limit aggregate resource consumption per Namespace.',
  RoleBindings: 'A RoleBinding grants the permissions defined in a Role to a user or set of users.',
  Roles: 'A Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.',
  RuntimeClasses: 'A RuntimeClass is a Kubernetes resource for defining different runtimes for Pods.',
  Secrets:
    'A Secret is an object that contains a small amount of sensitive data such as a password, a token, or a key.',
  ServiceAccountTokens: 'A ServiceAccountToken is a JWT that identifies a service account.',
  ServiceAccounts: 'A ServiceAccount provides an identity for processes that run in a Pod.',
  Services: 'A Service is an abstraction which defines a logical set of Pods and a policy by which to access them.',
  StatefulSets: 'A StatefulSet is the workload API object used to manage stateful applications.',
  StorageClasses: "A StorageClass provides a way for administrators to describe the 'classes' of storage they offer.",
  ValidatingWebhookConfigurations: 'Configures when and how to call a Validating Webhook.',
  VerticalPodAutoscalers:
    'A VerticalPodAutoscaler automatically sets resource requests and limits of a pod based on historical data.',
  VirtualServices: 'A VirtualService defines a set of traffic routing rules to apply when a host is addressed.',
  VolumeAttachments:
    'A VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.',
  'Zarf Packages': 'Artifact containing a set of components and configurations for air-gapped workloads.',
}
