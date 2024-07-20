# Contributing to UDS Runtime

Welcome :unicorn: to the UDS Runtime! If you'd like to contribute, please reach out to one of the [CODEOWNERS](CODEOWNERS) and we'll be happy to get you started!

Below are some notes on our core software design philosophies that should help guide contributors.

## Table of Contents

1. [Code Quality and Standards](#code-quality-and-standards)
1. [How to Contribute](#how-to-contribute)
   - [Building the app](#building-the-app)
1. [Local Development](#local-development)
   - [Pre-Commit Hooks and Linting](#pre-commit-hooks-and-linting)
   - [Testing](#testing)

## Code Quality and Standards

Below are some general guidelines for code quality and standards that make UDS Runtime :sparkles:

- **Write tests that give confidence**: Unless there is a technical blocker, every new feature and bug fix should be tested in the project's automated test suite.

- **Prefer readability over being clever**: We have a strong preference for code readability in UDS Runtime. Specifically, this means things like: naming variables appropriately, keeping functions to a reasonable size and avoiding complicated solutions when simple ones exist.

- **Design Decision**: We use [Architectural Decision Records](https://adr.github.io/) to document the design decisions that we make. These records live in the `adr` directory. We highly recommend reading through the existing ADRs to understand the context and decisions that have been made in the past, and to inform current development.

### Continuous Delivery

Continuous Delivery is core to our development philosophy. Check out [https://minimumcd.org](https://minimumcd.org/) for a good baseline agreement on what that means.

Specifically:

- We do trunk-based development (`main`) with short-lived feature branches that originate from the trunk, get merged into the trunk, and are deleted after the merge
- We don't merge code into `main` that isn't releasable
- We perform automated testing on all changes before they get merged to `main`
- We create immutable release artifacts

## How to Contribute

Please ensure there is a Gitub issue for your proposed change, this helps the UDS Runtime team to understand the context of the change and to track the progress of the work. If there isn't an issue for your change, please create one before starting work. The recommended workflow for contributing is as follows:

\*Before starting development, we highly recommend reading through the UDS Runtime [documentation](https://uds.defenseunicorns.com/) and our [ADRs](./adr).

1. **Fork this repo** and clone it locally
1. **Create a branch** for your changes
1. **Create, [test](#testing)** your changes
1. **Add docs** where appropriate
1. **Push your branch** to your fork
1. **Open a PR** against the `main` branch of this repo

## Local Development

Most of the actions needed for running and testing UDS Runtime are contained in tasks ran by UDS CLI's `run` feature (ie. vendored [Maru](https://github.com/defenseunicorns/maru-runner)). While the actions can be performed manually without running tasks, we recommend installing the [`uds` binary](https://uds.defenseunicorns.com/cli/quickstart-and-usage/) and using tasks as much as possible.

> !NOTE  
> Tasks are used in CI. See the [pull request workflow](.github/workflows/pr-tests.yaml) for an example.

A list of runnable tasks from `uds run --list-all`

| Name                 | Description                                                                                                    |
| -------------------- | -------------------------------------------------------------------------------------------------------------- |
| dev-server           | run the api server in dev mode (requires air https://github.com/air-verse/air?tab=readme-ov-file#installation) |
| dev-ui               | run the ui in dev mode                                                                                  |
| test:e2e             | run end-to-end tests (assumes api server is running on port 8080)                                              |
| test:go              | run api server unit tests                                                                                      |
| test:ui-unit         | run frontend unit tests                                                                                        |
| test:unit            | run all unit tests (backend and frontend)                                                                      |
| lint:all             | Run all linters                                                                                                |
| lint:golangci        | Run golang linters                                                                                             |
| lint:yaml            | Run yaml linters                                                                                               |
| lint:ui              | Run ui linters                                                                                                 |
| lint:format-ui       | Format ui code                                                                                                 |
| setup:build-api      | build the go api server                                                                                        |
| setup:build-ui       | build ui                                                                                                |
| setup:slim-cluster   | Create a k3d cluster and deploy core slim dev with metrics server                                              |
| setup:simple-cluster | Create a k3d cluster, no core                                                                                  |
| setup:golangci       | Install golangci-lint to GOPATH using install.sh                                                               |
| setup:clone-core     | Clone uds-core for custom slim dev setup                                                                       |
| setup:metrics-server | Create and deploy metrics server from cloned core                                                              |

### Pre-Commit Hooks and Linting

In this repo you can optionally use [pre-commit](https://pre-commit.com/) hooks for automated validation and linting, but if not CI will run these checks for you

### Testing

We strive to test all changes made to UDS Runtime. If you're adding a new feature or fixing a bug, please add tests to cover the new functionality. Unit tests and E2E tests are both welcome, but we leave it up to the contributor to decide which is most appropriate for the change. Below are some guidelines for testing:

#### Unit Tests

Unit tests reside alongside the source code in a `*_test.go` file or `*.test.ts` file. These tests should be used to test individual functions or components, or (in a more integration style not E2E) small flows of coupled functions / components.

For running unit tests:

- run all -- `uds run test:unit`
- run backend only -- `uds run test:go`
- run frontend only -- `uds run test:ui-unit`

#### E2E Tests

E2E tests reside in the `ui/tests/` directory and can be named `*.test.ts` or `*.spec.ts`. Run E2E tests via `uds run test:e2e`. This task will:

1. build the ui
1. build the api server
1. setup the slim cluster (core-slim-dev + metrics server)
1. run the e2e script, which starts the api server (serves ui) to test against.
