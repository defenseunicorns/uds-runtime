package cve

import (
	"database/sql"
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
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	clusterOverviews, err := FetchClusterOverview(db)
	if err != nil {
		log.Printf("Error fetching cluster overviews data: %v", err)
		return
	}

	findings, err := FetchFindings(db)
	if err != nil {
		log.Printf("Error fetching findings data: %v", err)
		return
	}

	securityReports := Reports{ClusterOverview: clusterOverviews, Findings: findings}

	cacheMutex.Lock()
	cache = securityReports
	cacheMutex.Unlock()
}

func GetReports() Reports {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	return cache
}
