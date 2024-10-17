// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package main

import (
	"embed"
	"os"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

//go:embed ui/build/*
var assets embed.FS

//go:embed hack/certs/cert.pem
var localCert []byte

//go:embed hack/certs/key.pem
var localKey []byte

func main() {
	message.SetLogLevel(message.DebugLevel)

	r, inCluster, err := api.Setup(&assets)
	if err != nil {
		message.WarnErr(err, "failed to start the API server")
		os.Exit(1)
	}

	err = api.Serve(r, localCert, localKey, inCluster)
	if err != nil {
		os.Exit(1)
	}
}
