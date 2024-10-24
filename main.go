// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package main

import (
	"embed"
	"fmt"
	"log/slog"
	"os"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api"
)

//go:embed ui/build/*
var assets embed.FS

//go:embed hack/certs/cert.pem
var localCert []byte

//go:embed hack/certs/key.pem
var localKey []byte

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	r, inCluster, err := api.Setup(&assets)
	if err != nil {
		slog.Warn(fmt.Sprintf("failed to start the API server: %s", err))
		os.Exit(1)
	}

	err = api.Serve(r, localCert, localKey, inCluster)
	if err != nil {
		os.Exit(1)
	}
}
