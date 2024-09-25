// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package security

import (
	"database/sql"
	// Driver needed for slq.Open() function call
	_ "github.com/mattn/go-sqlite3"
)

func FetchClusterOverview(dbPath string) ([]ClusterOverview, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
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
        lr.critical,
        lr.high,
        lr.total
    ORDER BY
        p.id;
    `

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clusterOverviews []ClusterOverview

	// Iterate over the rows and populate the reports slice
	for rows.Next() {
		var clusterOverview ClusterOverview
		if err := rows.Scan(
			&clusterOverview.PackageID,
			&clusterOverview.PackageName,
			&clusterOverview.Tag,
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
