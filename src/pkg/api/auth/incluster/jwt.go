// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package incluster

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var allowedGroups = []string{
	"/UDS Core/Admin",
	"/UDS Core/Auditor",
}

type contextKey string

const (
	GroupKey             contextKey = "group"
	PreferredUserNameKey contextKey = "preferred_username"
	NameKey              contextKey = "name"
)

// ValidateJWT checks if the request has a valid JWT token with the required groups.
func ValidateJWT(w http.ResponseWriter, r *http.Request) (*http.Request, bool) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return r, false
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// parse the JWT token without validation (authservice will validate it, we only need the groups here)
	token, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(tokenString, jwt.Claims(jwt.MapClaims{}))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return r, false
	}

	// Check if the token contains a "groups" claim
	if groups, ok := token.Claims.(jwt.MapClaims)["groups"].([]interface{}); ok {
		// Check if any of the token's groups match the allowed groups
		for _, group := range groups {
			for _, allowedGroup := range allowedGroups {
				if group == allowedGroup {
					// Group is allowed, add group and username to request context
					// set context values on request
					r = r.WithContext(context.WithValue(r.Context(), GroupKey, group))
					preferredUsername, preferredUsernameOk := token.Claims.(jwt.MapClaims)["preferred_username"].(string)
					name, nameOk := token.Claims.(jwt.MapClaims)["name"].(string)
					if preferredUsernameOk && nameOk {
						r = r.WithContext(context.WithValue(r.Context(), PreferredUserNameKey, preferredUsername))
						r = r.WithContext(context.WithValue(r.Context(), NameKey, name))
						return r, true
					}
					http.Error(w, "Invalid token claims", http.StatusUnauthorized)
					return r, false
				}
			}
		}
		// If we reach here, no matching group was found
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
		return r, false
	}

	http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	return r, false
}