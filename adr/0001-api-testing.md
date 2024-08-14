# 1. API Testing

Date: 14 August 2024

## Status

In progress

## Context

API testing is crucial for ensuring the quality and reliability of the UDS Runtime API. It helps to identify bugs, performance issues, and security vulnerabilities early in the development process. In this context, API testing refers to the process of testing the API endpoints, request/response payloads, and the overall behavior of the API.

### Considerations

#### Data

In order to reliably and repeatedly test the API, we need a consistent set of data that can be used to validate the API responses. This data should be representative of the data that the API will encounter in production. To that end, we need to decide what data should be present in the test Kubernetes cluster. Options include:

- Deploying the UDS core slim dev bundle and testing against known data from that deployment
- Seeding the cluster with fake data for each test or set of tests

#### Test Coverage

Because the UDS Runtime backend code is reasonably generic with respect to creating endpoints for retrieving Kubernetes resources, there isn't a need to test every type of resource. Furthermore, because there are parameters for each request (i.e., dense, sparse, SSE, etc.), we don't want to test every combination of parameters. To that end, we should base our testing strategy and coverage on the bind groupings of resources currently implemented by the API. For reference, those bind groupings are:

- Core resources
- Workload resources
- UDS resources
- Config resources
- Cluster Ops resources
- Network resources
- Storage resources

Furthermore, there are endpoints that are specific to UDS, such as Pepr endpoints, that are present under the `/monitor` path; for example
- `/api/v1/monitor/pepr`
- `/api/v1/monitor/pepr/{stream}`
- `/api/v1/monitor/cluster-overview`

#### Assertions

Options for test assertions include:
- Testing the exact expected JSON response of the various API endpoints
- Marshalling the received JSON response into a Kubernetes type and comparing the struct to expected data

## Decision

We will use the `net/http/httptest` package from the Go standard library to write tests for the UDS Runtime API.

- For test data, we will use a minimal version of UDS Core to include Pepr and Istio, and deploy a simple app such as podinfo that is integrated with UDS.
- For test coverage, we will test endpoints from each of the bind groupings listed above, with the addition of testing the more custom endpoints under the `/monitor` path.
- For assertions, we will marshal the received JSON response into a Kubernetes struct and test key fields from the expected response.
