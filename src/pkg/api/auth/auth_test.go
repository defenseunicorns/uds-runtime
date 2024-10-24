package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/incluster"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	"github.com/defenseunicorns/uds-runtime/src/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestConfigure(t *testing.T) {
	// Helper function to reset environment and config state
	reset := func() {
		os.Unsetenv("LOCAL_AUTH_ENABLED")
		os.Unsetenv("IN_CLUSTER_AUTH_ENABLED")
		config.LocalAuthEnabled = false
		config.InClusterAuthEnabled = false
		local.AuthToken = ""
	}

	tests := []struct {
		name                  string
		localAuthEnv          string
		inClusterAuthEnv      string
		expectedLocalAuth     bool
		expectedInClusterAuth bool
		shouldHaveLocalToken  bool
	}{
		{
			name:                  "Default values when env vars not set",
			localAuthEnv:          "",
			inClusterAuthEnv:      "",
			expectedLocalAuth:     true,
			expectedInClusterAuth: false,
			shouldHaveLocalToken:  true,
		},
		{
			name:                  "Local auth explicitly enabled",
			localAuthEnv:          "true",
			inClusterAuthEnv:      "",
			expectedLocalAuth:     true,
			expectedInClusterAuth: false,
			shouldHaveLocalToken:  true,
		},
		{
			name:                  "Local auth disabled, in-cluster auth enabled",
			localAuthEnv:          "false",
			inClusterAuthEnv:      "true",
			expectedLocalAuth:     false,
			expectedInClusterAuth: true,
			shouldHaveLocalToken:  false,
		},
		{
			name:                  "Invalid local auth value defaults to true",
			localAuthEnv:          "invalid",
			inClusterAuthEnv:      "",
			expectedLocalAuth:     true,
			expectedInClusterAuth: false,
			shouldHaveLocalToken:  true,
		},
		{
			name:                  "Invalid in-cluster auth value defaults to false",
			localAuthEnv:          "false",
			inClusterAuthEnv:      "invalid",
			expectedLocalAuth:     false,
			expectedInClusterAuth: false,
			shouldHaveLocalToken:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state before each test
			reset()

			// Set environment variables for test
			if tt.localAuthEnv != "" {
				os.Setenv("LOCAL_AUTH_ENABLED", tt.localAuthEnv)
			}
			if tt.inClusterAuthEnv != "" {
				os.Setenv("IN_CLUSTER_AUTH_ENABLED", tt.inClusterAuthEnv)
			}

			// Run the configuration
			Configure()

			// Check local auth configuration
			require.Equal(t, tt.expectedLocalAuth, config.LocalAuthEnabled)

			// Check in-cluster auth configuration
			require.Equal(t, tt.expectedInClusterAuth, config.InClusterAuthEnabled)

			// Check local auth token
			if tt.shouldHaveLocalToken {
				require.NotEmpty(t, local.AuthToken)
				require.Len(t, local.AuthToken, 96)
			} else {
				require.Empty(t, local.AuthToken)
			}
		})
	}
}

func TestLocalAuthRequestHandler(t *testing.T) {
	config.LocalAuthEnabled = true
	local.AuthToken = "test-token"

	// Expected response in local mode unless an error has occurred
	userResp := &UserResponse{
		InClusterAuth: false,
	}

	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedResp   *UserResponse
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
				local.SessionID = sessionID
				req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID})
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RequestHandler)

			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code, "handler returned wrong status code")

			// Verify response body
			if tt.expectedResp != nil {
				var resp UserResponse
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
		expectedResp   *UserResponse
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
			expectedResp: &UserResponse{
				InClusterAuth:     true,
				Group:             "admin-group",
				Name:              "Doug Unicorn",
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
				var resp UserResponse
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
