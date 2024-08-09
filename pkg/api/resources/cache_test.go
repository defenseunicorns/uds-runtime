// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	autoScalingV2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	dynamicFake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
)

func TestBindCoreResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Create a mock Node
	mockNode := &corev1.Node{}
	mockNodeName := "test-node"
	mockNode.SetName(mockNodeName)
	mockNode.SetUID("123e4567-e89b-12d3-a456-426614174N0D3")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().Nodes().Create(context.Background(), mockNode, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory: informers.NewSharedInformerFactory(clientset, time.Minute*10),
		stopper: make(chan struct{}),
	}

	// Bind resources
	c.bindCoreResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.Nodes.GetResources("", mockNodeName)[0].GetName(), mockNode.Name)
}

func TestBindWorkloadResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Create a mock Pod
	mockPod := &corev1.Pod{}
	mockPodName := "test-pod"
	mockPod.SetName(mockPodName)
	mockPod.SetUID("123e4567-e89b-12d3-a456-426614174P0D")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().Pods("default").Create(context.Background(), mockPod, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory: informers.NewSharedInformerFactory(clientset, time.Minute*10),
		stopper: make(chan struct{}),
	}

	// Bind resources
	c.bindWorkloadResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.Pods.GetResources("default", mockPodName)[0].GetName(), mockPod.Name)
}

func TestBindConfigResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Create a mock Secret
	mockSecret := &corev1.Secret{}
	mockSecretName := "test-secret"
	mockSecret.SetName(mockSecretName)
	mockSecret.SetUID("123e4567-e89b-12d3-a456-426614174S3CR3T")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().Secrets("default").Create(context.Background(), mockSecret, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory: informers.NewSharedInformerFactory(clientset, time.Minute*10),
		stopper: make(chan struct{}),
	}

	// Bind resources
	c.bindConfigResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.Secrets.GetResources("default", mockSecretName)[0].GetName(), mockSecret.Name)
}

func TestBindClusterOpsResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Create a mock HPA
	mockHPA := &autoScalingV2.HorizontalPodAutoscaler{}
	mockHPAName := "test-hpa"
	mockHPA.SetName(mockHPAName)
	mockHPA.SetUID("123e4567-e89b-12d3-a456-426614174S34HP4")

	// Add the mocks to the fake clientset
	_, err := clientset.AutoscalingV2().HorizontalPodAutoscalers("default").Create(context.Background(), mockHPA, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory: informers.NewSharedInformerFactory(clientset, time.Minute*10),
		stopper: make(chan struct{}),
	}

	// Bind resources
	c.bindClusterOpsResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.HPAs.GetResources("default", mockHPAName)[0].GetName(), mockHPA.Name)
}

func TestBindNetworkResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Set up dynamic client for VirtualService which is bound in bindNetworkResources()
	runtimeScheme := runtime.NewScheme()
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1", Kind: "VirtualService"}, &unstructured.UnstructuredList{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1", Kind: "VirtualServiceList"}, &unstructured.UnstructuredList{})
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtimeScheme)

	// Create a mock service
	mockService := &corev1.Service{}
	mockServiceName := "test-service"
	mockService.SetName(mockServiceName)
	mockService.SetUID("123e4567-e89b-12d3-a456-426614174S34SVC")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().Services("default").Create(context.Background(), mockService, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory:        informers.NewSharedInformerFactory(clientset, time.Minute*10),
		dynamicFactory: dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, time.Minute*10),
		stopper:        make(chan struct{}),
	}

	// Bind resources
	c.bindNetworkResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()

	// Create a new context with a timeout for dynamic informer factory
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// start dynamic informer factory
	go func(ctx context.Context) {
		c.dynamicFactory.Start(c.stopper)
	}(ctx)

	c.dynamicFactory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.Services.GetResources("default", mockServiceName)[0].GetName(), mockService.Name)
}

func TestBindStorageResources(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// Create multiple mock persistent volumes
	mockPV := &corev1.PersistentVolume{}
	mockPV.SetName("test-pv")
	mockPV.SetUID("123e4567-e89b-12d3-a456-426614174P3RS1ST3NT")

	mockPV2 := &corev1.PersistentVolume{}
	mockPV2.SetName("test-pv2")
	mockPV2.SetUID("123e4567-e89b-12d3-a456-42661417P3RS1ST3NT2")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().PersistentVolumes().Create(context.Background(), mockPV, metav1.CreateOptions{})
	require.NoError(t, err)
	_, err = clientset.CoreV1().PersistentVolumes().Create(context.Background(), mockPV2, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory: informers.NewSharedInformerFactory(clientset, time.Minute*10),
		stopper: make(chan struct{}),
	}

	// Bind resources
	c.bindStorageResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	// Test multiple resources
	require.Equal(t, len(c.PersistentVolumes.GetResources("", "")), 2)
	pvs := c.PersistentVolumes.GetResources("", "")
	names := []string{
		pvs[0].GetName(),
		pvs[1].GetName(),
	}
	require.Contains(t, names, mockPV.Name)
	require.Contains(t, names, mockPV2.Name)
}

func TestBindUDSResources(t *testing.T) {
	// Create fake client
	runtimeScheme := runtime.NewScheme()
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "Package"}, &unstructured.Unstructured{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "PackageList"}, &unstructured.UnstructuredList{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "Exemption"}, &unstructured.Unstructured{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "ExemptionList"}, &unstructured.UnstructuredList{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1", Kind: "VirtualService"}, &unstructured.UnstructuredList{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1", Kind: "VirtualServiceList"}, &unstructured.UnstructuredList{})
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtimeScheme)

	// Create a mock UDS Package
	mockUDSPackage := &unstructured.Unstructured{}
	mockUDSPackageName := "test-uds-package"
	mockUDSPackage.SetName(mockUDSPackageName)
	mockUDSPackage.SetUID("123e4567-e89b-12d3-a456-426614174UD5P4CK4G3")

	// Add the mock UDS Package to the fake dynamic client
	udsPackageGVR := schema.GroupVersionResource{
		Group:    "uds.dev",
		Version:  "v1alpha1",
		Resource: "packages",
	}
	_, err := dynamicClient.Resource(udsPackageGVR).Create(context.Background(), mockUDSPackage, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		dynamicFactory: dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, time.Minute*10),
		stopper:        make(chan struct{}),
	}

	// Bind resources
	c.bindUDSResources()

	// Create a new context with a timeout for dynamic informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// start dynamic informer factory
	go func(ctx context.Context) {
		c.dynamicFactory.Start(c.stopper)
	}(ctx)

	c.dynamicFactory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.UDSPackages.GetResources("", mockUDSPackageName)[0].GetName(), mockUDSPackage.GetName())
}

func TestSimpleAndDynamicClient(t *testing.T) {
	// create fake client
	clientset := fake.NewSimpleClientset()

	// create fake dynamic client
	runtimeScheme := runtime.NewScheme()
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "Package"}, &unstructured.Unstructured{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "PackageList"}, &unstructured.UnstructuredList{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "Exemption"}, &unstructured.Unstructured{})
	runtimeScheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "uds.dev", Version: "v1alpha1", Kind: "ExemptionList"}, &unstructured.UnstructuredList{})
	dynamicClient := dynamicFake.NewSimpleDynamicClient(runtimeScheme)

	// Create a mock Pod
	mockPod := &corev1.Pod{}
	mockPodName := "test-pod"
	mockPod.SetName(mockPodName)
	mockPod.SetUID("123e4567-e89b-12d3-a456-426614174P0D")

	// Create a mock UDS Package
	mockUDSPackage := &unstructured.Unstructured{}
	mockUDSPackageName := "test-uds-package"
	mockUDSPackage.SetName(mockUDSPackageName)
	mockUDSPackage.SetUID("123e4567-e89b-12d3-a456-426614174UD5P4CK4G3")

	// Add the mocks to the fake clientset
	_, err := clientset.CoreV1().Pods("default").Create(context.Background(), mockPod, metav1.CreateOptions{})
	require.NoError(t, err)

	// Add the mock UDS Package to the fake dynamic client
	udsPackageGVR := schema.GroupVersionResource{
		Group:    "uds.dev",
		Version:  "v1alpha1",
		Resource: "packages",
	}
	_, err = dynamicClient.Resource(udsPackageGVR).Create(context.Background(), mockUDSPackage, metav1.CreateOptions{})
	require.NoError(t, err)

	// Create Cache instance
	c := &Cache{
		factory:        informers.NewSharedInformerFactory(clientset, time.Minute*10),
		dynamicFactory: dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, time.Minute*10),
		stopper:        make(chan struct{}),
	}

	// Bind resources
	c.bindWorkloadResources()
	c.bindUDSResources()

	// Create a new context with a timeout for informer factory
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start informer factory
	go func(ctx context.Context) {
		c.factory.Start(c.stopper)
	}(ctx)

	c.factory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()

	// Create a new context with a timeout for dynamic informer factory
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// start dynamic informer factory
	go func(ctx context.Context) {
		c.dynamicFactory.Start(c.stopper)
	}(ctx)

	c.dynamicFactory.WaitForCacheSync(c.stopper)

	// wait for the context to be done
	<-ctx.Done()
	defer close(c.stopper)

	require.Equal(t, c.Pods.GetResources("default", mockPodName)[0].GetName(), mockPod.Name)
	require.Equal(t, c.UDSPackages.GetResources("", mockUDSPackageName)[0].GetName(), mockUDSPackage.GetName())
}
