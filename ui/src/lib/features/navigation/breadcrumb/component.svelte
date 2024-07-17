<!-- SPDX-License-Identifier: Apache-2.0 -->
<!-- SPDX-FileCopyrightText: 2024-Present The UDS Authors -->

<script lang="ts">
  import { ChevronRight } from 'carbon-icons-svelte'

  import { page } from '$app/stores'
  import { stringToSnakeCase } from '$lib/utils'

  import { routes } from '../routes'
  import { type Route, type RouteChild } from '../types'

  const flatRoutes = routes.flatMap((route) => {
    if (route.children) {
      return [route, ...route.children]
    }

    return route
  })

  let matchParent: Route
  let matchChild: RouteChild

  // Subscribe to the page store to get the current URL (regular $ doesnt seem to work here)
  page.subscribe((value) => {
    const { pathname } = value.url

    // Find the current route
    ;[matchParent, matchChild] = flatRoutes.filter((route) => {
      if (pathname === '/') {
        return false
      }

      if (route.path === '/' && pathname !== '/') {
        return false
      }

      return pathname.includes(route.path)
    })
  })
</script>

<nav class="flex" aria-label="Breadcrumb">
  <ol class="inline-flex items-center space-x-1 md:space-x-2 rtl:space-x-reverse">
    {#if matchParent}
      <li>
        <div class="flex items-center">
          <span
            class="inline-flex items-center text-sm font-medium text-gray-700 dark:text-gray-400"
            data-testid={`breadcrumb-item-${stringToSnakeCase(matchParent?.name)}`}
          >
            {matchParent?.name}
          </span>
        </div>
      </li>
    {/if}
    {#if matchChild}
      <li><ChevronRight class="w-5 h-5 text-gray-400 dark:text-gray-400" /></li>
      <li>
        <div class="flex items-center">
          <span
            class="inline-flex items-center text-sm font-medium text-gray-700 dark:text-gray-400"
            data-testid={`breadcrumb-item-${stringToSnakeCase(matchChild.name)}`}
          >
            {matchChild.name}
          </span>
        </div>
      </li>
    {/if}
  </ol>
</nav>
