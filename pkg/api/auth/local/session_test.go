// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package local

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
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

func TestRequireJWT(t *testing.T) {
	// Create a sample handler that the middleware will wrap
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create the middleware
	middleware := RequireJWT(nextHandler)

	// Helper function to create a JWT token without signing
	createToken := func(groups []string) string {
		claims := jwt.MapClaims{
			"groups": groups,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
		tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
		return tokenString
	}

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Valid token with allowed group",
			token:          createToken([]string{"/UDS Core/Admin"}),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid token without allowed group",
			token:          createToken([]string{"guest"}),
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "Valid token with empty group",
			token:          createToken([]string{}),
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "Invalid token",
			token:          "invalid.token.string",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Missing token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler
			req, _ := http.NewRequest("GET", "/", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the middleware
			middleware.ServeHTTP(rr, req)

			// Check the status code
			require.Equal(t, tt.expectedStatus, rr.Code, "handler returned wrong status code")
		})
	}
}
