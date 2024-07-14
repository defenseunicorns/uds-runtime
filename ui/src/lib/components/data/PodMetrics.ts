// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023-Present The UDS Authors

/**
 * Parse CPU usage string
 *
 * @param cpuString CPU usage string
 * @returns CPU usage in millicpu
 * @throws Error if the input is not a valid string
 */
export function parseCPU(cpuString: string): number {
  // Ensure the input is a valid string
  if (typeof cpuString !== 'string') {
    throw new Error('Invalid input: CPU string must be a string')
  }

  // If the CPU usage is reported in nanocores (ending with 'n'), convert to millicores
  if (cpuString.endsWith('n')) {
    return parseFloat(cpuString.slice(0, -1)) / 1e6
  }

  // Otherwise, assume the CPU usage is reported in cores and convert to millicores
  return parseFloat(cpuString) * 1000
}

/**
 * Convert memory usage to Mi
 *
 * @param memoryString Memory usage string
 * @returns Memory usage in Mi
 * @throws Error if the input is not a valid string
 */
export function parseMemory(memoryString: string): number {
  // Ensure the input is a valid string
  if (typeof memoryString !== 'string') {
    throw new Error('Invalid input: Memory string must be a string')
  }

  // Parse the numeric value from the string
  const value = parseFloat(memoryString)

  // If the memory usage is reported in kibibytes (ending with 'Ki'), convert to mebibytes
  if (memoryString.endsWith('Ki')) {
    return value / 1024
  }

  // If the memory usage is reported in mebibytes (ending with 'Mi'), return the value as-is
  if (memoryString.endsWith('Mi')) {
    return value
  }

  // If the memory usage is reported in gibibytes (ending with 'Gi'), convert to mebibytes
  if (memoryString.endsWith('Gi')) {
    return value * 1024
  }

  // Otherwise, assume the memory usage is reported in bytes and convert to mebibytes
  return value / (1024 * 1024)
}

/**
 * Format CPU usage to a human-readable string
 *
 * @param millicpu CPU usage in millicpu
 * @returns A human-readable string
 */
export function formatCPU(millicpu: number): string {
  // If the CPU usage is 1000 millicores or more, format as cores
  if (millicpu >= 1000) {
    return `${(millicpu / 1000).toFixed(2)} cpu`
  }
  // If the CPU usage is 100 to 999 millicores, format as integer millicores
  else if (millicpu >= 100) {
    return `${millicpu.toFixed(0)} m`
  }
  // If the CPU usage is 10 to 99 millicores, format as one decimal place millicores
  else if (millicpu >= 10) {
    return `${millicpu.toFixed(1)} m`
  }
  // If the CPU usage is less than 10 millicores, format as two decimal places millicores
  else {
    return `${millicpu.toFixed(2)} m`
  }
}

/**
 * Format memory usage to a human-readable string
 *
 * @param mb Memory usage in Mi
 * @returns A human-readable string
 */
export function formatMemory(mb: number): string {
  // If the memory usage is 1024 Mi or more, format as Gi
  if (mb >= 1024) {
    return `${(mb / 1024).toFixed(2)} Gi`
  }
  // If the memory usage is 100 to 1023 Mi, format as rounded integer Mi
  else if (mb >= 100) {
    return `${Math.round(mb)} Mi`
  }
  // If the memory usage is less than 100 Mi, format as one decimal place Mi
  else {
    return `${mb.toFixed(1)} Mi`
  }
}
