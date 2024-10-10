// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package main

import (
	"embed"
	"os"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
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
