// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package auth

import (
	"crypto/rand"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/defenseunicorns/uds-runtime/pkg/config"
)

// Configure sets the config vars for local or in-cluster auth
func Configure() {
	// check for local auth first
	localAuthEnabled, err := strconv.ParseBool(strings.ToLower(os.Getenv("LOCAL_AUTH_ENABLED")))
	if err != nil {
		slog.Warn("invalid value for LocalAuthEnabled, must be 'true' or 'false'. Defaulting to 'true'")
		localAuthEnabled = true
	}

	config.LocalAuthEnabled = localAuthEnabled
	if localAuthEnabled {
		slog.Info("Local auth enabled")
		token, err := randomString(96)
		if err != nil {
			slog.Error("Failed to generate local auth token")
			os.Exit(1)
		}
		config.LocalAuthToken = token
		return
	}

	// If local auth is disabled, check for in-cluster auth
	inClusterAuthEnabled, err := strconv.ParseBool(strings.ToLower(os.Getenv("IN_CLUSTER_AUTH_ENABLED")))
	if err != nil {
		slog.Warn("invalid value for InClusterAuthEnabled, must be 'true' or 'false'. Defaulting to 'false'")
		inClusterAuthEnabled = false
	}

	if inClusterAuthEnabled {
		config.InClusterAuthEnabled = inClusterAuthEnabled
		slog.Info("In-cluster auth enabled")
	}
}

// Very limited special chars for git / basic auth
// https://owasp.org/www-community/password-special-characters has complete list of safe chars.
const randomStringChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~-"

// randomString generates a secure random string of the specified length.
func randomString(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = randomStringChars[b%byte(len(randomStringChars))]
	}

	return string(bytes), nil
}
