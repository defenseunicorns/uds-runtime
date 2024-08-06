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
uds create /tmp/bundle/ --confirm -o /tmp/bundle/
uds deploy /tmp/bundle/uds-bundle-runtime-test-*.tar.zst --confirm --set uds-k3d-dev.K3D_EXTRA_ARGS="--k3s-arg --tls-san=$public_ip@server:*"

# Edit kubeconfig for remote access
mkdir -p /home/ubuntu/.kube
k3d kubeconfig get interview > /home/ubuntu/.kube/config
kubeconfig=$(sed 's/0\.0\.0\.0/'"$public_ip"'/g' /home/ubuntu/.kube/config)
echo "$kubeconfig" > /home/ubuntu/kubeconfig-remote.yaml
