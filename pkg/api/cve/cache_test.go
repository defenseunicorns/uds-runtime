// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package cve

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetFindings(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		page     int
		limit    int
		cache    []Finding
		expected []Finding
		total    int
		err      error
	}{
		{
			name:     "Empty Cache",
			page:     1,
			limit:    10,
			cache:    []Finding{},
			expected: nil,
			total:    0,
			err:      nil,
		},
		{
			name:     "Invalid Limit",
			page:     1,
			limit:    0,
			cache:    []Finding{{}, {}, {}},
			expected: nil,
			total:    3,
			err:      errors.New("limit must be greater than 0"),
		},
		{
			name:     "Invalid Page",
			page:     0,
			limit:    10,
			cache:    []Finding{{}, {}, {}},
			expected: nil,
			total:    3,
			err:      errors.New("page must be greater than 0"),
		},
		{
			name:     "Valid Pagination",
			page:     1,
			limit:    2,
			cache:    []Finding{{}, {}, {}},
			expected: []Finding{{}, {}},
			total:    3,
			err:      nil,
		},
		{
			name:     "Offset Exceeds Total",
			page:     2,
			limit:    10,
			cache:    []Finding{{}, {}, {}},
			expected: nil,
			total:    3,
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the cache
			findingCacheMutex.Lock()
			findingCache = tt.cache
			findingCacheMutex.Unlock()

			// Call the function
			result, total, err := GetFindings(tt.page, tt.limit)

			// Check the result
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}

			// Check the total
			if total != tt.total {
				t.Errorf("expected total %d, got %d", tt.total, total)
			}

			// Check the error
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
