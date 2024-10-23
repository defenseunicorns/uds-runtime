// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package local

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
)

// BrowserSession is a struct that holds the session ID of the current session
// The session ID is generated once tokens are validated during local auth mode and is stored in a cookie
type BrowserSession struct {
	sessionID string
	mutex     sync.RWMutex
}

// AuthToken is the token used for local auth
var AuthToken = ""

// NewBrowserSession creates a new BrowserSession
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

// Session is a global variable that holds the current session
var Session = NewBrowserSession()

// Auth validates tokens and session cookies for local authentication
func Auth(w http.ResponseWriter, r *http.Request) bool {
	token := r.URL.Query().Get("token")
	if token == "" {
		// Handle session cookie validation
		if valid := ValidateSessionCookie(w, r); valid {
			w.WriteHeader(http.StatusOK)
			return false
		}
		w.WriteHeader(http.StatusUnauthorized)
		return false
	} else if token != AuthToken {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	// valid token, generate session id and set cookie
	sessionID := GenerateSessionID(w)
	Session.Store(sessionID)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	return true
}

// ValidateSessionCookie validates the session cookie in the request
func ValidateSessionCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil || !Session.Validate(cookie.Value) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}

// GenerateSessionID is a util function to generate a random session ID
func GenerateSessionID(w http.ResponseWriter) string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
		return ""
	}
	return hex.EncodeToString(bytes)
}
