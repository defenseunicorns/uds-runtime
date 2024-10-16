// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defenseunicorns/uds-runtime/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	// Mock next handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	var session = local.NewBrowserSession()

	tests := []struct {
		name                 string
		localAuthEnabled     bool
		inClusterAuthEnabled bool
		path                 string
		expectedStatusCode   int
		setup                func(*http.Request)
	}{
		{
			name:                 "Local Auth - Allow List Path",
			localAuthEnabled:     true,
			inClusterAuthEnabled: false,
			path:                 "/api/v1/auth",
			expectedStatusCode:   http.StatusOK,
			setup:                func(*http.Request) {},
		},
		{
			name:                 "Local Auth - Valid Session",
			localAuthEnabled:     true,
			inClusterAuthEnabled: false,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusOK,
			setup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{
					Name:  "session_id",
					Value: "valid_session",
				})
				session.Store("valid_session")
				local.Session = session
			},
		},
		{
			name:                 "Local Auth - Invalid Session",
			localAuthEnabled:     true,
			inClusterAuthEnabled: false,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusUnauthorized,
			setup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{
					Name:  "session_id",
					Value: "invalid_session",
				})
				session.Store("valid_session")
				local.Session = session
			},
		},
		{
			name:                 "Local Auth - No Session Cookie",
			localAuthEnabled:     true,
			inClusterAuthEnabled: false,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusUnauthorized,
			setup:                func(*http.Request) {},
		},
		{
			name:                 "In-cluster auth - Valid JWT",
			localAuthEnabled:     false,
			inClusterAuthEnabled: true,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusOK,
			setup: func(r *http.Request) {
				// create jwt for test and set header
				jot := jwt.New(jwt.SigningMethodNone)
				jot.Claims = jwt.MapClaims{
					"groups": []string{"/UDS Core/Admin"},
				}
				token, _ := jot.SignedString(jwt.UnsafeAllowNoneSignatureType)
				r.Header.Set("Authorization", token)
			},
		},
		{
			name:                 "In-cluster auth - Invalid JWT",
			localAuthEnabled:     false,
			inClusterAuthEnabled: true,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusForbidden,
			setup: func(r *http.Request) {
				// create jwt for test and set header
				jot := jwt.New(jwt.SigningMethodNone)
				jot.Claims = jwt.MapClaims{
					"groups": []string{"/UDS Core/bad-group"},
				}
				token, _ := jot.SignedString(jwt.UnsafeAllowNoneSignatureType)
				r.Header.Set("Authorization", token)
			},
		},
		{
			name:                 "In-cluster auth - No JWT",
			localAuthEnabled:     false,
			inClusterAuthEnabled: true,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusUnauthorized,
			setup:                func(*http.Request) {},
		},
		{
			name:                 "Both Auths Disabled",
			localAuthEnabled:     false,
			inClusterAuthEnabled: false,
			path:                 "/api/v1/workloads/pods",
			expectedStatusCode:   http.StatusOK,
			setup:                func(*http.Request) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			config.LocalAuthEnabled = tt.localAuthEnabled
			config.InClusterAuthEnabled = tt.inClusterAuthEnabled

			// Create a request to pass to our handler
			req, err := http.NewRequest("GET", tt.path, nil)
			require.NoError(t, err)

			// Setup session (if any)
			tt.setup(req)

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Create the handler with our middleware
			handler := Auth(nextHandler)

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check the status code
			require.Equal(t, tt.expectedStatusCode, rr.Code)
		})
	}
}
