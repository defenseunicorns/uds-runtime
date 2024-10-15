// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

//go:build unit

package cluster

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestValidateJWT(t *testing.T) {
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
			name:           "Valid token with another allowed group",
			token:          createToken([]string{"/UDS Core/Auditor"}),
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

			// Call the function directly
			result := ValidateJWT(rr, req)

			// Check the status code
			require.Equal(t, tt.expectedStatus, rr.Code, "handler returned wrong status code")

			// Check the return value
			expectedResult := tt.expectedStatus == http.StatusOK
			require.Equal(t, expectedResult, result, "ValidateJWT returned unexpected result")
		})
	}
}
