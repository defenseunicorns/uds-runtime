// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package local

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestStoreSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)

	require.Equal(t, sessionID, storage.sessionID, "expected sessionID to be stored correctly")
}

func TestValidateSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)

	require.True(t, storage.Validate(sessionID), "expected sessionID to be valid")

	invalidSessionID := "invalid-session-id"
	require.False(t, storage.Validate(invalidSessionID), "expected invalid sessionID to be invalid")
}

func TestRemoveSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)
	storage.Remove()

	require.Empty(t, storage.sessionID, "expected sessionID to be empty after removal")
	require.False(t, storage.Validate(sessionID), "expected sessionID to be invalid after removal")
}

func TestAuthHandler(t *testing.T) {
	// Save the original values and restore them after the test
	originalLocalAuthEnabled := config.LocalAuthEnabled
	originalLocalAuthToken := auth.LocalAuthToken
	defer func() {
		config.LocalAuthEnabled = originalLocalAuthEnabled
		auth.LocalAuthToken = originalLocalAuthToken
	}()

	config.LocalAuthEnabled = true
	auth.LocalAuthToken = "test-token"

	tests := []struct {
		name           string
		token          string
		expectedStatus int
		withCookie     bool
	}{
		{name: "Valid token", token: "test-token", expectedStatus: http.StatusOK, withCookie: false},
		{name: "Invalid token", token: "wrong-token", expectedStatus: http.StatusUnauthorized, withCookie: false},
		{name: "No token", token: "", expectedStatus: http.StatusUnauthorized, withCookie: false},
		{name: "Valid cookie", token: "", expectedStatus: http.StatusOK, withCookie: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/auth", nil)
			require.NoError(t, err)

			// if token provided, add it to the URL as a query param
			if tt.token != "" {
				q := req.URL.Query()
				q.Add("token", tt.token)
				req.URL.RawQuery = q.Encode()
			}

			// create cookie if specified
			if tt.withCookie {
				sessionID := generateSessionID(httptest.NewRecorder())
				Session.Store(sessionID)
				req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID})
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(AuthHandler)

			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code, "handler returned wrong status code")
		})
	}
}
