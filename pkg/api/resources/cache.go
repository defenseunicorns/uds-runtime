// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/defenseunicorns/uds-engine/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type Handler cache.ResourceEventHandlerFuncs

type Cache struct {
	stopper        chan struct{}
	factory        informers.SharedInformerFactory
	Events         *ResourceList[*v1.Event]
	Namespaces     *ResourceList[*v1.Namespace]
	Pods           *ResourceList[*v1.Pod]
	Deployments    *ResourceList[*appsv1.Deployment]
	Daemonsets     *ResourceList[*appsv1.DaemonSet]
	Statefulsets   *ResourceList[*appsv1.StatefulSet]
	PodMetrics     *PodMetrics
	MetricsChanges chan struct{}
}

func NewCache(ctx context.Context) (*Cache, error) {
	k8s, err := k8s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the cluster: %v", err)
	}

	factory := informers.NewSharedInformerFactory(k8s.Clientset, time.Minute*10)

	c := &Cache{
		factory:        factory,
		stopper:        make(chan struct{}),
		PodMetrics:     NewPodMetrics(),
		MetricsChanges: make(chan struct{}, 1),
	}

	c.Events = NewResourceList[*v1.Event](factory.Core().V1().Events().Informer())
	c.Namespaces = NewResourceList[*v1.Namespace](factory.Core().V1().Namespaces().Informer())
	c.Pods = NewResourceList[*v1.Pod](factory.Core().V1().Pods().Informer())
	c.Deployments = NewResourceList[*appsv1.Deployment](factory.Apps().V1().Deployments().Informer())
	c.Daemonsets = NewResourceList[*appsv1.DaemonSet](factory.Apps().V1().DaemonSets().Informer())
	c.Statefulsets = NewResourceList[*appsv1.StatefulSet](factory.Apps().V1().StatefulSets().Informer())

	// start the informer
	go c.factory.Start(c.stopper)

	// Wait for the caches to sync
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
