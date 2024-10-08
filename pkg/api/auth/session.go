// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

type InMemoryStorage struct {
	sessionID string
	mutex     sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) StoreSession(sessionID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Replace the old session with the new one
	s.sessionID = sessionID
}

func (s *InMemoryStorage) ValidateSession(sessionID string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Check if the provided sessionID matches the stored session
	return s.sessionID == sessionID
}

func (s *InMemoryStorage) RemoveSession() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Clear the session
	s.sessionID = ""
}

var storage = NewInMemoryStorage()

// TokenAuthenticator ensures the request has a valid token.
func TokenAuthenticator(validToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.URL.Query().Get("token")
			if token == "" {
				ValidateSessionCookie(next, w, r)
			} else if token != validToken {
				// If a token is passed in and its not valid, return unauthorized
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				// If a token is passed in and its valid, set the session ID and continue
				if token != "" && token == validToken {
					sessionID := generateSessionID()
					storage.StoreSession(sessionID)
					http.SetCookie(w, &http.Cookie{
						Name:     "session_id",
						Value:    sessionID,
						HttpOnly: true,
						Secure:   true,
						SameSite: http.SameSiteStrictMode,
						Path:     "/",
					})

					next.ServeHTTP(w, r)
				}
			}
		})
	}
}

func ValidateSessionCookie(next http.Handler, w http.ResponseWriter, r *http.Request) {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil || !storage.ValidateSession(cookie.Value) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w, r)
}

func generateSessionID() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		// Handle error
		return ""
	}
	return hex.EncodeToString(bytes)
}

// Connect is a head-only request to test the connection.
func Connect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var allowedGroups = []string{
	"/UDS Core/Admin",
	"/UDS Core/Auditor",
}

// RequireJWT is a middleware that checks if the request has a valid JWT token with the required groups.
func RequireJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// parse the JWT token without validation (authservice will validate it, we only need the groups here)
		token, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(tokenString, jwt.Claims(jwt.MapClaims{}))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token contains a "groups" claim
		if groups, ok := token.Claims.(jwt.MapClaims)["groups"].([]interface{}); ok {
			// Check if any of the token's groups match the allowed groups
			for _, group := range groups {
				for _, allowedGroup := range allowedGroups {
					if group == allowedGroup {
						// Group is allowed, continue to the next handler
						next.ServeHTTP(w, r)
						return
					}
				}
			}
			// If we reach here, no matching group was found
			http.Error(w, "Insufficient permissions", http.StatusForbidden)
			return
		}

		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
	})
}
