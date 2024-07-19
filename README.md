# UDS Engine

## Development

To start the backend server, run the following command:

```bash
air
```

> **NOTE**: If you do not have air installed, you can find insturctions for how to install at [here](https://github.com/air-verse/air)

To start the frontend server, run the following command:

```bash
cd ui && npm run dev
```

## Techstack

- Backend:

  - [Golang](https://go.dev/)
  - [Chi HTTP Router](https://github.com/go-chi/chi)
  - [K8s client-go](https://github.com/kubernetes/client-go)

- Frontend:

  - [Sveltekit](https://kit.svelte.dev/)
  - [Vite](https://vitejs.dev/)
  - [TailwindCSS](https://tailwindcss.com/) ([Flowbite](https://flowbite.com/))
  - [Carbon Icons](https://www.carbondesignsystem.com/guidelines/icons/library)

- Networking:

  - [Server Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
  - [REST API](https://restfulapi.net/)
  - [K8s Shared Informers](https://pkg.go.dev/k8s.io/client-go/informers)
