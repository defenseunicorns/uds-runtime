// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cve

import (
	"database/sql"
	// Driver needed for slq.Open() function call
	_ "github.com/mattn/go-sqlite3"
)

func FetchClusterOverview(db *sql.DB) ([]Overview, error) {
	rows, err := db.Query(`
    WITH LatestReports AS (
        SELECT
            id,
            package_name,
            tag,
            critical,
            high,
            total,
            created_at,
            ROW_NUMBER() OVER (PARTITION BY package_name, tag ORDER BY created_at DESC) AS rn
        FROM
            reports
    )
    SELECT
        p.id AS package_id,
        p.name AS package_name,
        p.tag,
        p.repository,
        p.updated_at,
        lr.critical,
        lr.high,
        lr.total,
        COUNT(s.id) AS scan_count
    FROM
        packages p
    LEFT JOIN
        LatestReports lr ON p.id = lr.id AND p.name = lr.package_name AND p.tag = lr.tag AND lr.rn = 1
    LEFT JOIN
        scans s ON p.id = s.package_id
	WHERE
	    lr.rn = 1
    GROUP BY
        p.id,
        p.updated_at,
        p.name,
        p.tag,
		p.repository,
        lr.critical,
        lr.high,
        lr.total
    ORDER BY
        p.id;
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clusterOverviews []Overview

	// Iterate over the rows and populate the reports slice
	for rows.Next() {
		var clusterOverview Overview
		if err := rows.Scan(
			&clusterOverview.PackageID,
			&clusterOverview.PackageName,
			&clusterOverview.Tag,
			&clusterOverview.Repository,
			&clusterOverview.UpdatedAt,
			&clusterOverview.Critical,
			&clusterOverview.High,
			&clusterOverview.Total,
			&clusterOverview.ImagesWithPackage,
		); err != nil {
			return nil, err
		}
		clusterOverviews = append(clusterOverviews, clusterOverview)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clusterOverviews, nil
}

func FetchFindings(db *sql.DB) ([]Finding, error) {
	rows, err := db.Query(`
        SELECT
            digest,
            pkg_name,
            installed_version,
			fixed_version,
            type,
            vulnerability_id,
            severity,
            severity_source
        FROM
            vulnerabilities;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var findings []Finding

	// Iterate over the rows and populate the reports slice
	for rows.Next() {
		var finding Finding
		if err := rows.Scan(
			&finding.ImageID,
			&finding.AppName,
			&finding.AppVersion,
			&finding.FixedVersion,
			&finding.Author,
			&finding.Vulnerability,
			&finding.Severity,
			&finding.Reporter,
		); err != nil {
			return nil, err
		}
		findings = append(findings, finding)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return findings, nil
}
