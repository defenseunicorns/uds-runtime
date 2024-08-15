# 2. Front End Data Test IDs (data-testid)

## Status

proposed

## Context

Front-end testing, it is often necessary to select and interact with elements on the web page to validate functionality, whether you are running unit tests or integration/ end-to-end tests. There are many different ways you can select elements, including CSS selectors, IDs, and attributes like data-testid. The choice of method impacts the maintainability and reliability of the tests

When using a framework, the included locators can be used, and should be by default (i.e. getByRole, getByText, getByPlaceholder). For cases where those do not work, rather than trying to use the structure of the html to locate elements, we should use data-testid's because the framework we are using for testing provides a getByTestId locator

## Decision

The decision is to use the data-testid attribute for selecting elements during front-end testing when one of the default test framework selectors/ locators do not apply

In the case that there might be a collision with test id's (i.e. `data-table`), in an environment where multiple components might be loaded into a page, consider using the component name like a namespace i.e. `data-testid="pods-data-table"` and `data-testid="jobs-data-table"`

When iterating over a list, a unique id associated with the given item or the mapping structure index can be used i.e. `data-testid="jobs-data-listitem-{row.id}"`

## Rationale

### Stability of Test Selectors:

- Avoiding Changes Due to Styling Updates: CSS classes and IDs are often tied to the styling or structural organization of the page. Any changes to the UI, even purely visual, can break tests if these selectors are used. Using data-testid ensures that tests are insulated from such changes, as it is solely intended for testing purposes.

### Clarity and Intent:

- Explicit Purpose: The data-testid attribute makes it clear that the element is being selected for testing. This avoids confusion for developers and testers, making the intent behind the selection clear and reducing the risk of accidentally removing or altering a selector that is crucial for tests.

### Maintainability:

- Isolation from Production Code: By using a dedicated attribute like data-testid, test selectors are decoupled from production code. This isolation minimizes the impact on production features and avoids cluttering the markup with test-related classes or IDs that could confuse or complicate code maintenance.

### Consistency Across Tests:

- Standardization: Adopting data-testid as the standard practice across all tests promotes consistency, making the test suite easier to understand and maintain. It simplifies the review process and onboarding for new developers, who can rely on a uniform approach to element selection.

### Best Practices:

- Industry Adoption: The use of data-testid is a widely accepted best practice in front-end testing communities, especially with tools like React Testing Library, Cypress and PLaywright. Following this convention aligns the project with industry standards, facilitating collaboration and leveraging community support.

## Consequences

### Positive:

- Increased test stability and robustness.
- Clear separation of concerns between test code and production code.
- Enhanced maintainability and readability of the test suite.

### Negative:

- Leads to additional work of adding data-testid's to elements and following best practice for naming them
