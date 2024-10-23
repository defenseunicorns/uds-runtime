package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestAuthRequestHandlerLocal(t *testing.T) {
	config.LocalAuthEnabled = true
	local.AuthToken = "test-token"

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
				sessionID := local.GenerateSessionID(httptest.NewRecorder())
				local.Session.Store(sessionID)
				req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID})
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RequestHandler)

			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code, "handler returned wrong status code")
		})
	}
}

// todo: test in-cluster
