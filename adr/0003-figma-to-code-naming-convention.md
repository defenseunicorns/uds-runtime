# 3 Figma to Code Naming Convention

## Status

proposed

## Context

Devs and designers are currently using different language to describe various Flowbite components. Designers use language from Flowbite's Figma library and devs are using generic language. This discrepancy in word choice is causing confusion when reviewing PRs

## Decision

A decision was made to name components based on the naming used by the flowbite Figma library, for components that we are using that are either blocks or what Figma calls "widgets".

### Examples...

In this image you can see 3 different things outlined

- The Name of the component at large which is called Stats Widget
- The selected component
- The name of the selected compoent or type i.e. "With right icon"

![Alt text](../ui/static/stats-widgets.png?raw=true)

In this case, this would have been a component that has variants

![Alt text](../ui/static/stats-widget-code.png?raw=true)

In this image we can see...

- The name of the folder `StatsWidget`
- The name of the compoennt `WithRightIcon`
- The component
- The attributes for the component which line up with the attributes that the component has i.e.
  - Stats text
  - Helper text

![Alt text](../ui/static/stats-widget-attributes.png?raw=true)

All attributes were not included, but the ones we did include were named based on this components attributes in Figma

## Rationale

### Creates a Convention

In deciding to name components based on Figma (Flowbite) designs, we are using a naming convention that everyone can follow.

## Consequences

If Flowbite updates their component library in Figma and the names change, then there is going to be a disconnect between the library naming and the front end components.
