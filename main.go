package main

import (
	"embed"
	"os"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

//go:embed ui/build/*
var assets embed.FS

func main() {
	message.SetLogLevel(message.DebugLevel)
	if err := api.Start(assets); err != nil {
		// Log the error and exit
		message.WarnErr(err, "failed to start the API server")
		os.Exit(1)
	}
}
