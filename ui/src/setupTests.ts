// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import '@testing-library/jest-dom'

const mockComponent = () => ({
  $$: {
    on_mount: [],
    on_destroy: [],
    before_update: [],
    after_update: [],
  },
})

// Mock the entire component
vi.mock('$components', () => {
  return {
    DataTable: vi.fn().mockImplementation(mockComponent),
    Link: vi.fn().mockImplementation(mockComponent),
  }
})
