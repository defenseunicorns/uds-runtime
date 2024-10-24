<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { Card } from '$components'
  import { type CarbonIcon } from 'carbon-icons-svelte'

  type IconThemeType = 'primary' | 'warning' | 'critical'

  const themeMapping = {
    primary: {
      bgColor: 'bg-blue-200',
      bgColorDark: 'dark:bg-blue-900',
      iconColor: '#3f83f8',
    },
    warning: {
      bgColor: 'bg-yellow-100',
      bgColorDark: 'bg-yellow-100',
      iconColor: '#E3A008',
    },
    critical: {
      bgColor: 'bg-pink-200',
      bgColorDark: 'bg-pink-200',
      iconColor: '#F98080',
    },
  } as const

  export let statText: string
  export let helperText: string
  export let icon: typeof CarbonIcon
  export let link: string = ''
  export let iconTheme: IconThemeType = 'primary'

  $: iconThemeOption = themeMapping[iconTheme]
</script>

<Card {link}>
  <div class="flex items-center justify-start space-x-4">
    <div class="{iconThemeOption['bgColor']} {iconThemeOption['bgColorDark']} p-3 rounded-md">
      <svelte:component this={icon} size={24} color={iconThemeOption['iconColor']} />
    </div>

    <div class="flex flex-col items-start">
      <dt
        class="text-3xl font-semibold text-black dark:text-white truncate"
        data-testid={`resource-count-${helperText.split(' ')[0].toLowerCase()}`}
      >
        {statText}
      </dt>
      <dd class="mt-1 text-sm text-gray-500 dark:text-gray-400">
        {helperText}
      </dd>
    </div>
  </div>
</Card>
