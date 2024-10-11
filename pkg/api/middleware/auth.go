// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package middleware

import (
	"net/http"
	"strings"

	clusterAuth "github.com/defenseunicorns/uds-runtime/pkg/api/auth/cluster"
	localAuth "github.com/defenseunicorns/uds-runtime/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/pkg/config"
)

// validateLocalAuthSession is a middleware that checks if the request has a valid local auth session cookie.
func validateLocalAuthSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.LocalAuthEnabled {
			localAuth.ValidateSessionCookie(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// validateJWTAuth is a middleware that checks if the request has a valid JWT token with the required groups.
func validateJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.InClusterAuthEnabled {
			if valid := clusterAuth.ValidateJWT(w, r); valid {
				next.ServeHTTP(w, r)
			}
			return
		}
	})
}

// Auth is a middleware that handles all API authentication for UDS Runtime
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow list endpoints (used for local auth only)
		apiAllowList := []string{
			"/api/v1/auth",
		}
		if config.LocalAuthEnabled {
			// check if the request is in the allow list
			if strings.HasPrefix(r.URL.Path, "/api/") {
				for _, path := range apiAllowList {
					if r.URL.Path == path {
						next.ServeHTTP(w, r)
						return
					}
				}
				validateLocalAuthSession(next)
			} else if config.InClusterAuthEnabled {
				validateJWTAuth(next)
			}
		}
		next.ServeHTTP(w, r)
	})
}
