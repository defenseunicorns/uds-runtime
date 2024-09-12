# UDS Runtime

UDS Runtime is the frontend for all things UDS, providing views and insights into your UDS cluster.

## Quickstart Deploy

> !**WARNING**  
> UDS Runtime is in early alpha, expect breaking changes to functionality

### Pre-requisites

Recommended:
* [UDS-CLI](https://github.com/defenseunicorns/UDS-CLI#install)

If building locally:
* `Go >= 1.22.0`
* `Node >= v21.1.0`

### In Cluster

Assumes a K8s cluster is running and the appropriate K8s context has been selected

```bash
uds zarf package deploy oci://ghcr.io/defenseunicorns/packages/uds/uds-runtime:<tag> --confirm
```

**See [all tags](https://github.com/defenseunicorns/uds-runtime/pkgs/container/packages%2Fuds%2Fuds-runtime)*

### Locally (Out of Cluster)

1. clone this repo
1. compile: `uds run compile`
1. run: `./build/main`

## Quickstart Development

For a full guide on developing for UDS Runtime, please read the [CONTRIBUTING.md](./CONTRIBUTING.md)

### To start the backend development server, run the following command:

**With UDS-CLI**
```bash
uds run dev-server
```

**Without UDS-CLI**
```bash
air
```

> **NOTE**: If you do not have air installed, you can find instructions for how to install at [here](https://github.com/air-verse/air)

### To start the frontend server, run the following command:

**With UDS-CLI**
```bash
uds run dev-ui
```

**Wihtout UDS-CLI**
```bash
cd ui
npm ci
npm run dev
```

## Nightly Releases

UDS Runtime publishes a canary release of latest changes every night tagged `nightly-unstable`

```bash
uds zarf package deploy oci://ghcr.io/defenseunicorns/packages/uds/uds-runtime:nightly-unstable
```

## Tech Stack

- Backend:

  - [Golang](https://go.dev/)
  - [Chi HTTP Router](https://github.com/go-chi/chi)
  - [K8s client-go](https://github.com/kubernetes/client-go)

- Frontend:

  - [Sveltekit](https://kit.svelte.dev/)
  - [Vite](https://vitejs.dev/)
  - [Typescript](https://typescriptlang.org/)
  - [TailwindCSS](https://tailwindcss.com/) ([Flowbite](https://flowbite.com/))
  - [Carbon Icons](https://www.carbondesignsystem.com/guidelines/icons/library)
  - [svelte-apexcharts](https://github.com/bn3t/svelte-apexcharts)

- Networking:

  - [Server Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
  - [REST API](https://restfulapi.net/)
  - [K8s Shared Informers](https://pkg.go.dev/k8s.io/client-go/informers)
