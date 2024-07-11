package resources

import (
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

const (
	Added    = "ADDED"
	Modified = "MODIFIED"
	Deleted  = "DELETED"
)

// ResourceList is a thread-safe struct to store the list of resources and notify subscribers of changes.
type ResourceList[T metav1.Object] struct {
	mutex     sync.RWMutex
	resources map[string]T
	HasSynced cache.InformerSynced
	Changes   chan struct{}
}

// NewResourceList initializes a ResourceList and sets up event handlers for resource changes.
func NewResourceList[T metav1.Object](informer cache.SharedIndexInformer) *ResourceList[T] {
	r := &ResourceList[T]{
		resources: make(map[string]T),
		HasSynced: informer.HasSynced,
		Changes:   make(chan struct{}, 1),
	}

	// Handlers to update the ResourceList
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			r.notifyChange(obj, Added)
		},
		UpdateFunc: func(oldObj, newObj any) {
			r.notifyChange(newObj, Modified)
		},
		DeleteFunc: func(obj any) {
			r.notifyChange(obj, Deleted)
		},
	})

	return r
}

// GetResources returns a slice of the current resources.
func (r *ResourceList[T]) GetResources() []T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	resources := make([]T, 0, len(r.resources))
	for _, resource := range r.resources {
		resources = append(resources, resource)
	}
	return resources
}

// notifyChange updates the ResourceList based on the event type and notifies subscribers of changes.
func (r *ResourceList[T]) notifyChange(obj interface{}, eventType string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var resource T
	var uid string

	// Try to cast directly to T first, as this will be the most common case
	if typedObj, ok := obj.(T); ok {
		resource = typedObj
		uid = string(resource.GetUID())
	} else if unstructuredObj, ok := obj.(*unstructured.Unstructured); ok {
		// Fall back to handling unstructured data
		uidValue, found, err := unstructured.NestedString(unstructuredObj.Object, "metadata", "uid")
		if err != nil || !found {
			// Handle error or log it
			return
		}
		uid = uidValue

		// Convert unstructured to T
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.Object, &resource)
		if err != nil {
			// Handle error or log it
			return
		}
	} else {
		// Neither T nor *unstructured.Unstructured
		return
	}

	switch eventType {
	case Added, Modified:
		r.resources[uid] = resource
	case Deleted:
		delete(r.resources, uid)
	}

	// Notify subscribers of the change
	select {
	case r.Changes <- struct{}{}:
	default:
	}
}
