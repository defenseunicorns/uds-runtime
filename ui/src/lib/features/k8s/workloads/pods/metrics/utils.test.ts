// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

import { expect, suite, test } from 'vitest'

import { formatCPU, formatMemory, parseCPU, parseMemory } from './utils'

suite('CPU and Memory Utility Functions', () => {
  suite('parseCPU', () => {
    test('should parse CPU usage in nanocores', () => {
      expect(parseCPU('1000000n')).toBe(1)
      expect(parseCPU('500000n')).toBe(0.5)
    })

    test('should parse CPU usage in cores', () => {
      expect(parseCPU('1')).toBe(1000)
      expect(parseCPU('0.5')).toBe(500)
    })

    test('should throw an error for invalid input', () => {
      expect(() => parseCPU(123 as unknown as string)).toThrow('Invalid input: CPU string must be a string')
    })
  })

  suite('parseMemory', () => {
    test('should parse memory usage in Ki', () => {
      expect(parseMemory('1024Ki')).toBe(1)
      expect(parseMemory('2048Ki')).toBe(2)
    })

    test('should parse memory usage in Mi', () => {
      expect(parseMemory('1Mi')).toBe(1)
      expect(parseMemory('2Mi')).toBe(2)
    })

    test('should parse memory usage in Gi', () => {
      expect(parseMemory('1Gi')).toBe(1024)
      expect(parseMemory('2Gi')).toBe(2048)
    })

    test('should parse memory usage in bytes', () => {
      expect(parseMemory('1048576')).toBe(1)
      expect(parseMemory('2097152')).toBe(2)
    })

    test('should throw an error for invalid input', () => {
      expect(() => parseMemory(123 as unknown as string)).toThrow('Invalid input: Memory string must be a string')
    })
  })

  suite('formatCPU', () => {
    test('should format CPU usage >= 1000 millicores', () => {
      expect(formatCPU(1000)).toBe('1.00 cpu')
      expect(formatCPU(1500)).toBe('1.50 cpu')
    })

    test('should format CPU usage between 100 and 999 millicores', () => {
      expect(formatCPU(500)).toBe('500 m')
      expect(formatCPU(999)).toBe('999 m')
    })

    test('should format CPU usage between 10 and 99 millicores', () => {
      expect(formatCPU(50)).toBe('50.0 m')
      expect(formatCPU(99.9)).toBe('99.9 m')
    })

    test('should format CPU usage < 10 millicores', () => {
      expect(formatCPU(1)).toBe('1.00 m')
      expect(formatCPU(9.99)).toBe('9.99 m')
    })
  })

  suite('formatMemory', () => {
    test('should format memory usage >= 1024 Mi', () => {
      expect(formatMemory(1024)).toBe('1.00 Gi')
      expect(formatMemory(2048)).toBe('2.00 Gi')
    })

    test('should format memory usage between 100 and 1023 Mi', () => {
      expect(formatMemory(500)).toBe('500 Mi')
      expect(formatMemory(1023)).toBe('1023 Mi')
    })

    test('should format memory usage < 100 Mi', () => {
      expect(formatMemory(50)).toBe('50.0 Mi')
      expect(formatMemory(99.9)).toBe('99.9 Mi')
    })
  })
})
