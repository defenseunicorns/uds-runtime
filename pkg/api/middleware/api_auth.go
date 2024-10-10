// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package middleware

import (
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/auth"
	"github.com/defenseunicorns/uds-runtime/pkg/config"
)

func ValidateLocalAuthSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.LocalAuthEnabled {
			auth.ValidateSessionCookie(next, w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
