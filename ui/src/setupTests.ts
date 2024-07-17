// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import '@testing-library/jest-dom'

// Mock the entire component
vi.mock('$components', () => {
  const mockComponent = () => ({
    $$: {
      on_mount: [],
      on_destroy: [],
      before_update: [],
      after_update: [],
    },
  })

  return {
    DataTable: vi.fn().mockImplementation(mockComponent),
  }
})
