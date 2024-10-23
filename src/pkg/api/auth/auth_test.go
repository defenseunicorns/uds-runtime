package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/incluster"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestLocalAuthRequestHandler(t *testing.T) {
	config.LocalAuthEnabled = true
	local.AuthToken = "test-token"

	// Expected response in local mode unless an error has occurred
	userResp := &userResponse{
		InClusterAuth: false,
	}

	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedResp   *userResponse
		withCookie     bool
	}{
		{
			name:           "Valid token",
			token:          "test-token",
			expectedStatus: http.StatusOK,
			withCookie:     false,
			expectedResp:   userResp,
		},
		{
			name:           "Invalid token",
			token:          "wrong-token",
			expectedStatus: http.StatusUnauthorized,
			withCookie:     false,
			expectedResp:   nil,
		},
		{
			name:           "No token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			withCookie:     false,
			expectedResp:   nil,
		},
		{
			name:           "Valid cookie",
			token:          "",
			expectedStatus: http.StatusOK,
			withCookie:     true,
			expectedResp:   nil,
		},
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

			// Verify response body
			if tt.expectedResp != nil {
				var resp userResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				require.NoError(t, err, "failed to unmarshal response")
				require.Equal(t, tt.expectedResp, &resp,
					"handler returned unexpected response: %v")
			}
		})
	}
}

func TestInClusterRequestHandler(t *testing.T) {
	config.LocalAuthEnabled = false
	config.InClusterAuthEnabled = true

	tests := []struct {
		name           string
		setupContext   func(context.Context) context.Context
		expectedStatus int
		expectedResp   *userResponse
	}{
		{
			name: "InCluster auth success",
			setupContext: func(ctx context.Context) context.Context {
				ctx = context.WithValue(ctx, incluster.GroupKey, "admin-group")
				ctx = context.WithValue(ctx, incluster.PreferredUserNameKey, "doug@defenseunicorns.com")
				ctx = context.WithValue(ctx, incluster.NameKey, "Doug Unicorn")
				return ctx
			},
			expectedStatus: http.StatusOK,
			expectedResp: &userResponse{
				InClusterAuth:     true,
				Group:             "admin-group",
				Name:              "Dough Unicorn",
				PreferredUsername: "doug@defenseunicorns.com",
			},
		},
		{
			name: "InCluster auth missing values",
			setupContext: func(ctx context.Context) context.Context {
				ctx = context.WithValue(ctx, incluster.GroupKey, "admin-group")
				// Missing PreferredUserNameKey and NameKey
				return ctx
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request with test context
			req, err := http.NewRequest("GET", "/auth", nil)
			require.NoError(t, err)
			req = req.WithContext(tt.setupContext(req.Context()))

			// Create response recorder and execute handler
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RequestHandler)
			handler.ServeHTTP(rr, req)

			// Check status code
			require.Equal(t, tt.expectedStatus, rr.Code,
				"handler returned wrong status code: %v")

			// Verify response body
			if tt.expectedResp != nil {
				var resp userResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				require.NoError(t, err, "failed to unmarshal response")
				require.Equal(t, tt.expectedResp, &resp,
					"handler returned unexpected response: %v")
			}
		})
	}
}

func TestNoAuthEnabled(t *testing.T) {
	config.LocalAuthEnabled = false
	config.InClusterAuthEnabled = false

	req, err := http.NewRequest("GET", "/auth", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestHandler)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
}

// todo: test configure?
