// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cve

type Reports struct {
	ClusterOverview []ClusterOverview `json:"cluster_overview"`
	ByImage         []ByImage         `json:"by_image"`
}

type ClusterOverview struct {
	PackageID         int    `json:"package_id"`
	PackageName       string `json:"package_name"`
	Tag               string `json:"package_version"`
	Repository        string `json:"repository"`
	UpdatedAt         string `json:"build_date"`
	Critical          int    `json:"critical"`
	High              int    `json:"high"`
	Total             int    `json:"cve_count"`
	ImagesWithPackage int    `json:"images_with_package"`
}

type ByImage struct {
	ImageID       string `json:"image_id"`
	Component     string `json:"component"`
	AppName       string `json:"app_name"`
	AppVersion    string `json:"app_version"`
	Author        string `json:"author"`
	Vulnerability string `json:"vulnerability"`
	Severity      string `json:"severity"`
	Reporter      string `json:"reporter"`
	VexStatus     string `json:"vex_status"`
	Justified     string `json:"justified"`
}
