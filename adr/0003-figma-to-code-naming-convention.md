# 3 Figma to Code Naming Convention

## Status

proposed

## Context

In order to try to keep some sort of connection between the Figma designs and the components on the front end that are being built.

## Decision

A decision was made to name components based on the naming used by the flowbite Figma library, for components that we are using that are either blocks or what Figma calls "widgets".

### Examples...

In this image you can see 3 different things outlined

- The Name of the component at large which is called Stats Widget
- The selected component
- The name of the selected compoent or type i.e. "With right icon"

<img width="1881" alt="Screenshot 2024-09-30 at 12 27 02 PM" src="https://github.com/user-attachments/assets/479f4b4d-b495-4f07-92b9-21cd112df4c4">

In this case, this would have been a component that has variants, but rather than built extra abstractions and complication, the components were built out as individual components inside of a folder called `StatsWidget`

<img width="1606" alt="Screenshot 2024-09-30 at 12 30 54 PM" src="https://github.com/user-attachments/assets/a3518cf7-f703-4cbf-842a-f65561b57731">

In this image we can see...

- The name of the folder `StatsWidget`
- The name of the compoennt `WithRightIcon`
- The component
- The attributes for the component which line up with the attributes that the component has i.e.
  - Stats text
  - Helper text

<img width="920" alt="Screenshot 2024-09-30 at 12 32 38 PM" src="https://github.com/user-attachments/assets/e1095ab8-60bb-4cab-a4c8-51ce17194ce9">

All attributes were not included, but the ones we did include were named based on this components attributes in Figma

## Rationale

### Creates a Convention

In deciding to name components based on Figma (flowbite) designs, we are using a naming convention that everyone can follow.

## Consequences

If flowbite updates their component library in Figma and the names change, then there is going to be a disconnect between the library naming and the front end components.
