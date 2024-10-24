// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package local

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
)

// AuthToken is the token used for local auth
var AuthToken = ""

// SessionID is the session ID for local auth; generated after token is validated
var SessionID = ""

// Auth validates tokens and session cookies for local authentication
func Auth(w http.ResponseWriter, r *http.Request) bool {
	var once sync.Once
	token := r.URL.Query().Get("token")
	if token == "" {
		// Handle session cookie validation
		if valid := ValidateSessionCookie(w, r); valid {
			w.WriteHeader(http.StatusOK)
			return true
		}
		w.WriteHeader(http.StatusUnauthorized)
		return false
	} else if token != AuthToken {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	// valid token, generate session id and set cookie
	once.Do(func() {
		SessionID = GenerateSessionID(w)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    SessionID,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
	})
	return true
}

// ValidateSessionCookie validates the session cookie in the request
func ValidateSessionCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil || SessionID != cookie.Value {
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
