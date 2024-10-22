<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  import { onMount } from 'svelte'

  import Convert from 'ansi-to-html'

  const convert = new Convert({
    newline: true,
    stream: true,
    escapeXML: true,
  })
  let termElement: HTMLElement | null
  let scrollAnchor: Element | null | undefined

  // exported in parent component to handle incoming SSE messages
  export const addMessage = (message: string) => {
    let html = convert.toHtml(message)
    // Print the html or a non-breaking space if the message is empty to preserve line breaks
    html = `<div class="zarf-terminal-line">${html || '&nbsp;'}</div>`
    scrollAnchor?.insertAdjacentHTML('beforebegin', html)
    scrollAnchor?.scrollIntoView()
  }

  onMount(() => {
    termElement = document.getElementById('terminal')
    scrollAnchor = termElement?.lastElementChild
  })
</script>

<div class="scroll-anchor" />
