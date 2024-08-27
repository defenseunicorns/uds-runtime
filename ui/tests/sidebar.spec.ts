// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

import { expect, test } from '@playwright/test'

test.describe('Sidebar', async () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/')
  })

  test('should open and close sidebar', async ({ page }) => {
    const sidebar = page.getByTestId('main-sidebar-test-id')
    const sidebarToggle = page.getByTestId('toggle-sidebar')

    await expect(sidebar).toBeVisible()
    await expect(sidebarToggle).toBeVisible()
    const sideBarText = page.getByText('Workloads')
    await expect(sideBarText).toBeVisible()

    // close sidebar
    await sidebarToggle.click()
    await expect(sideBarText).not.toBeVisible()

    // hover over sidebar
    await sidebar.hover()
    await expect(sideBarText).toBeVisible()

    // move away from sidebar
    await page.mouse.move(1000, 0)
    await expect(sideBarText).not.toBeVisible()

    // open sidebar
    await sidebarToggle.click()
    await expect(sideBarText).toBeVisible()
  })

  test("should expand and collapse sidebar's sections", async ({ page }) => {
    const sideBarText = page.getByText('Workloads')
    await expect(sideBarText).toBeVisible()

    // expand Workloads section
    await sideBarText.click()
    const subMenuText = page.getByText('Pods', { exact: true })
    await expect(subMenuText).toBeVisible()

    // collapse Workloads section
    await sideBarText.click()
    await expect(subMenuText).not.toBeVisible()
  })

  test('filter sidebar items', async ({ page }) => {
    const nonFilteredText = page.getByText('Configs')
    await expect(nonFilteredText).toBeVisible()

    // type in searchbar
    const filterInput = page.getByPlaceholder('Filter Pages')
    await filterInput.pressSequentially('Pods')
    await expect(filterInput).toHaveValue('Pods')

    // check if Pods is visible and Configs is not after filtering
    const filteredText = page.getByText('Pods', { exact: true })
    await expect(filteredText).toBeVisible()
    await expect(nonFilteredText).not.toBeVisible()
  })
})
