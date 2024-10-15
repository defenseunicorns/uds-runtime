// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package local

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/defenseunicorns/uds-runtime/pkg/api/auth"
	"github.com/defenseunicorns/uds-runtime/pkg/config"
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

// session is a global variable that holds the current session
var session = NewBrowserSession()

// AuthHandler handle validating tokens and session cookies for local authentication
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if config.LocalAuthEnabled {
		token := r.URL.Query().Get("token")
		if token == "" {
			// Handle session cookie validation
			ValidateSessionCookie(w, r)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if token != auth.LocalAuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// valid token, generate session id and set cookie
		sessionID := generateSessionID(w)
		session.Store(sessionID)
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

func ValidateSessionCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil || !session.Validate(cookie.Value) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}

func generateSessionID(w http.ResponseWriter) string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
		return ""
	}
	return hex.EncodeToString(bytes)
}
