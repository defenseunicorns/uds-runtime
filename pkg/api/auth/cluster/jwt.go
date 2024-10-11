// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cluster

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var allowedGroups = []string{
	"/UDS Core/Admin",
	"/UDS Core/Auditor",
}

// ValidateJWT is a middleware that checks if the request has a valid JWT token with the required groups.
func ValidateJWT(w http.ResponseWriter, r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return false
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// parse the JWT token without validation (authservice will validate it, we only need the groups here)
	token, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(tokenString, jwt.Claims(jwt.MapClaims{}))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return false
	}

	// Check if the token contains a "groups" claim
	if groups, ok := token.Claims.(jwt.MapClaims)["groups"].([]interface{}); ok {
		// Check if any of the token's groups match the allowed groups
		for _, group := range groups {
			for _, allowedGroup := range allowedGroups {
				if group == allowedGroup {
					// Group is allowed
					return true
				}
			}
		}
		// If we reach here, no matching group was found
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
		return false
	}

	http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	return false
}
