// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package resources

import (
	"strings"
	"sync"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

const (
	Added    = "ADDED"
	Modified = "MODIFIED"
	Deleted  = "DELETED"
)

// ResourceList is a thread-safe struct to store the list of resources and notify subscribers of changes.
type ResourceList struct {
	mutex           sync.RWMutex
	Resources       map[string]*unstructured.Unstructured
	SparseResources map[string]*unstructured.Unstructured
	HasSynced       cache.InformerSynced
	Changes         chan struct{}
	gvk             schema.GroupVersionKind
	GVR             schema.GroupVersionResource
	CRDExists       bool
}

// initializeResourceList initializes the common fields of ResourceList.
func initializeResourceList(informer cache.SharedIndexInformer, gvk schema.GroupVersionKind) *ResourceList {
	r := &ResourceList{
		Resources:       make(map[string]*unstructured.Unstructured),
		SparseResources: make(map[string]*unstructured.Unstructured),
		Changes:         make(chan struct{}, 1),
		HasSynced:       informer.HasSynced,
		gvk:             gvk,
		CRDExists:       true,
		GVR:             schema.GroupVersionResource{},
	}

	//nolint:errcheck
	// Handlers to update the ResourceList
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			r.notifyChange(obj, Added)
		},
		UpdateFunc: func(_, newObj any) {
			r.notifyChange(newObj, Modified)
		},
		DeleteFunc: func(obj any) {
			r.notifyChange(obj, Deleted)
		},
	})

	return r
}

// NewResourceList initializes a ResourceList and sets up event handlers for resource changes.
func NewResourceList(informer cache.SharedIndexInformer, gvk schema.GroupVersionKind) *ResourceList {
	r := initializeResourceList(informer, gvk)
	return r
}

// NewDynamicResourceList initializes a ResourceList with a gvr.
func NewDynamicResourceList(informer cache.SharedIndexInformer, gvk schema.GroupVersionKind, gvr schema.GroupVersionResource) *ResourceList {
	r := initializeResourceList(informer, gvk)
	r.GVR = gvr
	return r
}

// GetResource returns a resource by UID.
func (r *ResourceList) GetResource(uid string) (unstructured.Unstructured, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	resource, found := r.Resources[uid]
	if !found {
		return unstructured.Unstructured{}, false
	}

	return *resource, true
}

// GetResources returns a slice of the current resources.
func (r *ResourceList) GetResources(namespace string, namePartial string) []unstructured.Unstructured {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	resources := make([]unstructured.Unstructured, 0, len(r.Resources))
	for _, resource := range r.Resources {
		// Check if the resource matches the namespace and name filter
		if r.isFilterMatch(resource, namespace, namePartial) {
			resources = append(resources, *resource)
		}
	}

	return resources
}

// GetSparseResources returns a slice of the current resources with only metadata and status fields.
func (r *ResourceList) GetSparseResources(namespace string, namePartial string) []unstructured.Unstructured {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	resources := make([]unstructured.Unstructured, 0, len(r.SparseResources))
	for _, resource := range r.SparseResources {
		// Check if the resource matches the namespace and name filter
		if r.isFilterMatch(resource, namespace, namePartial) {
			resources = append(resources, *resource)
		}
	}

	return resources
}

// CRDExistsInCluster returns the value of the MissingCRD field for the ResourceList.
func (r *ResourceList) CRDExistsInCluster() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.CRDExists
}

// notifyChange updates the ResourceList based on the event type and notifies subscribers of changes.
func (r *ResourceList) notifyChange(obj interface{}, eventType string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	resource, err := ToUnstructured(obj)
	if err != nil {
		// Handle error or log it
		return
	}

	// Add GVK because they wont exist without the typed informer
	resource.SetGroupVersionKind(r.gvk)

	// Extract UID
	uid := string(resource.GetUID())
	if uid == "" {
		// Handle error: UID is required
		return
	}

	// Always remove managedFields from the resource
	delete(resource.Object["metadata"].(map[string]interface{}), "managedFields")

	// Extract the sparse object
	sparseResource := r.extractSparseObject(resource)

	// Update the resource list based on the event type
	switch eventType {
	case Added, Modified:
		r.Resources[uid] = resource
		r.SparseResources[uid] = sparseResource
	case Deleted:
		delete(r.Resources, uid)
		delete(r.SparseResources, uid)
	}

	// Notify subscribers of the change
	select {
	case r.Changes <- struct{}{}:
	default:
	}
}

// isFilterMatch checks if the resource matches the namespace and name filter
func (r *ResourceList) isFilterMatch(resource *unstructured.Unstructured, namespace string, namePartial string) bool {
	if namespace != "" && resource.GetNamespace() != namespace {
		return false
	}

	if namePartial != "" && !strings.Contains(resource.GetName(), namePartial) {
		return false
	}

	return true
}

// extractSparseObject creates a sparse representation of the given unstructured object
func (r *ResourceList) extractSparseObject(obj *unstructured.Unstructured) *unstructured.Unstructured {
	sparseObj := &unstructured.Unstructured{
		Object: map[string]interface{}{},
	}

	// Safely extract apiVersion (GetAPIVersion() may not work)
	if apiVersion, exists := obj.Object["apiVersion"]; exists {
		sparseObj.Object["apiVersion"] = apiVersion
	}

	// Safely extract kind (GetKind() may not work)
	if kind, exists := obj.Object["kind"]; exists {
		sparseObj.Object["kind"] = kind
	}

	// Extract metadata and deep copy it to avoid mutating the original object
	if metadata, ok := obj.Object["metadata"].(map[string]interface{}); ok {
		sparseObj.Object["metadata"] = runtime.DeepCopyJSON(metadata)
	}

	// Extract type if it exists
	if typeStr, exists, _ := unstructured.NestedString(obj.Object, "type"); exists {
		sparseObj.Object["type"] = typeStr
	}

	// Extract data if it exists, but only preserve the keys
	if data, exists, _ := unstructured.NestedFieldCopy(obj.Object, "data"); exists {
		sparseObj.Object["data"] = map[string]interface{}{}
		for key := range data.(map[string]interface{}) {
			sparseObj.Object["data"].(map[string]interface{})[key] = nil
		}
	}

	// Include status if it exists
	if status, exists, _ := unstructured.NestedFieldCopy(obj.Object, "status"); exists {
		sparseObj.Object["status"] = status
	} else {
		// If status doesn't exist, still include it as nil
		sparseObj.Object["status"] = nil
	}

	// Strip the metadata annotations from the copy
	delete(sparseObj.Object["metadata"].(map[string]interface{}), "annotations")

	return sparseObj
}
