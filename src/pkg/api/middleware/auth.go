// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package middleware

import (
	"net/http"
	"strings"

	clusterAuth "github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/cluster"
	localAuth "github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
)

// Auth is a middleware that handles all API authentication for UDS Runtime
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow list endpoints (used for local auth only)
		apiAllowList := []string{
			"/api/v1/auth",
		}
		if config.LocalAuthEnabled {
			// check if the request is in the allow list
			if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/swagger") {
				for _, path := range apiAllowList {
					if r.URL.Path == path {
						next.ServeHTTP(w, r) // path allowed
						return
					}
				}
				if valid := localAuth.ValidateSessionCookie(w, r); valid {
					next.ServeHTTP(w, r)
					return
				}
			}
		} else if config.InClusterAuthEnabled {
			if valid := clusterAuth.ValidateJWT(w, r); valid {
				next.ServeHTTP(w, r)
			}
		}
		next.ServeHTTP(w, r)
	})
}
