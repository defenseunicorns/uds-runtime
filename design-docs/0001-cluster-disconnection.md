# Design Doc Title: Cluster Disconnection Detection

Author(s): Runtime Team
Date Created: Sept. 9, 2024
Status: Implemented
Ticket: https://github.com/defenseunicorns/uds-runtime/issues/10

## Problem Statement

It is important for real-time monitoring and maintaining of a kubernetes cluster, that users are made aware when their connection to the cluster is no longer healthy. This is especially important given that Runtime uses a cache (built from kubernetes informers), which continues to serve potentially outdated information upon cluster disconnection. This design aims to solve this problem for local (out of cluster) deployments of Runtime by implementing a system that detects cluster disconnection and automatically attempts to reconnect, while providing feedback to users.

## Proposal

The solution involves creating a mechanism that constantly monitors the health of the cluster connection. If the connection is lost, it will trigger a reconnection attempt in the background and notify the user via toast notifications in the frontend. Upon reconnection, a success message will be shown to the user, the current view will be updated, and the system will continue normal operations.

## Scope and Requirements

- Detect disconnection from the cluster
- Automatically attempt to reconnect when a disconnection is detected
- Notify users of both disconnection and successful reconnection via the frontend UI
- Ensure the system can handle reconnection attempts without causing a complete failure or downtime
- Keep monitoring the cluster connection state at regular intervals

## Implementation Details

The implementation will consist of the following components:

### Backend Implementation:

#### Detection:

Option 1: Use Informers

This approach uses the [watch error handler](https://github.com/kubernetes/client-go/blob/v0.20.5/tools/cache/shared_informer.go#L169-L182) that every informer already implements to detect disconnection. By doing so, we would not need to poll the cluster with a separate endpoint. We could then, upon disconnection, implement reconnection logic. The issue found in testing this option is that when informers connect successfully they don't then detect disconnection errors immediately. It seems as though they don't get the error until some timeout is hit, likely closing the TCP connection. An attempt at setting the timeout for TCP connections to a lower value, did not make a difference.

Option 2: Poll

We poll the cluster with a server health check. Initiating this health check requires the frontend to make a request to `/health`. If an error is encountered, the system will emit a disconnection event, triggering the reconnection process in a separate go routine.

**Solution**

Due to complications mentioned in option 1, we have chosen to implement option 2.

#### Reconnecting

When triggered by an update to the disconnected error channel, the reconnection handler will cancel the current cache context (officially stopping the informers), attempt to recreate the Kubernetes client, and reinitialize the cache. This loop will continue until the connection is restored or the application is stopped. **Note:** reconnection attempts will only be made to the originally connected cluster based on the original current-context and cluster name.

### Frontend Implementation:

When a user lands on the application, triggering the main `src/routes/layout.svelte`, an EventSource is created for `/health` that will now continuously receive updates from the server on cluster connection health. If an error is received, a toast will be displayed to the user. This error toast should remain on the screen (regardless of user navigation) until a reconnected message is received. Only a single toast will be added regardless of subsequent error messages. When the connection is restored, the toast is updated to indicate reconnection and then removed.

After the reconnection event is received, a custom event will be dispatched to trigger an event listener on the `Datatable` component. This event handler will then restart the table store, which re-fires the eventSource call and gets data from the new cache.

## Changes to Existing Systems:

- Add a new health check route (/health) to monitor the cluster's connection status.
- Introduce reconnection handling logic in the backend to manage the lifecycle of Kubernetes clients and caches.
- Wrap route handlers so they dynamically get the latest cache. Otherwise, routes will always maintain the cache reference they were initialized with and therefore never return data from the new informers.
- Add event listeners and dispatchers in UI and make stores from createStore(), in Datatable, reactive

## Current Problems and Questions

1. By intiating the health check via the frontend call to `/health`, this creates a poll per client. Could this cause strain on the server? Could all these potential error events (for the same cluster disconnection) cause unnecessary reconnection attempts?

1. If the connection check interval is too low, there is potential for unnecessarily initiating reconnection attempts. This occurs when an error kicks off reconnection handling, the reconnection is successful, but the check occurs with the old clientset before the new one can be set and therefore sends another error to the disconnected error channel.

1. Are there any side-effects of this running in-cluster?

## Non-Goals

This solution does not include handling other Kubernetes API failures unrelated to cluster disconnections. It does not provide full error recovery for all possible API failures.

## Future Improvements

It would be nice to figure out a way to use the informers for detecting disconnection events without the need for polling. This could also open up the ability to send more specific errors regarding resources to the frontend for admins to see.

e.g.

```console
E0909 09:36:59.759711 2545122 reflector.go:158] "Unhandled Error" err="pkg/mod/k8s.io/client-go@v0.31.0/tools/cache/reflector.go:243: Failed to watch uds.dev/v1alpha1, Resource=exemptions: failed to list uds.dev/v1alpha1, Resource=exemptions: the server could not find the requested resource" logger="UnhandledError"
```

## Other Considerations

- The impact of reconnection attempts on system performance should be monitored, particularly in environments with high traffic.
