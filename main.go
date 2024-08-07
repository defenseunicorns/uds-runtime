package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/defenseunicorns/uds-runtime/pkg/api"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

//go:embed ui/build/*
var assets embed.FS

func main() {
	message.SetLogLevel(message.DebugLevel)
	r, err := api.Setup(&assets)
	if err != nil {
		// Log the error and exit
		message.WarnErr(err, "failed to start the API server")
		os.Exit(1)
	}
	log.Println("Starting server on :8080")
	//nolint:gosec
	if err = http.ListenAndServe(":8080", r); err != nil {
		message.WarnErrf("server failed to start: %w", err.Error())
		os.Exit(1)
	}
}
