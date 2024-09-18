// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package middleware

import (
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/auth"
)

func ValidateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth.ValidateSessionCookie(next, w, r)
	})
}
