// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package resources

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

func HasCRD(targetGVR schema.GroupVersionResource, CRDs *ResourceList) bool {
	crds := CRDs.GetResources("", "")

	for _, crd := range crds {
		name := crd.GetName()
		if name == fmt.Sprintf("%s.%s", targetGVR.Resource, targetGVR.Group) {
			return true
		}
	}

	return false
}

// AddCustomListeners adds additional listeners to a shared informer for updating Custom Resource informers
func AddCustomListeners(informer cache.SharedIndexInformer, runtimeCache *Cache) {
	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(_ interface{}) {
			notifyCustomResources(runtimeCache)
		},
		UpdateFunc: func(_, _ any) {
			notifyCustomResources(runtimeCache)
		},
		DeleteFunc: func(_ any) {
			notifyCustomResources(runtimeCache)
		},
	})
	if err != nil {
		log.Printf("Error setting up CRD informer: %v", err)
	}
}

func notifyCustomResources(c *Cache) {
	// Send to UDSExemptions channel if initialized
	if c.UDSExemptions != nil {
		c.UDSExemptions.mutex.Lock()
		defer c.UDSExemptions.mutex.Unlock()
		select {
		case c.UDSExemptions.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("UDSExemptions is nil")
	}

	// Send to UDSPackages channel if initialized
	if c.UDSPackages != nil {
		c.UDSPackages.mutex.Lock()
		defer c.UDSPackages.mutex.Unlock()
		select {
		case c.UDSPackages.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("UDSPackages is nil")
	}

	// Send to VirtualServices channel if initialized
	if c.VirtualServices != nil {
		c.VirtualServices.mutex.Lock()
		defer c.VirtualServices.mutex.Unlock()
		select {
		case c.VirtualServices.Changes <- struct{}{}:
		default:
		}
	} else {
		log.Println("VirtualServices is nil")
	}
}
