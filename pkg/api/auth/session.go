// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"sync"

	"github.com/defenseunicorns/uds-runtime/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type BrowserSession struct {
	sessionID string
	mutex     sync.RWMutex
}

func NewBrowserSession() *BrowserSession {
	return &BrowserSession{}
}

func (s *BrowserSession) Store(sessionID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Replace the old session with the new one
	s.sessionID = sessionID
}

func (s *BrowserSession) Validate(sessionID string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Check if the provided sessionID matches the stored session
	return s.sessionID == sessionID
}

func (s *BrowserSession) Remove() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Clear the session
	s.sessionID = ""
}

// todo: where is this used? rename to session or something
var storage = NewBrowserSession()

// LocalAuthHandler handle validating tokens and session cookies for local authentication
func LocalAuthHandler(w http.ResponseWriter, r *http.Request) {
	if config.LocalAuthEnabled {
		token := r.URL.Query().Get("token")
		if token == "" {
			// Handle session cookie validation
			validateSessionCookie(w, r)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if token != config.LocalAuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// valid token, generate session id and set cookie
		sessionID := generateSessionID()
		storage.Store(sessionID)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		w.WriteHeader(http.StatusOK)
	}

	// not using local auth, return ok
	w.WriteHeader(http.StatusOK)
}

func ValidateSessionCookie(next http.Handler, w http.ResponseWriter, r *http.Request) {
	validateSessionCookie(w, r)
	next.ServeHTTP(w, r)
}

func validateSessionCookie(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil || !storage.Validate(cookie.Value) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func generateSessionID() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		// Handle error
		return ""
	}
	return hex.EncodeToString(bytes)
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
