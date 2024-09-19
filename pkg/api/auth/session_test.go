// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreSession(t *testing.T) {
	storage := NewInMemoryStorage()
	sessionID := "test-session-id"

	storage.StoreSession(sessionID)

	require.Equal(t, sessionID, storage.sessionID, "expected sessionID to be stored correctly")
}

func TestValidateSession(t *testing.T) {
	storage := NewInMemoryStorage()
	sessionID := "test-session-id"

	storage.StoreSession(sessionID)

	require.True(t, storage.ValidateSession(sessionID), "expected sessionID to be valid")

	invalidSessionID := "invalid-session-id"
	require.False(t, storage.ValidateSession(invalidSessionID), "expected invalid sessionID to be invalid")
}

func TestRemoveSession(t *testing.T) {
	storage := NewInMemoryStorage()
	sessionID := "test-session-id"

	storage.StoreSession(sessionID)
	storage.RemoveSession()

	require.Empty(t, storage.sessionID, "expected sessionID to be empty after removal")
	require.False(t, storage.ValidateSession(sessionID), "expected sessionID to be invalid after removal")
}
