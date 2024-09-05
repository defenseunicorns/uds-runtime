<script lang="ts">
  export let text: string = ''

  console.log(text)
</script>

<!-- <div class="flex relative group w-full z-50 transform -translate-y-0">
  <div class="w-full">{text}</div>
  <div class="absolute top-0 transform -translate-y-0 p-2 bg-gray-900 text-white text-xs rounded shadow-lg opacity-0 group-hover:opacity-90 transition-opacity duration-300 pointer-events-none whitespace-nowrap">{text}</div>
</div> -->

<div class="tooltip-2 w-full truncate overflow-visible z-50">
  {text}
  <span class="tooltiptext w-full z-50">{text}</span>
</div>

<style lang="css">
  /* Tooltip container */
  .tooltip-2 {
    position: relative;
  }

  /* Tooltip text */
  .tooltip-2 .tooltiptext {
    visibility: hidden;
    background-color: black;
    color: #fff;
    text-align: center;
    /* padding: 5px 0; */
    /* height: 10px; */

    /* Position the tooltip text - see examples below! */
    position: absolute;
    left: 0;
    height: 30px;
    width: 300px;
  }

  /* Show the tooltip text when you mouse over the tooltip container */
  .tooltip-2:hover .tooltiptext {
    visibility: visible;
  }
</style>
