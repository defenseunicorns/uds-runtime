package main

import (
	"embed"

	"github.com/defenseunicorns/uds-engine/pkg/api"
	"github.com/defenseunicorns/zarf/src/pkg/message"
)

//go:embed ui/build/*
var assets embed.FS

func main() {
	message.SetLogLevel(message.DebugLevel)
	if err := api.Start(assets); err != nil {
		message.Fatal(err, "Failed to start server")
	}
}
