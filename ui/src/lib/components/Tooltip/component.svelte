<!-- Copyright 2024 Defense Unicorns -->
<!-- SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial -->

<script lang="ts">
  export let title = ''
  let isHovered = false
  let xOffset: number
  let yOffset: number
  let isWider: boolean = false

  type MouseEventType = MouseEvent & {
    currentTarget: EventTarget & HTMLDivElement
  }

  const mouseOver = (event: MouseEventType) => {
    const createElWidth = getElementWidth('div', event.currentTarget.children[0].textContent)

    // This offset is the padding for the table > tr > td element in app.postcss. The class of px-4 is 1em or 16px
    if (createElWidth - event.currentTarget.offsetLeft * 2 > event.currentTarget.clientWidth) {
      isWider = true
    } else {
      isWider = false
    }

    isHovered = true
    xOffset = event.pageX
    yOffset = event.pageY
  }
  const mouseMove = (event: MouseEventType) => {
    // calculate if the columns is a certain percent off right side and render on the left side
    const diff = ((window.innerWidth - event.pageX) / window.innerWidth) * 100

    const tooltipOffset = diff < 15 ? 500 : 200
    xOffset = event.pageX - tooltipOffset
    yOffset = event.pageY - 110
  }

  const mouseLeave = () => (isHovered = false)

  const getElementWidth = (type: string, textContent: string | null): number => {
    if (textContent) {
      const newEl = document.createElement(type)
      const newContent = document.createTextNode(textContent)

      // add the text node to the newly created div
      newEl.appendChild(newContent)
      newEl.setAttribute('id', 'test-el')
      // Inline block so it does not take full width of screen
      newEl.style.display = 'inline-block'
      // Ensure the element is not visible
      newEl.style.position = 'fixed'
      newEl.style.zIndex = '0'

      // add the newly created element and its content into the DOM
      const bodyEl = document.querySelector('body')!
      bodyEl.appendChild(newEl)
      const createdEL = document.getElementById('test-el')
      const createElWidth = createdEL?.clientWidth
      bodyEl.removeChild(newEl)

      return createElWidth || 0
    }

    return 0
  }
</script>

<div
  id="container"
  role="button"
  tabindex={0}
  on:mouseover={mouseOver}
  on:mouseleave={mouseLeave}
  on:mousemove={mouseMove}
  on:focus
>
  <slot />
</div>

{#if isHovered && isWider}
  <div
    class="absolute bg-gray-900 text-white opacity-90 text-xs focus:border-gray-200 focus:outline-none focus:ring-0 dark:border-gray-700 p-4 rounded-[4px]"
    style="top: {yOffset}px; left: {xOffset}px; padding: 10px"
  >
    {title}
  </div>
{/if}
