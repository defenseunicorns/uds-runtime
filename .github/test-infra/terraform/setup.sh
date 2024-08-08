#!/bin/bash -x

# Get instance public IP
public_ip="$(curl -s https://checkip.amazonaws.com/)"

# Create k3d cluster
# k3d cluster create uds \
#     --api-port 6443 \
#     --port 80:80@loadbalancer \
#     --port 443:443@loadbalancer \
#     --k3s-arg "--tls-san=$public_ip@server:*"

# Build and Deploy Bundle
cat <<EOF > /tmp/uds-bundle.yaml
    kind: UDSBundle
    metadata:
      name: runtime-test
      description: A UDS bundle for deploying UDS Runtime with UDS Core
      version: 0.1.0

    packages:
      - name: uds-k3d-dev
        repository: ghcr.io/defenseunicorns/packages/uds-k3d
        ref: 0.8.0

      - name: init
        repository: ghcr.io/defenseunicorns/packages/init
        ref: v0.35.0

      - name: core
        repository: ghcr.io/defenseunicorns/packages/uds/core
        ref: 0.25.0-upstream
        optionalComponents:
          - istio-passthrough-gateway
          - metrics-server

      - name: runtime
        repository: ghcr.io/defenseunicorns/packages/uds/uds-runtime
        ref: nightly-unstable
        overrides:
          uds-runtime:
            uds-runtime:
              variables:
                - name: IMG_TAG
                  path: image.tag
                  default: nightly-unstable
EOF

uds create /tmp --confirm -o /tmp || exit 1
uds deploy /tmp/uds-bundle-runtime-test-*.tar.zst --confirm || exit 1
