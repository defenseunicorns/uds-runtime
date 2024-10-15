// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package resources

import (
	"log"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

// CRDs is a thread-safe struct to store the list of CRDs and notify subscribers of changes.
type CRDs struct {
	mutex       sync.RWMutex
	definitions map[string]bool
}

func NewCRDs() *CRDs {
	return &CRDs{
		definitions: make(map[string]bool),
	}
}

func (c *Cache) setupCRDInformer() {
	crdGVR := schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}
	crdInformer := c.dynamicFactory.ForResource(crdGVR).Informer()
	_, err := crdInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.CRDs.addCRD(obj)
			notifyDynamicResources(c)
		},
		UpdateFunc: func(_, newObj any) {
			c.CRDs.addCRD(newObj)
			notifyDynamicResources(c)
		},
		DeleteFunc: func(obj any) {
			c.CRDs.removeCRD(obj)
			notifyDynamicResources(c)
		},
	})
	if err != nil {
		log.Printf("Error setting up CRD informer: %v", err)
	}
}

func (c *CRDs) addCRD(crd interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	resource, err := ToUnstructured(crd)
	if err != nil {
		log.Printf("Error converting CRD to unstructured: %v", err)
		return
	}
	crdName := resource.Object["metadata"].(map[string]interface{})["name"].(string)
	c.definitions[crdName] = true
}

func (c *CRDs) removeCRD(crd interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	resource, err := ToUnstructured(crd)
	if err != nil {
		log.Printf("Error converting CRD to unstructured: %v", err)
		return
	}
	crdName := resource.Object["metadata"].(map[string]interface{})["name"].(string)
	delete(c.definitions, crdName)
}

func (c *CRDs) Contains(targetGVR schema.GroupVersionResource) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, exists := c.definitions[targetGVR.Resource+"."+targetGVR.Group]
	return exists
}

func notifyDynamicResources(c *Cache) {
	c.CRDs.mutex.Lock()
	defer c.CRDs.mutex.Unlock()

	// Send to UDSExemptions channel if initialized
	if c.UDSExemptions != nil {
		select {
		case c.UDSExemptions.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("UDSExemptions is nil")
	}

	// Send to UDSPackages channel if initialized
	if c.UDSPackages != nil {
		select {
		case c.UDSPackages.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("UDSPackages is nil")
	}

	// Send to VirtualServices channel if initialized
	if c.VirtualServices != nil {
		select {
		case c.VirtualServices.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("VirtualServices is nil")
	}
}
