// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package local

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)

	require.Equal(t, sessionID, storage.sessionID, "expected sessionID to be stored correctly")
}

func TestValidateSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)

	require.True(t, storage.Validate(sessionID), "expected sessionID to be valid")

	invalidSessionID := "invalid-session-id"
	require.False(t, storage.Validate(invalidSessionID), "expected invalid sessionID to be invalid")
}

func TestRemoveSession(t *testing.T) {
	storage := NewBrowserSession()
	sessionID := "test-session-id"

	storage.Store(sessionID)
	storage.Remove()

	require.Empty(t, storage.sessionID, "expected sessionID to be empty after removal")
	require.False(t, storage.Validate(sessionID), "expected sessionID to be invalid after removal")
}
