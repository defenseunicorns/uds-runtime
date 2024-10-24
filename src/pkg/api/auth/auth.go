// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package auth

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/incluster"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
)

// Very limited special chars for git / basic auth
// https://owasp.org/www-community/password-special-characters has complete list of safe chars.
const randomStringChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~-"

// UserResponse is the response for the /auth endpoint
type UserResponse struct {
	Group             string `json:"group"`
	PreferredUsername string `json:"preferred-username"`
	Name              string `json:"name"`
	InClusterAuth     bool   `json:"in-cluster-auth"`
}

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
		local.AuthToken = token
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

// RequestHandler is the main handler for the /auth endpoint; it returns a userResponse struct
// indicating whether the request was authenticated via local or in-cluster auth, and relevant user data
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	resp := UserResponse{
		InClusterAuth: false,
	}
	if config.LocalAuthEnabled && !local.Auth(w, r) {
		// auth failed, response is already written
		return
	} else if config.InClusterAuthEnabled {
		// grab values from context set by Auth JWT middleware
		group := r.Context().Value(incluster.GroupKey)
		username := r.Context().Value(incluster.PreferredUserNameKey)
		name := r.Context().Value(incluster.NameKey)

		resp.InClusterAuth = true

		// ensure values are valid
		if group != nil && username != nil && name != nil {
			resp.Group = group.(string)
			resp.Name = name.(string)
			resp.PreferredUsername = username.(string)
		} else {
			slog.Warn("Failed to get group and username from context")
			http.Error(w, "authorization failure", http.StatusInternalServerError)
			return
		}
	}

	// write response
	w.WriteHeader(http.StatusOK)
	bodyBytes, err := json.Marshal(resp)
	if err != nil {
		slog.Debug("failed to marshal response", "error", err)
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bodyBytes)
	if err != nil {
		slog.Debug(fmt.Sprintf("failed to write response: %s", err.Error()))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

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
