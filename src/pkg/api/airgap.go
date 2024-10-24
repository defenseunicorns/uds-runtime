// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"log/slog"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

// serveAirgap starts a server assuming airgap and uses self-signed certificates
func serveAirgap(r *chi.Mux) error {
	err := generateCerts()
	if err != nil {
		return errors.New("failed to generate certs")
	}
	defer cleanupCerts()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		//nolint:gosec,govet
		err := http.ListenAndServeTLS(":8443", "airgap-cert.pem", "airgap-key.pem", r)
		if err != nil {
			slog.Error("Failed to start server:", "error", err)
			os.Exit(1)
		}
	}()
	<-stop
	slog.Info("Shutting down server...")
	return nil
}

// isAirgapped checks if we're in the airgap by checking Google and Cloudflare DNS servers
func isAirgapped(timeout time.Duration) bool {
	googleDNS := "8.8.8.8:53"
	cloudflareDNS := "1.1.1.1:53"

	// Check Google DNS
	googleConn, googleErr := net.DialTimeout("udp", googleDNS, timeout)
	if googleErr == nil {
		defer googleConn.Close()
	}

	// Check Cloudflare DNS
	cloudflareConn, cloudflareErr := net.DialTimeout("udp", cloudflareDNS, timeout)
	if cloudflareErr == nil {
		defer cloudflareConn.Close()
	}

	return !(googleErr == nil && cloudflareErr == nil)
}

// generateCerts creates self-signed certificates for running locally in the airgap
func generateCerts() error {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour), // Valid for 1 year
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
	}

	// Create certificate using template
	derBytes, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&privateKey.PublicKey,
		privateKey,
	)
	if err != nil {
		return err
	}

	// Save certificate to file
	certFile, err := os.Create("airgap-cert.pem")
	if err != nil {
		return err
	}
	defer certFile.Close()

	err = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	if err != nil {
		return err
	}

	// Save private key to file
	keyFile, err := os.Create("airgap-key.pem")
	if err != nil {
		return err
	}
	defer keyFile.Close()

	err = pem.Encode(keyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return err
	}

	return nil
}

func cleanupCerts() {
	slog.Info("Cleaning up short-lived airgap certs")
	os.Remove("airgap-cert.pem")
	os.Remove("airgap-key.pem")
}
