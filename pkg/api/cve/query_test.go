// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cve

import (
	"database/sql"
	// "os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

// note: error will be logged on init since test db won't be created yet
func TestFetchClusterOverview(t *testing.T) {
	// Create an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	// Create tables and insert test data
	_, err = db.Exec(`
        CREATE TABLE packages (
            id INTEGER PRIMARY KEY,
            name TEXT,
            tag TEXT,
			repository TEXT,
            updated_at TEXT
        );
        CREATE TABLE reports (
            id INTEGER PRIMARY KEY,
            package_name TEXT,
            tag TEXT,
            critical INTEGER,
            high INTEGER,
            total INTEGER,
            created_at TEXT
        );
        CREATE TABLE scans (
            id INTEGER PRIMARY KEY,
            package_id INTEGER
        );

        INSERT INTO packages (id, name, tag, repository, updated_at) VALUES
        (1, 'package1', 'v1.0', 'defenseunicorns', '2023-01-01'),
        (2, 'package1', 'v1.0', 'defenseunicorns', '2023-01-02'),
        (3, 'package2', 'v1.0', 'defenseunicorns', '2023-01-02');

        INSERT INTO reports (id, package_name, tag, critical, high, total, created_at) VALUES
        (1, 'package1', 'v1.0', 5, 10, 15, '2023-01-01'),
        (2, 'package1', 'v1.0', 3, 6, 9, '2023-01-02'), -- Latest report
        (3, 'package2', 'v1.0', 2, 4, 6, '2023-01-01');

        INSERT INTO scans (id, package_id) VALUES
        (1, 2),
        (2, 2),
        (3, 3);
    `)

	require.NoError(t, err)

	// Call the function
	overviews, err := FetchClusterOverview(db)
	require.NoError(t, err)

	// Verify the results
	require.Len(t, overviews, 2)

	require.Equal(t, 2, overviews[0].PackageID)
	require.Equal(t, "package1", overviews[0].PackageName)
	require.Equal(t, "v1.0", overviews[0].Tag)
	require.Equal(t, "defenseunicorns", overviews[0].Repository)
	require.Equal(t, "2023-01-02", overviews[0].UpdatedAt)
	require.Equal(t, 3, overviews[0].Critical)
	require.Equal(t, 6, overviews[0].High)
	require.Equal(t, 9, overviews[0].Total)
	require.Equal(t, 2, overviews[0].ImagesWithPackage)

	require.Equal(t, 3, overviews[1].PackageID)
	require.Equal(t, "package2", overviews[1].PackageName)
	require.Equal(t, "v1.0", overviews[1].Tag)
	require.Equal(t, "defenseunicorns", overviews[1].Repository)
	require.Equal(t, "2023-01-02", overviews[1].UpdatedAt)
	require.Equal(t, 2, overviews[1].Critical)
	require.Equal(t, 4, overviews[1].High)
	require.Equal(t, 6, overviews[1].Total)
	require.Equal(t, 1, overviews[1].ImagesWithPackage)
}

func TestFetchFindings(t *testing.T) {
	// Create an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	// Create the vulnerabilities table
	_, err = db.Exec(`
        CREATE TABLE vulnerabilities (
            digest TEXT,
            pkg_name TEXT,
            installed_version TEXT,
			fixed_version TEXT,
            type TEXT,
            vulnerability_id TEXT,
            severity TEXT,
            severity_source TEXT
        );
    `)
	require.NoError(t, err)

	// Insert test data
	_, err = db.Exec(`
        INSERT INTO vulnerabilities (digest, pkg_name, installed_version, fixed_version, type, vulnerability_id, severity, severity_source) VALUES
        ('sha256:abc123', 'testpkg', '1.0.0', '1.1.0', 'library', 'CVE-1234', 'high', 'NVD'),
        ('sha256:def456', 'anotherpkg', '2.0.0', '2.1.0', 'library', 'CVE-5678', 'medium', 'NVD');
    `)
	require.NoError(t, err)

	// Call the FetchFindings function
	findings, err := FetchFindings(db)
	require.NoError(t, err)

	// Verify the results
	require.Len(t, findings, 2)

	require.Equal(t, "sha256:abc123", findings[0].ImageID)
	require.Equal(t, "testpkg", findings[0].AppName)
	require.Equal(t, "1.0.0", findings[0].AppVersion)
	require.Equal(t, "1.1.0", findings[0].FixedVersion)
	require.Equal(t, "library", findings[0].Author)
	require.Equal(t, "CVE-1234", findings[0].Vulnerability)
	require.Equal(t, "high", findings[0].Severity)
	require.Equal(t, "NVD", findings[0].Reporter)

	require.Equal(t, "sha256:def456", findings[1].ImageID)
	require.Equal(t, "anotherpkg", findings[1].AppName)
	require.Equal(t, "2.0.0", findings[1].AppVersion)
	require.Equal(t, "2.1.0", findings[1].FixedVersion)
	require.Equal(t, "library", findings[1].Author)
	require.Equal(t, "CVE-5678", findings[1].Vulnerability)
	require.Equal(t, "medium", findings[1].Severity)
	require.Equal(t, "NVD", findings[1].Reporter)
}
