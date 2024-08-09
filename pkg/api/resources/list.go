// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package resources

import (
	"strings"
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
	Resources       map[string]*unstructured.Unstructured
	SparseResources map[string]*unstructured.Unstructured
	HasSynced       cache.InformerSynced
	Changes         chan struct{}
	gvk             schema.GroupVersionKind
}

// NewResourceList initializes a ResourceList and sets up event handlers for resource changes.
func NewResourceList(informer cache.SharedIndexInformer, gvk schema.GroupVersionKind) *ResourceList {
	r := &ResourceList{
		Resources:       make(map[string]*unstructured.Unstructured),
		SparseResources: make(map[string]*unstructured.Unstructured),
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

	// Extract spec.nodeName if it exists
	if nodeName, exists, _ := unstructured.NestedString(obj.Object, "spec", "nodeName"); exists {
		sparseObj.Object["spec"] = map[string]interface{}{
			"nodeName": nodeName,
		}
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
