// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package security

type Reports struct {
	ClusterOverview []ClusterOverview `json:"cluster-overview"`
	ByImage         []ByImage         `json:"by-image"`
}

type ClusterOverview struct {
	PackageID   int    `json:"package_id"`
	PackageName string `json:"package-name"`
	Tag         string `json:"package-version"`
	UpdatedAt   string `json:"build-date"`
	//TODO: Author            string `json:"author"`
	Critical          int `json:"critical"`
	High              int `json:"high"`
	Total             int `json:"cve-count"`
	ImagesWithPackage int `json:"images-with-package"`
}

type ByImage struct {
	ImageID       string `json:"image-id"`
	Component     string `json:"component"`
	AppName       string `json:"app-name"`
	AppVersion    string `json:"app-version"`
	Author        string `json:"author"`
	Vulnerability string `json:"vulnerability"`
	Severity      string `json:"severity"`
	Reporter      string `json:"reporter"`
	VexStatus     string `json:"vex-status"`
	Justified     string `json:"justified"`
}
