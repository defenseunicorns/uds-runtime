// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"sync"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	resources       map[string]*unstructured.Unstructured
	sparseResources map[string]*unstructured.Unstructured
	HasSynced       cache.InformerSynced
	Changes         chan struct{}
	gvk             schema.GroupVersionKind
}

// NewResourceList initializes a ResourceList and sets up event handlers for resource changes.
func NewResourceList(informer cache.SharedIndexInformer, gvk schema.GroupVersionKind) *ResourceList {
	r := &ResourceList{
		resources:       make(map[string]*unstructured.Unstructured),
		sparseResources: make(map[string]*unstructured.Unstructured),
		Changes:         make(chan struct{}, 1),
		HasSynced:       informer.HasSynced,
		gvk:             gvk,
	}

	//nolint:errcheck
	// Handlers to update the ResourceList
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			r.notifyChange(obj, Added)
		},
		//nolint:revive
		UpdateFunc: func(oldObj, newObj any) {
			r.notifyChange(newObj, Modified)
		},
		DeleteFunc: func(obj any) {
			r.notifyChange(obj, Deleted)
		},
	})

	return r
}

// GetResource returns a resource by UID.
func (r *ResourceList) GetResource(uid string) (unstructured.Unstructured, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	resource, found := r.resources[uid]
	return *resource, found
}

// GetResources returns a slice of the current resources.
func (r *ResourceList) GetResources() []unstructured.Unstructured {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	resources := make([]unstructured.Unstructured, 0, len(r.resources))
	for _, resource := range r.resources {
		resources = append(resources, *resource)
	}

	return resources
}

// GetSparseResources returns a slice of the current resources with only metadata and status fields.
func (r *ResourceList) GetSparseResources() []unstructured.Unstructured {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	resources := make([]unstructured.Unstructured, 0, len(r.sparseResources))
	for _, resource := range r.sparseResources {
		resources = append(resources, *resource)
	}

	return resources
}

// notifyChange updates the ResourceList based on the event type and notifies subscribers of changes.
func (r *ResourceList) notifyChange(obj interface{}, eventType string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	resource, err := toUnstructured(obj)
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

	// Extract the sparse object
	sparseResource := r.extractSparseObject(resource)

	// Update the resource list based on the event type
	switch eventType {
	case Added, Modified:
		r.resources[uid] = resource
		r.sparseResources[uid] = sparseResource
	case Deleted:
		delete(r.resources, uid)
		delete(r.sparseResources, uid)
	}

	// Notify subscribers of the change
	select {
	case r.Changes <- struct{}{}:
	default:
	}
}

// extractSparseObject creates a sparse representation of the given unstructured object
func (r *ResourceList) extractSparseObject(obj *unstructured.Unstructured) *unstructured.Unstructured {
	sparseObj := &unstructured.Unstructured{
		Object: map[string]interface{}{},
	}

	// Safely extract apiVersion
	if apiVersion, exists := obj.Object["apiVersion"]; exists {
		sparseObj.Object["apiVersion"] = apiVersion
	}

	// Safely extract kind
	if kind, exists := obj.Object["kind"]; exists {
		sparseObj.Object["kind"] = kind
	}

	// Extract metadata
	if metadata, exists := obj.Object["metadata"]; exists {
		sparseObj.Object["metadata"] = metadata
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

	// Strip the metadata managed fields
	delete(sparseObj.Object["metadata"].(map[string]interface{}), "managedFields")

	return sparseObj
}
