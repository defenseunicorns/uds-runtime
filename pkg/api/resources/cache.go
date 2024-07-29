// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/k8s"

	admissionRegV1 "k8s.io/api/admissionregistration/v1"
	appsV1 "k8s.io/api/apps/v1"
	autoScalingV2 "k8s.io/api/autoscaling/v2"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	nodeV1 "k8s.io/api/node/v1"
	policyV1 "k8s.io/api/policy/v1"
	schedulingV1 "k8s.io/api/scheduling/v1"
	storageV1 "k8s.io/api/storage/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	dynamicInformer "k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type Handler cache.ResourceEventHandlerFuncs

type Cache struct {
	stopper        chan struct{}
	factory        informers.SharedInformerFactory
	dynamicFactory dynamicInformer.DynamicSharedInformerFactory

	// Core resources
	Events     *ResourceList
	Namespaces *ResourceList
	Nodes      *ResourceList

	// Workload resources
	Pods         *ResourceList
	Deployments  *ResourceList
	Daemonsets   *ResourceList
	Statefulsets *ResourceList
	Jobs         *ResourceList
	CronJobs     *ResourceList

	// UDS resources
	UDSPackages   *ResourceList
	UDSExemptions *ResourceList

	// Config resources
	Configmaps *ResourceList
	Secrets    *ResourceList

	// Cluster ops resources
	MutatingWebhooks     *ResourceList
	ValidatingWebhooks   *ResourceList
	HPAs                 *ResourceList
	PriorityClasses      *ResourceList
	RuntimeClasses       *ResourceList
	PodDisruptionBudgets *ResourceList
	LimitRanges          *ResourceList
	ResourceQuotas       *ResourceList

	// Network resources
	Services        *ResourceList
	NetworkPolicies *ResourceList
	Endpoints       *ResourceList
	VirtualServices *ResourceList

	// Storage resources
	PersistentVolumes      *ResourceList
	PersistentVolumeClaims *ResourceList
	StorageClasses         *ResourceList

	// Metrics
	PodMetrics     *PodMetrics
	MetricsChanges chan struct{}
}

func NewCache(ctx context.Context) (*Cache, error) {
	k8s, err := k8s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the cluster: %v", err)
	}

	c := &Cache{
		factory:        informers.NewSharedInformerFactory(k8s.Clientset, time.Minute*10),
		stopper:        make(chan struct{}),
		PodMetrics:     NewPodMetrics(),
		MetricsChanges: make(chan struct{}, 1),
	}

	// Create the dynamic client and factory
	dynamicClient, err := dynamic.NewForConfig(k8s.Config)
	if err != nil {
		return nil, fmt.Errorf("unable to create dynamic client: %v", err)
	}
	c.dynamicFactory = dynamicInformer.NewFilteredDynamicSharedInformerFactory(dynamicClient, time.Minute*10, metaV1.NamespaceAll, nil)

	c.bindCoreResources()
	c.bindWorkloadResources()
	c.bindUDSResources()
	c.bindConfigResources()
	c.bindClusterOpsResources()
	c.bindNetworkResources()
	c.bindStorageResources()

	// start the informer
	go c.factory.Start(c.stopper)
	go c.dynamicFactory.Start(c.stopper)

	// Wait for the pod cache to sync as they it is required for metrics collection
	if !cache.WaitForCacheSync(ctx.Done(), c.Pods.HasSynced) {
		return nil, fmt.Errorf("timed out waiting for caches to sync")
	}

	// Start metrics collection
	go c.StartMetricsCollection(ctx, k8s.MetricsClient)

	// Stop the informer when the context is done
	go func() {
		<-ctx.Done()
		close(c.stopper)
	}()

	return c, nil
}

func (c *Cache) bindCoreResources() {
	nodeGVK := coreV1.SchemeGroupVersion.WithKind("Node")
	eventGVK := coreV1.SchemeGroupVersion.WithKind("Event")
	namespaceGVK := coreV1.SchemeGroupVersion.WithKind("Namespace")

	c.Nodes = NewResourceList(c.factory.Core().V1().Nodes().Informer(), nodeGVK)
	c.Events = NewResourceList(c.factory.Core().V1().Events().Informer(), eventGVK)
	c.Namespaces = NewResourceList(c.factory.Core().V1().Namespaces().Informer(), namespaceGVK)
}

func (c *Cache) bindWorkloadResources() {
	podGVK := coreV1.SchemeGroupVersion.WithKind("Pod")
	deploymentGVK := appsV1.SchemeGroupVersion.WithKind("Deployment")
	daemonsetGVK := appsV1.SchemeGroupVersion.WithKind("DaemonSet")
	statefulsetGVK := appsV1.SchemeGroupVersion.WithKind("StatefulSet")
	jobGVK := batchV1.SchemeGroupVersion.WithKind("Job")
	cronJobGVK := batchV1.SchemeGroupVersion.WithKind("CronJob")

	c.Pods = NewResourceList(c.factory.Core().V1().Pods().Informer(), podGVK)
	c.Deployments = NewResourceList(c.factory.Apps().V1().Deployments().Informer(), deploymentGVK)
	c.Daemonsets = NewResourceList(c.factory.Apps().V1().DaemonSets().Informer(), daemonsetGVK)
	c.Statefulsets = NewResourceList(c.factory.Apps().V1().StatefulSets().Informer(), statefulsetGVK)
	c.Jobs = NewResourceList(c.factory.Batch().V1().Jobs().Informer(), jobGVK)
	c.CronJobs = NewResourceList(c.factory.Batch().V1().CronJobs().Informer(), cronJobGVK)
}

func (c *Cache) bindConfigResources() {
	configMapGVK := coreV1.SchemeGroupVersion.WithKind("ConfigMap")
	secretGVK := coreV1.SchemeGroupVersion.WithKind("Secret")

	c.Configmaps = NewResourceList(c.factory.Core().V1().ConfigMaps().Informer(), configMapGVK)
	c.Secrets = NewResourceList(c.factory.Core().V1().Secrets().Informer(), secretGVK)
}

func (c *Cache) bindClusterOpsResources() {
	mutatingWebhookGVK := admissionRegV1.SchemeGroupVersion.WithKind("MutatingWebhookConfiguration")
	validatingWebhookGVK := admissionRegV1.SchemeGroupVersion.WithKind("ValidatingWebhookConfiguration")
	hpaGVK := autoScalingV2.SchemeGroupVersion.WithKind("HorizontalPodAutoscaler")
	runtimeClassGVK := nodeV1.SchemeGroupVersion.WithKind("RuntimeClass")
	priorityClassGVK := schedulingV1.SchemeGroupVersion.WithKind("PriorityClass")
	podDisruptionBudgetGVK := policyV1.SchemeGroupVersion.WithKind("PodDisruptionBudget")
	limitRangesGVK := coreV1.SchemeGroupVersion.WithKind("LimitRange")
	resourceQuotaGVK := coreV1.SchemeGroupVersion.WithKind("ResourceQuotas")

	c.MutatingWebhooks = NewResourceList(c.factory.Admissionregistration().V1().MutatingWebhookConfigurations().Informer(), mutatingWebhookGVK)
	c.ValidatingWebhooks = NewResourceList(c.factory.Admissionregistration().V1().ValidatingWebhookConfigurations().Informer(), validatingWebhookGVK)
	c.HPAs = NewResourceList(c.factory.Autoscaling().V2().HorizontalPodAutoscalers().Informer(), hpaGVK)
	c.RuntimeClasses = NewResourceList(c.factory.Node().V1().RuntimeClasses().Informer(), runtimeClassGVK)
	c.PriorityClasses = NewResourceList(c.factory.Scheduling().V1().PriorityClasses().Informer(), priorityClassGVK)
	c.PodDisruptionBudgets = NewResourceList(c.factory.Policy().V1().PodDisruptionBudgets().Informer(), podDisruptionBudgetGVK)
	c.LimitRanges = NewResourceList(c.factory.Core().V1().LimitRanges().Informer(), limitRangesGVK)
	c.ResourceQuotas = NewResourceList(c.factory.Core().V1().ResourceQuotas().Informer(), resourceQuotaGVK)
}

func (c *Cache) bindNetworkResources() {
	serviceGVK := coreV1.SchemeGroupVersion.WithKind("Service")
	networkPolicyGVK := coreV1.SchemeGroupVersion.WithKind("NetworkPolicy")
	endpointGVK := coreV1.SchemeGroupVersion.WithKind("Endpoints")
	isitoVSGVK := schema.FromAPIVersionAndKind("networking.istio.io/v1", "VirtualService")

	c.Services = NewResourceList(c.factory.Core().V1().Services().Informer(), serviceGVK)
	c.NetworkPolicies = NewResourceList(c.factory.Networking().V1().NetworkPolicies().Informer(), networkPolicyGVK)
	c.Endpoints = NewResourceList(c.factory.Core().V1().Endpoints().Informer(), endpointGVK)

	// VirtualServices are not part of the core informer factory
	istioVSGVR := schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1",
		Resource: "virtualservices",
	}

	c.VirtualServices = NewResourceList(c.dynamicFactory.ForResource(istioVSGVR).Informer(), isitoVSGVK)
}

func (c *Cache) bindStorageResources() {
	persistentVolumeGVK := coreV1.SchemeGroupVersion.WithKind("PersistentVolume")
	persistentVolumeClaimGVK := coreV1.SchemeGroupVersion.WithKind("PersistentVolumeClaim")
	storageClassGVK := storageV1.SchemeGroupVersion.WithKind("StorageClass")

	c.PersistentVolumes = NewResourceList(c.factory.Core().V1().PersistentVolumes().Informer(), persistentVolumeGVK)
	c.PersistentVolumeClaims = NewResourceList(c.factory.Core().V1().PersistentVolumeClaims().Informer(), persistentVolumeClaimGVK)
	c.StorageClasses = NewResourceList(c.factory.Storage().V1().StorageClasses().Informer(), storageClassGVK)
}

func (c *Cache) bindUDSResources() {
	udsPackageGVK := schema.FromAPIVersionAndKind("uds.dev/v1alpha1", "Package")
	udsExemptionGVK := schema.FromAPIVersionAndKind("uds.dev/v1alpha1", "Exemption")

	udsPackageGVR := schema.GroupVersionResource{
		Group:    "uds.dev",
		Version:  "v1alpha1",
		Resource: "packages",
	}

	udsExemptionsGVR := schema.GroupVersionResource{
		Group:    "uds.dev",
		Version:  "v1alpha1",
		Resource: "exemptions",
	}

	c.UDSPackages = NewResourceList(c.dynamicFactory.ForResource(udsPackageGVR).Informer(), udsPackageGVK)
	c.UDSExemptions = NewResourceList(c.dynamicFactory.ForResource(udsExemptionsGVR).Informer(), udsExemptionGVK)
}
