// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

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
