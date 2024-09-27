// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cve

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"
)

var (
	overviewCache            []Overview
	overviewCacheMutex       sync.RWMutex
	overviewCacheInitialized bool

	findingCache            []Finding
	findingCacheMutex       sync.RWMutex
	findingCacheInitialized bool
)

func updateOverviewCache() {
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

	overviewCacheMutex.Lock()
	overviewCache = clusterOverviews
	overviewCacheMutex.Unlock()
	overviewCacheInitialized = true
}

func updateFindingCache() {
	dbPath := os.Getenv("SECURITY_HUB_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	findings, err := FetchFindings(db)
	if err != nil {
		log.Printf("Error fetching findings data: %v", err)
		return
	}

	findingCacheMutex.Lock()
	findingCache = findings
	findingCacheMutex.Unlock()
	findingCacheInitialized = true
}

func GetOverviews() []Overview {
	if !overviewCacheInitialized {
		updateOverviewCache()
	}
	overviewCacheMutex.RLock()
	defer overviewCacheMutex.RUnlock()
	return overviewCache
}

func GetFindings(page, limit int) ([]Finding, int, error) {
	if !findingCacheInitialized {
		updateFindingCache()
	}

	findingCacheMutex.RLock()
	defer findingCacheMutex.RUnlock()

	total := len(findingCache)
	if total == 0 {
		return nil, 0, nil
	}

	if limit <= 0 {
		return nil, total, errors.New("limit must be greater than 0")
	}

	if page <= 0 {
		return nil, total, errors.New("page must be greater than 0")
	}

	offset := (page - 1) * limit
	if offset >= total {
		return nil, total, nil
	}

	end := offset + limit
	if end > total {
		end = total
	}

	paginatedFindings := findingCache[offset:end]
	return paginatedFindings, total, nil
}
