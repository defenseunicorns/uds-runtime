package main

import (
	"crypto/tls"
	"embed"
	"log"
	"net/http"
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
		// Log the error and exit
		message.WarnErr(err, "failed to start the API server")
		os.Exit(1)
	}

	//nolint:gosec,govet
	if inCluster {
		log.Println("Starting server on :8080")

		if err = http.ListenAndServe(":8080", r); err != nil {
			message.WarnErrf(err, "server failed to start: %s", err.Error())
			os.Exit(1)
		}
	} else {
		// create tls config from embedded cert and key
		cert, err := tls.X509KeyPair(localCert, localKey)
		if err != nil {
			log.Fatalf("Failed to load embedded certificate: %v", err)
		}
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		// Create a server with TLS config
		server := &http.Server{
			Addr:      ":8443",
			Handler:   r,
			TLSConfig: tlsConfig,
		}

		log.Println("Starting server on :8443")
		if err = server.ListenAndServeTLS("", ""); err != nil {
			message.WarnErrf(err, "server failed to start: %s", err.Error())
			os.Exit(1)
		}
	}
}
