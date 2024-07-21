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
	autoscalingV2 "k8s.io/api/autoscaling/v2"
	batchV1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
	storageV1 "k8s.io/api/storage/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	Events     *ResourceList[*v1.Event]
	Namespaces *ResourceList[*v1.Namespace]
	Nodes      *ResourceList[*v1.Node]

	// Workload resources
	Pods         *ResourceList[*v1.Pod]
	Deployments  *ResourceList[*appsV1.Deployment]
	Daemonsets   *ResourceList[*appsV1.DaemonSet]
	Statefulsets *ResourceList[*appsV1.StatefulSet]
	Jobs         *ResourceList[*batchV1.Job]
	CronJobs     *ResourceList[*batchV1.CronJob]

	// UDS resources
	UDSPackages   *ResourceList[*unstructured.Unstructured]
	UDSExemptions *ResourceList[*unstructured.Unstructured]

	// Config resources
	Configmaps *ResourceList[*v1.ConfigMap]
	Secrets    *ResourceList[*v1.Secret]

	// Cluster ops resources
	MutatingWebhooks   *ResourceList[*admissionRegV1.MutatingWebhookConfiguration]
	ValidatingWebhooks *ResourceList[*admissionRegV1.ValidatingWebhookConfiguration]
	HPAs               *ResourceList[*autoscalingV2.HorizontalPodAutoscaler]

	// Network resources
	Services        *ResourceList[*v1.Service]
	NetworkPolicies *ResourceList[*networkingV1.NetworkPolicy]
	Endpoints       *ResourceList[*v1.Endpoints]
	VirtualServices *ResourceList[*unstructured.Unstructured]

	// Storage resources
	PersistentVolumes      *ResourceList[*v1.PersistentVolume]
	PersistentVolumeClaims *ResourceList[*v1.PersistentVolumeClaim]
	StorageClasses         *ResourceList[*storageV1.StorageClass]

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
	c.Nodes = NewResourceList[*v1.Node](c.factory.Core().V1().Nodes().Informer())
	c.Events = NewResourceList[*v1.Event](c.factory.Core().V1().Events().Informer())
	c.Namespaces = NewResourceList[*v1.Namespace](c.factory.Core().V1().Namespaces().Informer())
}

func (c *Cache) bindWorkloadResources() {
	c.Pods = NewResourceList[*v1.Pod](c.factory.Core().V1().Pods().Informer())
	c.Deployments = NewResourceList[*appsV1.Deployment](c.factory.Apps().V1().Deployments().Informer())
	c.Daemonsets = NewResourceList[*appsV1.DaemonSet](c.factory.Apps().V1().DaemonSets().Informer())
	c.Statefulsets = NewResourceList[*appsV1.StatefulSet](c.factory.Apps().V1().StatefulSets().Informer())
	c.Jobs = NewResourceList[*batchV1.Job](c.factory.Batch().V1().Jobs().Informer())
	c.CronJobs = NewResourceList[*batchV1.CronJob](c.factory.Batch().V1().CronJobs().Informer())
}

func (c *Cache) bindConfigResources() {
	c.Configmaps = NewResourceList[*v1.ConfigMap](c.factory.Core().V1().ConfigMaps().Informer())
	c.Secrets = NewResourceList[*v1.Secret](c.factory.Core().V1().Secrets().Informer())
}

func (c *Cache) bindClusterOpsResources() {
	c.MutatingWebhooks = NewResourceList[*admissionRegV1.MutatingWebhookConfiguration](c.factory.Admissionregistration().V1().MutatingWebhookConfigurations().Informer())
	c.ValidatingWebhooks = NewResourceList[*admissionRegV1.ValidatingWebhookConfiguration](c.factory.Admissionregistration().V1().ValidatingWebhookConfigurations().Informer())
	c.HPAs = NewResourceList[*autoscalingV2.HorizontalPodAutoscaler](c.factory.Autoscaling().V2().HorizontalPodAutoscalers().Informer())
}

func (c *Cache) bindNetworkResources() {
	c.Services = NewResourceList[*v1.Service](c.factory.Core().V1().Services().Informer())
	c.NetworkPolicies = NewResourceList[*networkingV1.NetworkPolicy](c.factory.Networking().V1().NetworkPolicies().Informer())
	c.Endpoints = NewResourceList[*v1.Endpoints](c.factory.Core().V1().Endpoints().Informer())

	// VirtualServices are not part of the core informer factory
	c.VirtualServices = NewResourceList[*unstructured.Unstructured](c.dynamicFactory.ForResource(schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1",
		Resource: "virtualservices",
	}).Informer())
}

func (c *Cache) bindStorageResources() {
	c.PersistentVolumes = NewResourceList[*v1.PersistentVolume](c.factory.Core().V1().PersistentVolumes().Informer())
	c.PersistentVolumeClaims = NewResourceList[*v1.PersistentVolumeClaim](c.factory.Core().V1().PersistentVolumeClaims().Informer())
	c.StorageClasses = NewResourceList[*storageV1.StorageClass](c.factory.Storage().V1().StorageClasses().Informer())
}

func (c *Cache) bindUDSResources() {
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

	c.UDSPackages = NewResourceList[*unstructured.Unstructured](c.dynamicFactory.ForResource(udsPackageGVR).Informer())
	c.UDSExemptions = NewResourceList[*unstructured.Unstructured](c.dynamicFactory.ForResource(udsExemptionsGVR).Informer())
}
