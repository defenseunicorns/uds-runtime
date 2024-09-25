package cve

import (
	"log"
	"os"
	"sync"
)

var (
	cache      Reports
	cacheMutex sync.RWMutex
)

func init() {
	// Initialize the cache on startup
	updateCache()
}

func updateCache() {
	dbPath := os.Getenv("SECURITY_HUB_DB_PATH")
	clusterOverviews, err := FetchClusterOverview(dbPath)
	if err != nil {
		log.Printf("Error fetching cluster overviews: %v", err)
		return
	}

	//TODO: FetchImageData implementation
	byImage := []ByImage{}

	securityReports := Reports{ClusterOverview: clusterOverviews, ByImage: byImage}

	cacheMutex.Lock()
	cache = securityReports
	cacheMutex.Unlock()
}

func GetReports() Reports {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	return cache
}
